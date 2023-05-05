import React, { useState } from "react";
import sidebar from '../styles/SideBar.module.css'
import axios from "axios";

export const SideBar = () => {
  const [isOn1, setIsOn1] = useState(true);
  const [isOn2, setIsOn2] = useState(false);

  const handleToggle1 = async() => {
    setIsOn1(true);
    setIsOn2(false);
    console.log(isOn1);
    try {
      const response = await axios.post('http://localhost:8080/api/radio-button', {"value": isOn1 });
      console.log(response.data);
      console.log('Radio value submitted!');
    } catch (error) {
      console.error(error);
    }
    // console.log(response);
  };

  const handleToggle2 = async () => {
    setIsOn1(false);
    setIsOn2(true);
    console.log(isOn1);
    try {
      const response = await axios.post('http://localhost:8080/api/radio-button', { "value": isOn1 });
      console.log(response.data);
      console.log('Radio value submitted!');
    } catch (error) {
      console.error(error);
    }
  };

  const handleAddButton = () => {
    window.location.reload();
  }


  return (
    <div className={sidebar.sidebar}>
      <div className={sidebar.sidebarTop}>
        <ul>
          <li><button className={sidebar.newButton} onClick={handleAddButton}>+ New Chat</button></li>
        </ul>
      </div>
      <div className={sidebar.sidebarBottom}>
        <ul>
          <li>
            <label className={sidebar.switch}>
              <input type="checkbox" checked={isOn1} onChange={handleToggle1} name="my-radio-button"/>
              <span className={sidebar.slider}></span>
            </label>
            <span className={isOn1 ? sidebar.orange : '    '}>  KMP</span>
          </li>
          <li>
            <label className={sidebar.switch}>
              <input type="checkbox" checked={isOn2} onChange={handleToggle2} name="my-radio-button"/>
              <span className={sidebar.slider}></span>
            </label>
            <span className={isOn2 ? sidebar.orange : '   '}>  BM</span>
          </li>
        </ul>
      </div>
    </div>
  );
};
