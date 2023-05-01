import Head from 'next/head'
import { SideBar } from "./SideBar";
import { TextBox } from './TextBox';

export default function Home() {
  return (
    <div>
      <Head>
        <title>Halaman Utama</title>
        <link rel="icon" href="/favicon.ico" />
        <style>{`
          body {
            background-color: #393E46;
          }
        `}</style>
      </Head>
      <TextBox />
      <SideBar />
    </div>
  )
}
