import React from "react";
import "./LoginForm.css"

const SignupForm:React.FC = () => {
    return (
        <div className="formContainer">
            <form>
                <h1>アカウント作成</h1>
                <hr />
                <div className="uiForm">
                    <div className="formField">
                        <label>ユーザー名</label>
                        <input type="text" placeholder="ユーザー名" name="username" />
                    </div>
                    <div className="formField">
                        <label>メールアドレス</label>
                        <input type="text" placeholder="メールアドレス" name="email" />
                    </div>
                    <div className="formField">
                        <label>パスワード</label>
                        <input type="password" placeholder="パスワード" name="password" />
                    </div>
                    <button type="submit" className="loginButton">ログイン</button>
                </div>
            </form>
        </div>
    );
};

export default SignupForm