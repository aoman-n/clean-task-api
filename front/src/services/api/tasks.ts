import axios from 'axios'
import { Response, getBaseUrl, createAuthHeader } from '~/utils/api'
import { Task } from '~/services/model'

export const fetchTasks = async (
  projectId: number,
  jwt: string | undefined,
) => {
  const headers = createAuthHeader(jwt)

  const res = await axios.get<Response<Task[]>>(
    `${getBaseUrl()}/projects/${projectId}/tasks`,
    {
      headers,
    },
  )

  return res.data.data
}

export const postTask = async (
  token: string,
  projectId: string,
  params: { name: string },
) => {
  const headers = createAuthHeader(token)

  const res = await axios.post<Response<Task>>(
    `${getBaseUrl()}/projects/${projectId}/tasks`,
    params,
    {
      headers,
    },
  )

  return res.data.data
}
