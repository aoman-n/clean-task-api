import axios from 'axios'

const BASE_URL_ON_SERVER = 'http://api:8080/api'
const BASE_URL_ON_FRONT = 'http://localhost:4000/api'

const isServer = typeof window === 'undefined'

export const getBaseUrl = () => {
  if (isServer) return BASE_URL_ON_SERVER

  return BASE_URL_ON_FRONT
}

export const createAuthHeader = (token: string | undefined) => ({
  Authorization: `Bearer ${token}`,
})

/* Response Data Structure */
export type Response<T> = {
  message: string
  data: T
}

interface Login {
  id: string
  token: string
}

export const loginApi = async (params: {
  loginName: string
  password: string
}) => {
  try {
    const res = await axios.post<Response<Login>>(
      `${getBaseUrl()}/login`,
      params,
    )

    return res.data.data
  } catch (e) {
    console.log('e: ', e)

    throw e
  }
}

export const getPrivateMessage = async (token: string) => {
  return axios.get<{ message: string }>(`${getBaseUrl()}/private`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })
}

export const fetchHello = async () => {
  const resp = await axios.get<{ message: string }>(
    `${getBaseUrl()}/cookie_sample`,
  )

  return resp.data
}

export const fetchCookie = async () => {
  const resp = await axios.get<{ message: string }>(
    `http://localhost:4000/cookie_sample`,
  )

  return resp.data
}

export const postProfile = async (
  token: string,
  formData: { displayName: string },
) => {
  const headers = createAuthHeader(token)
  return axios.post(`${getBaseUrl()}/profile`, formData, { headers })
}
