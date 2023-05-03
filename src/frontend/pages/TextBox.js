import React, { useState } from 'react';
import styles from '../styles/TextBox.module.css';

export const TextBox = () => {
  const [text, setText] = useState('');
  const [messages, setMessages] = useState([]);
  const [lastSender, setLastSender] = useState("");

  const handleTextChange = (event) => {
    setText(event.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    if (text.trim() !== "") {
      setMessages([...messages, { text, sender: "me" }]);
      setText("");
      setLastSender("me");
    }
  };
  

  return (
    <div className={styles.TextBox}>
      <form onSubmit={handleSubmit}>
        <input type="text" value={text} onChange={handleTextChange} placeholder="Send a message."/>
      </form>
      <div className={`${styles.container}`}>
        {messages.map((message, index) => (
        <div key={index}>
        <div className={`${styles.bubble} ${styles[message.sender]}`}>
          {message.text}
          {message.sender === "me" && (
            <div className={styles.arrow} />
          )}
      </div>
          {message.sender === "me" && (
          <div className={`${styles.answerBubble} ${styles.bot}`}>
          ada yang bisa saya bantu?
        </div>
      )}
    </div>
  ))}
</div>
    </div>
  );
};

// return (
//   <div className={styles.TextBox}>
//     <form onSubmit={handleSubmit}>
//       <input
//         type="text"
//         value={text}
//         onChange={handleTextChange}
//         placeholder="Send a message."
//       />
//     </form>
//     <div className={styles.messagesContainer}>
//       {messages.map((message, index) => (
//         <div key={index}>
//           <div className={`${styles.bubble} ${styles[message.sender]}`}>
//             {message.text}
//             {message.sender === 'me' && <div className={styles.arrow} />}
//           </div>
//           {message.sender === 'me' && (
//             <div className={`${styles.answerBubble} ${styles.bot}`}>
//               ada yang bisa saya bantu?
//             </div>
//           )}
//         </div>
//       ))}
//     </div>
//   </div>
// );
// };