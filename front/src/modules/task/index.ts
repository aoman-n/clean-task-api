import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Task } from '~/services/model'

export interface TaskState {
  // tasks: { [key in string]: Task }
  tasks: Task[]
}

export const taskInitialState: TaskState = {
  // tasks: {},
  tasks: [],
}

const taskModule = createSlice({
  name: 'tasks',
  initialState: taskInitialState,
  reducers: {
    setTasks: (state: TaskState, action: PayloadAction<Task[]>) => {
      return { ...state, tasks: action.payload }
    },
    addTask: (state: TaskState, action: PayloadAction<Task>) => {
      return { ...state, tasks: [...state.tasks, action.payload] }
    },
  },
})

export const { setTasks, addTask } = taskModule.actions

export default taskModule
