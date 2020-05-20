export interface Project {
  id: number
  title: string
  description: string
  role: string
}

export interface Task {
  id: number
  name: string
  status: number
}

export interface Tag {
  id: number
  name: string
  color: string
}
