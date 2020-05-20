import React from 'react'
import Head from 'next/head'
import App, { AppContext } from 'next/app'
import Router, { useRouter } from 'next/router'
import NProgress from 'nprogress'
import { ThemeProvider } from 'styled-components'
import { loadAuthFromCookie, Auth } from '~/auth'
import theme from '~/theme'
import { GlobalStyle } from '~/styles'
import { Provider } from 'react-redux'
import { makeStore } from '~/store'
import withRedux from 'next-redux-wrapper'
import { Store } from 'redux'
import { RootState } from '~/modules/rootState'
import { setToken } from '~/modules/auth'

import 'antd/dist/antd.css'

Router.events.on('routeChangeStart', (url) => {
  console.log(`Loading: ${url}`)
  NProgress.start()
})
Router.events.on('routeChangeComplete', () => NProgress.done())
Router.events.on('routeChangeError', () => NProgress.done())

const MyComponent: React.FC<{ children: React.ReactElement }> = ({
  children,
}) => {
  const router = useRouter()

  return (
    <>
      <Head>
        {/* Import CSS for nprogress */}
        <link rel="stylesheet" type="text/css" href="/nprogress.css" />
      </Head>
      <GlobalStyle />
      <ThemeProvider theme={theme}>
        {React.cloneElement(children, {
          key: router.route,
        })}
      </ThemeProvider>
    </>
  )
}

type Props = {
  Component: React.Component
  store: Store<RootState>
}

class MyApp extends App<Props> {
  static async getInitialProps({ Component, ctx }: AppContext) {
    let pageProps = {}
    let auth: Auth = { jwt: undefined }

    // サーバーで実行する時はcookieからjwtを取得
    // フロントで実行する時はstoreからjwtを取得
    if (ctx.req) {
      auth = loadAuthFromCookie(ctx)
      ctx.store.dispatch(setToken(auth.jwt))
    } else {
      const state: RootState = ctx.store.getState()
      auth = { jwt: state.auth.jwt }
    }

    if (Component.getInitialProps) {
      pageProps = await (Component as any).getInitialProps({ ...ctx, auth })
    }

    return { pageProps, auth }
  }

  render() {
    const { Component, pageProps, store } = this.props

    return (
      <MyComponent>
        <Provider store={store}>
          <Component {...pageProps} />
        </Provider>
      </MyComponent>
    )
  }
}

export default withRedux(makeStore, {
  debug: false,
})(MyApp)
