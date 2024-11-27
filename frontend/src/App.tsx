import { useState } from 'react'; 
import { Routes, Route } from "react-router-dom";
import Header from './component/Header';
import Sidebar from './component/Sidebar';
import Openbutton from './component/Openbutton';
import menuContext from './component/appstate';
import "./App.css";
import Timetable from './pages/Timetable';
import Nomatch from './pages/Nomatch';
import Classes from './pages/Classes';
import LoginForm from './pages/LoginForm';
import SignupForm from './pages/SignupForm';

const App: React.FC = () => {
  const [isOpened, setOpened] = useState(true);
  const [isLogin,setLogin] = useState(false);
  const [isSignup,setSignup] = useState(false)

    return (
      <menuContext.Provider value={{ isOpened, setOpened }}>
        <>
        {isLogin? (
          <>
            <Header />
            <Openbutton />
            <Routes>
              <Route path='/timetable' element={<Timetable />} />
              <Route path='/classes' element={<Classes />} />
              <Route path='*' element={<Nomatch />} />
            </Routes>
          </>
        ):isSignup ?(
          <LoginForm />
        ):(
          <SignupForm />
        )}
      </> 
      </menuContext.Provider>
  );
};

export default App;