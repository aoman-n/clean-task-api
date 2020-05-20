import { combineReducers } from '@reduxjs/toolkit'
import sampleModule, { SampleState } from './sampleModule'
import projectModule, { ProjectState } from './project'
import taskModule, { TaskState } from './task'
import authModule, { AuthState } from './auth'

export interface RootState {
  sample: SampleState
  project: ProjectState
  task: TaskState
  auth: AuthState
}

export const rootReducer = combineReducers({
  sample: sampleModule.reducer,
  project: projectModule.reducer,
  task: taskModule.reducer,
  auth: authModule.reducer,
})
