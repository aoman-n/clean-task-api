import { combineReducers } from '@reduxjs/toolkit'
import projectModule, { ProjectState } from './project'
import taskModule, { TaskState } from './task'
import authModule, { AuthState } from './auth'

export interface RootState {
  project: ProjectState
  task: TaskState
  auth: AuthState
}

export const rootReducer = combineReducers({
  project: projectModule.reducer,
  task: taskModule.reducer,
  auth: authModule.reducer,
})
