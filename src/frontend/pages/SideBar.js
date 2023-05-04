import React, { useState } from "react";
import sidebar from '../styles/SideBar.module.css'

export const SideBar = () => {
  const [isOn1, setIsOn1] = useState(false);
  const [isOn2, setIsOn2] = useState(false);

  const handleToggle1 = () => {
    setIsOn1(true);
    setIsOn2(false);
  };

  const handleToggle2 = () => {
    setIsOn1(false);
    setIsOn2(true);
  };

  return (
    <div className={sidebar.sidebar}>
      <div className={sidebar.sidebarTop}>
        <ul>
          <li><button className={sidebar.newButton}>+ New Chat</button></li>
          <li><button>histori 1</button></li>
          <li><button>histori 2</button></li>
          <li><button>histori 3</button></li>
        </ul>
      </div>
      <div className={sidebar.sidebarBottom}>
        <ul>
          <li>
            <label className={sidebar.switch}>
              <input type="checkbox" checked={isOn1} onChange={handleToggle1} />
              <span className={sidebar.slider}></span>
            </label>
            <span className={isOn1 ? sidebar.orange : '    '}>  KMP</span>
          </li>
          <li>
            <label className={sidebar.switch}>
              <input type="checkbox" checked={isOn2} onChange={handleToggle2} />
              <span className={sidebar.slider}></span>
            </label>
            <span className={isOn2 ? sidebar.orange : '   '}>  BM</span>
          </li>
        </ul>
      </div>
    </div>
  );
};
