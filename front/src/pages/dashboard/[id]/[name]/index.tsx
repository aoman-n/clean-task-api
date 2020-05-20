import { NextPage, ExNextPageContext } from 'next'
import Dashboard2 from '~/components/templates/Dashboard2'
import Footer from '~/components/molcules/Footer'
import MainContents from '~/components/organisms/MainContents'
import { fetchTasks } from '~/services/api/tasks'
import { redirect } from '~/routes'
import { setTasks } from '~/modules/task'
import { select } from '~/modules/project'

const Dashboard: NextPage = () => {
  return (
    <Dashboard2 MainComponent={<MainContents />} FooterComponent={<Footer />} />
  )
}

Dashboard.getInitialProps = async (ctx: ExNextPageContext) => {
  if (!ctx.auth.jwt) {
    redirect(ctx.res, '/login')
    return
  }

  const projectId = Number(ctx.query.id)
  ctx.store.dispatch(select(projectId))

  try {
    const tasks = await fetchTasks(projectId, ctx.auth.jwt)

    ctx.store.dispatch(setTasks(tasks))
    return
  } catch (e) {
    console.log('fetchTasks error: ', e)
  }
}

export default Dashboard
