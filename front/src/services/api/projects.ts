import axios from 'axios'
import { Response, getBaseUrl, createAuthHeader } from '~/utils/api'
import { Project } from '~/services/model'

export const fetchProjects = async (token: string) => {
  const headers = createAuthHeader(token)

  const res = await axios.get<Response<Project[]>>(`${getBaseUrl()}/projects`, {
    headers,
  })

  return res.data.data
}
