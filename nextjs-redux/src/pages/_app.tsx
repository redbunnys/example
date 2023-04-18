import '@/styles/globals.css'
import type { AppProps } from 'next/app'
import { store } from '@/store/store'
import { Provider } from 'react-redux'
import Link from 'next/link'


export default function App({ Component, pageProps }: AppProps) {
  return (
    <Provider store={store}>
      <h1>nihao</h1>
      <Link href='/'>index</Link>
      <Link href='/counter'>counter</Link>
      <Link href='/num'>num</Link>
      <Component {...pageProps} />
    </Provider>
  )

}
