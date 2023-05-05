import { useState, useEffect, useRef } from 'react';
import { faPaperPlane } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import styles from '../styles/ChatBox.module.css';

export function ChatBox() {
  const [messages, setMessages] = useState([]);
  const messageContainerRef = useRef(null);

  const handleSendMessage = (e) => {
    e.preventDefault();
    const inputEl = e.target.elements.message;
    const message = inputEl.value.trim();
    if (message) {
      setMessages([...messages, { sender: 'me', message }]);
      inputEl.value = '';
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
              <div className={`${styles.bubble} ${msg.sender === 'me' ? styles.left:styles.right}`}>
                gw males jawab chat lu
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
