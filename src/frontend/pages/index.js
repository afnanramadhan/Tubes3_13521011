// import Head from 'next/head'
// import { SideBar } from "./SideBar";
// import { TextBox } from './TextBox';
// import styles from '../styles/Home.module.css';

// export default function Home() {
//   return (
//     <div>
//       <Head>
//         <title>Halaman Utama</title>
//         <link rel="icon" href="/favicon.ico" />
//         <style>{`
//           body {
//             background-color: #393E46;
//           }
//         `}</style>
//       </Head>
//       <TextBox />
//       <SideBar />
//     </div>
//   )
// }


import Head from 'next/head';
import { SideBar } from "./SideBar";
import { TextBox } from './TextBox';
import styles from '../styles/Home.module.css';

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
        <div className={styles['home-atas']}>
        </div>
        <div className={styles['home-bawah']}>
          <TextBox />
        </div>
      </div>
      {/* <div className={`${styles.column} ${styles['home-bawah']}`}></div> */}
    </div>
  );
}

