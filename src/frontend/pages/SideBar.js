import React, { useState } from "react";
import sidebar from '../styles/SideBar.module.css'

export const SideBar = () => {
  const [isOn, setIsOn] = useState(false);

  const handleToggle = () => {
    setIsOn(prev => !prev);
  };

  return (
    <div className={sidebar.sidebar}>
      <ul>
        <li><button className={sidebar.newButton}>+ New Chat</button></li>
        <li><button>histori 1</button></li>
        <li><button>histori 2</button></li>
        <li>
          <label className={sidebar.switch}>
            <input type="checkbox" checked={isOn} onChange={handleToggle} />
            <span className={sidebar.slider}></span>
          </label>
        </li>
      </ul>
    </div>
  );
};
