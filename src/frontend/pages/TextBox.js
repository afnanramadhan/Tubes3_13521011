import React, { useState } from 'react';
import styles from '../styles/TextBox.module.css';

export const TextBox = () => {
  const [text, setText] = useState('');

  const handleTextChange = (event) => {
    setText(event.target.value);
  };

  return (
    <div className={styles.TextBox} >
      <input type="text" value={text} onChange={handleTextChange} placeholder="Send a message."/>
    </div>
  );
};

