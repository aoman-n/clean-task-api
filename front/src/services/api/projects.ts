import axios from 'axios'
import { Response, getBaseUrl, createAuthHeader } from '~/utils/api'
import { Project, Tag } from '~/services/model'

export const fetchProjects = async (token: string | undefined) => {
  const headers = createAuthHeader(token)

  const res = await axios.get<Response<Project[]>>(`${getBaseUrl()}/projects`, {
    headers,
  })

  return res.data.data
}

export const fetchTags = async (
  projectId: string,
  token: string | undefined,
) => {
  const headers = createAuthHeader(token)

  const res = await axios.get<Response<Tag[]>>(
    `${getBaseUrl()}/projects/${projectId}/tags`,
    {
      headers,
    },
  )

  return res.data.data
}
