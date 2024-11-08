from flask import Flask, render_template, redirect ,session, request
from sqlalchemy import func,Column,Integer,String,Boolean,JSON
from flask_login import UserMixin, LoginManager, login_user,logout_user, login_required
from werkzeug.security import generate_password_hash, check_password_hash
from setting import Session_local, Base, engine
import requests


app = Flask(__name__)
app.secret_key = 'Sk264575'

login_manager = LoginManager()
login_manager.init_app(app)



class AppClass(Base):
    __tablename__ = 'mandatory_cl'
    class_id = Column(Integer, primary_key=True)
    name = Column(String(100))
    faculty = Column(String(100))
    department = Column(String(100))
    is_mandatory = Column(Boolean)
    datetime = Column(String(10))


class UserInfo(Base):
    __tablename__ = 'user_inform'
    id = Column(Integer, primary_key=True, autoincrement=True)
    user_id = Column(Integer)
    class_data = Column(JSON)
    
class User(UserMixin,Base):
    __tablename__ = 'users'
    user_id = Column(Integer, primary_key=True)
    password = Column(String(1000),nullable=False)

    def get_id(self):
        return self.user_id
    
Base.metadata.create_all(bind=engine)

@login_manager.user_loader
def load_user(user_id):
    with Session_local() as db_session:
        return db_session.query(User).get(int(user_id))
 

def get_mandatory(faculty,department):
    with Session_local() as db_session:
        classlist = db_session.query(AppClass).filter(
             AppClass.faculty==faculty,
             AppClass.department==department,
             AppClass.is_mandatory==True
             ).all()
        
        return classlist


@app.route('/')
def get_index():
    return render_template('index.html')


#サインアップ
@app.route('/signup.html', methods=['POST','GET'])
def get_signup():
    with Session_local() as db_session:
        if request.method == "POST":
            user_id = request.form['user_id']
            password = request.form['password']
            
            user = db_session.query(User).filter_by(user_id=user_id).first()

            if user is None:
                new_user = User(user_id=user_id,password=generate_password_hash(password, method="pbkdf2:sha256"))
                db_session.add(new_user) 
                db_session.commit() 
                return redirect('/login.html')
          
            else:
                msg = 'このユーザー表は使われています'
                return render_template('/signup.html',msg=msg,user=user)
        else:
            return render_template('/signup.html')
        
    
#ログインページ
@app.route('/login.html', methods=['POST','GET'])
def get_login():    
    with Session_local() as db_session:
        if request.method == "POST":
            user_id = request.form['user_id']
            password = request.form['password']
            session['user_id'] = user_id
            
            user = db_session.query(User).filter(User.user_id==user_id).first()
            if user and check_password_hash(user.password,password):
                login_user(user)
                return redirect('/infoform.html')  
            elif not user:
                return 'ユーザーが見つかりません'
            else:
                return 'パスワードが間違っている'
        else:
            return render_template('/login.html')
       

#学部学科入力
@app.route('/infoform.html')
def get_infoform():
    return render_template('infoform.html')

#ログアウト
@app.route('/logout')
@login_required
def get_logout():
    logout_user()
    return redirect('/login.html')


#授業一覧を表示
@app.route('/classlist.html' ,methods=["POST","GET"])
@login_required
def class_get(): 
    if request.method == 'POST':
        faculty = request.form['faculty']
        department = request.form['department']
        user_name = request.form['user_name']
        session['faculty'] = faculty
        session['department'] = department
        session['user_name'] = user_name
       
    else:
        faculty = session.get('faculty')
        department = session.get('department')
    
    classlist = get_mandatory(faculty,department)
    return render_template('classlist.html',classlist=classlist, faculty=faculty, department=department)



#登録ボタンを押して送信
@app.route('/register_class', methods=['POST'])
def register_class():
    class_name =  request.form['class_name']
    datetime = request.form['class_datetime']
    redirect_url = f'/schedule.html?class_name={class_name}&datetime={datetime}'
    return redirect(redirect_url)



#スケジュールを表示
@app.route('/schedule.html')
@login_required
def schedule_get():
    with Session_local() as db_session:
        class_name = request.args.get('class_name','')
        datetime = request.args.get('datetime','')
        user_id = session.get('user_id')
        user_name = session.get('user_name')
        organized_schedule ={}
        userinfos = db_session.query(UserInfo).filter_by(user_id=user_id).first()
        if userinfos:
            db_session.query(UserInfo).filter_by(user_id=user_id).update({
                UserInfo.class_data: func.json_set(UserInfo.class_name, f'$."{datetime}"','class_name')
            })
        else:
            organized_schedule[datetime] = class_name
            new_user = UserInfo(user_id=user_id,class_data=organized_schedule)
            db_session.add(new_user)
        db_session.commit()    
        get_class =db_session.query(UserInfo).filter_by(user_id=user_id).first().class_data
        return render_template('schedule.html',get_class=get_class,user_name=user_name)

if __name__ == "__main__":
    app.run(debug=True) 