import React, { useContext, useState } from "react";
import IconButton from "@mui/material/IconButton";
import menuContext from "./appstate";
import textContent from "./appstate";
import Sidebar from "./Sidebar";
import { text } from "stream/consumers";


const Openbutton: React.FC = () => {
  const { isOpened, setOpened } = useContext(menuContext);
  const [buttontext,setText] = useState("open");

  const toggleOpen = () => {
    setOpened(!isOpened);  
    setText(isOpened ? "メニュー":"閉じる")
  };
  
  return (
    <div>
      <button onClick={toggleOpen} id="Openbutton">{buttontext}</button>
      { isOpened && <Sidebar />}
    </div>
  );
};

export default Openbutton;
