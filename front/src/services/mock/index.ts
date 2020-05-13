import { Project, Task } from '~/services/model'

export const mockProjects: Project[] = [
  {
    id: 1,
    title: 'プロジェクト1',
    description: 'Description 1111111111111',
    role: 'admin',
  },
  {
    id: 2,
    title: 'プロジェクト222',
    description: 'Description 2222222222222',
    role: 'admin',
  },
  {
    id: 3,
    title: 'プロジェクト3',
    description: 'Description 3333333333333333',
    role: 'admin',
  },
  {
    id: 4,
    title: 'プロジェクト4',
    description: 'Description 444444444444444',
    role: 'admin',
  },
]

export const mockTasks: Task[] = [
  {
    id: 1,
    name: 'Ant Design Title 1',
    status: 1,
  },
  {
    id: 2,
    name: 'Ant Design Title 2',
    status: 1,
  },
  {
    id: 3,
    name: 'Ant Design Title 3',
    status: 1,
  },
  {
    id: 4,
    name: 'Ant Design Title 4',
    status: 1,
  },
]
