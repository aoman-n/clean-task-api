import React from 'react'
// import { NextPageContext } from 'next'
// import Head from 'next/head'
import App, { AppContext } from 'next/app'
// import Router, { useRouter } from 'next/router'
// import NProgress from 'nprogress'
// import { ThemeProvider } from 'styled-components'
import { loadAuthFromCookie, AuthContext, Auth } from '~/auth'
// import theme from '~/theme'
// import { GlobalStyle } from '~/styles'
// Provider を読み込む
import { Provider } from 'react-redux'
// setupStore を読み込む
import { makeStore } from '~/store'
import withRedux from 'next-redux-wrapper'
import { Store } from 'redux'
import { RootState } from '~/modules/rootState'

// import { RootState, rootReducer } from '~/modules/rootState'

console.log('process.env.NODE_ENV: ', process.env.NODE_ENV)

// const store = setupStore()

// const makeStore: MakeStore = (initialState: RootState) => {
//   return createStore(rootReducer, initialState)
// }

// Router.events.on('routeChangeStart', url => {
//   console.log(`Loading: ${url}`)
//   NProgress.start()
// })
// Router.events.on('routeChangeComplete', () => NProgress.done())
// Router.events.on('routeChangeError', () => NProgress.done())

// const MyComponent: React.FC<{ children: React.ReactElement }> = ({
//   children,
// }) => {
//   const router = useRouter()

//   return (
//     <>
//       <Head>
//         {/* Import CSS for nprogress */}
//         <link rel="stylesheet" type="text/css" href="/nprogress.css" />
//       </Head>
//       <GlobalStyle />
//       <Provider store={store}>
//         <ThemeProvider theme={theme}>
//           {React.cloneElement(children, {
//             key: router.route,
//           })}
//         </ThemeProvider>
//       </Provider>
//     </>
//   )
// }

type Props = {
  Component: React.Component
  store: Store<RootState>
}

class MyApp extends App<Props> {
  static async getInitialProps({ Component, ctx }: AppContext) {
    console.log('In _app getInitialProps')
    let pageProps = {}

    const auth = loadAuthFromCookie(ctx)

    if (Component.getInitialProps) {
      pageProps = await (Component as any).getInitialProps({ ...ctx, auth })
      // pageProps = await Component.getInitialProps(ctx)
    }

    return { pageProps, auth }
  }

  render() {
    const { Component, pageProps, store } = this.props

    console.log('############### _appのrender内 store: ', store)

    // return (
    //   <MyComponent>
    //     <AuthContext.Provider value={auth}>
    //       <Component {...pageProps} />
    //     </AuthContext.Provider>
    //   </MyComponent>
    // )

    return (
      <Provider store={store}>
        <Component {...pageProps} />
      </Provider>
    )
  }
}

export default withRedux(makeStore, {
  debug: false,
})(MyApp)
