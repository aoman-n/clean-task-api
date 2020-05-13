import { ExNextPageContext } from 'next'
import Router from 'next/router'

export const dashboardTask = `/dashboard/:id/tasks`

export const redirect = (res: ExNextPageContext['res'], url: string) => {
  if (res) {
    res.writeHead(302, { Location: url }).end()
  } else {
    Router.push(url)
  }
}
