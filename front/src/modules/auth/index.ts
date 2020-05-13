import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface AuthState {
  token: string
}

export const authInitialState: AuthState = {
  token: '',
}

const taskModule = createSlice({
  name: 'auth',
  initialState: authInitialState,
  reducers: {
    setToken: (state: AuthState, action: PayloadAction<string>) => {
      return { ...state, token: action.payload }
    },
  },
})

export const { setToken } = taskModule.actions

export default taskModule
