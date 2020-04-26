import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface Task {
  id: number
  title: string
  description: string
}

export interface ProjectState {
  list: Task[]
  selected: number
}

export const projectInitialState: ProjectState = {
  list: [],
  selected: 0,
}

const projectModule = createSlice({
  name: 'project',
  initialState: projectInitialState,
  reducers: {
    setTasks: (state: ProjectState, action: PayloadAction<Task[]>) => {
      return { ...state, list: action.payload }
    },
  },
})

export const { setTasks } = projectModule.actions

export default projectModule
