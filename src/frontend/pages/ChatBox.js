import { useState, useEffect, useRef } from 'react';
import { faPaperPlane } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import axios from 'axios';

import styles from '../styles/ChatBox.module.css';

const bot = [];
export function ChatBox() {
  const [messages, setMessages] = useState([]);
  const messageContainerRef = useRef(null);
  const [botMessage, setBotMessage] = useState([]);

  const handleSendMessage = async (e) => {
    e.preventDefault();
    const inputEl = e.target.elements.message;
    const message = inputEl.value.trim();
    if (message) {
      setMessages([...messages, { sender: 'me', message }]);
      inputEl.value = '';
      const response = await axios.get(`http://localhost:8080/api/product/${message}`);
      const dataa = response.data.product;
      console.log(dataa);
      if (dataa){
        setBotMessage([...botMessage, { sender: 'me', dataa }]);
        bot.push(dataa);
        console.log(botMessage);
        console.log(message);
      }
    }
  };

  useEffect(() => {
    if (messageContainerRef.current) {
      messageContainerRef.current.scrollTo({
        top: messageContainerRef.current.scrollHeight,
        behavior: 'smooth',
      });
    }
  }, [messages]);

  return (
    <div>
      <div className={styles.chatbox}>
        <div ref={messageContainerRef} className={styles.messageContainer}>
          {messages.map((msg, idx) => (
            <div key={idx} className={styles.bubbleContainer}>
              <div className={`${styles.bubble} ${msg.sender === 'me' ? styles.right : styles.left}`}>
                {msg.message}
              </div>
              <div key={idx} className={`${styles.bubble} ${msg.sender === 'me' ? styles.left:styles.right}`}>
                {bot[idx]}
              </div>
            </div>
          ))}
        </div>
      </div>
      <div className={styles.TextBox}>
        <form onSubmit={handleSendMessage} className={styles.form}>
          <input
            type="text"
            name="message"
            placeholder="Type a message..."
            className={styles.input}
          />
          <button type="submit" className={styles.sendButton}>
            <FontAwesomeIcon icon={faPaperPlane} className={styles.faPaperPlane} />
          </button>
        </form>
      </div>
    </div>
  );
}
