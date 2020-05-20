import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { Project, Tag } from '~/services/model'

export interface ProjectState {
  projects: Project[]
  tags: Tag[]
  selected: number
}

export const projectInitialState: ProjectState = {
  projects: [],
  tags: [],
  selected: 0,
}

const projectModule = createSlice({
  name: 'project',
  initialState: projectInitialState,
  reducers: {
    setProjects: (state: ProjectState, action: PayloadAction<Project[]>) => {
      return { ...state, projects: action.payload }
    },
    select: (state: ProjectState, action: PayloadAction<number>) => {
      return { ...state, selected: action.payload }
    },
  },
})

export const { setProjects, select } = projectModule.actions

export default projectModule
