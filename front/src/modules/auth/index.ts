import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface AuthState {
  jwt?: string
}

export const authInitialState: AuthState = {
  jwt: undefined,
}

const taskModule = createSlice({
  name: 'auth',
  initialState: authInitialState,
  reducers: {
    setToken: (state: AuthState, action: PayloadAction<string | undefined>) => {
      return { ...state, jwt: action.payload }
    },
  },
})

export const { setToken } = taskModule.actions

export default taskModule
