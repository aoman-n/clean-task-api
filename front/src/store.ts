import {
  configureStore,
  getDefaultMiddleware,
  EnhancedStore,
} from '@reduxjs/toolkit'
import { MakeStore } from 'next-redux-wrapper'
import { rootReducer, RootState } from './modules/rootState'
import logger from 'redux-logger'

const isDev = process.env.NODE_ENV === 'development'

export const makeStore: MakeStore = (
  initialState?: RootState,
): EnhancedStore => {
  const middlewares = [...getDefaultMiddleware()]

  if (isDev) {
    middlewares.push(logger)
  }

  const store = configureStore({
    reducer: rootReducer,
    middleware: middlewares,
    devTools: isDev,
    preloadedState: initialState,
  })

  return store
}
