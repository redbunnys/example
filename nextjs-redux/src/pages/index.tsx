import Head from 'next/head'
import Image from 'next/image'
import { Inter } from 'next/font/google'
import styles from '@/styles/Home.module.css'
import { RootState } from '@/store/store'
import { useSelector } from 'react-redux'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  const count = useSelector((state: RootState) => state.counter.value)
  return (
    <>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className={styles.main}>
        <a href='counter'>33</a>
        {count}
      </main>
    </>
  )
}
