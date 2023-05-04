import React, { useState } from 'react';
import styles from '../styles/TextBox.module.css';
import axios from 'axios';

export const TextBox = () => {
  const [text, setText] = useState('');
  const [messages, setMessages] = useState([]);
  const [lastSender, setLastSender] = useState("");
  const [botMessage, setBotMessage] = useState([]);

  const handleTextChange = (event) => {
    setText(event.target.value);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    if (text.trim() !== "") {
      setMessages([...messages, { text, sender: "me" }]);
      setText("");
      setLastSender("me");
    }
    console.log(text.trim());
    try{
      const response = await axios.get(`http://localhost:8080/api/product/${text.trim()}`);
      const dataa = response.data.product;
      console.log(dataa)
      if (dataa){
        const nbotMessage = {text : dataa, sender: "bot"};
        setBotMessage([...botMessage, nbotMessage]);
        console.log(nbotMessage);
        console.log(botMessage);
        setLastSender("bot");
      }
    }
    catch(error){
      console.log(error);
      const nbotMessage = {text : "Maaf, saya tidak mengerti", sender: "bot"};
      setBotMessage([...botMessage, nbotMessage]);
      setLastSender("bot");
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
          {message.sender === "me" && lastSender === "bot" &&(
          <div className={`${styles.answerBubble} ${styles.bot}`}>
            {botMessage[index].text}

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