import Head from 'next/head';
import { SideBar } from "./SideBar";
import styles from '../styles/Home.module.css';
import { ChatBox } from './ChatBox';
// import { ChatBox } from './ChatBox';

export default function Home() {
  return (
    <div>
      <Head>
        <title>Halaman Utama</title>
        <link rel="icon" href="/favicon.ico" />
        <style>{`
          body {
            margin: 0;
            padding: 0;
            overflow: hidden; /* prevent scrolling */
          }
        `}</style>
      </Head>
      <div className={`${styles.column} ${styles['sidebar']}`}>
        <SideBar />
      </div>
      <div className={`${styles.column} ${styles['home']}`}>
        <div className={styles.bottomBarContainer}>
          <ChatBox />
          {/* <TextBox /> */}
        </div>
      </div>
    </div>
  );
}
