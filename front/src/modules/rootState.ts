import { combineReducers } from '@reduxjs/toolkit'
import sampleModule, { SampleState } from './sampleModule'
import projectModule, { ProjectState } from './projectModule'

export interface RootState {
  sample: SampleState
  project: ProjectState
}

export const rootReducer = combineReducers({
  sample: sampleModule.reducer,
  project: projectModule.reducer,
})
