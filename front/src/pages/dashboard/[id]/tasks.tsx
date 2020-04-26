import { NextPage } from 'next'
import Dashboard from '~/components/templates/Dashboard'
import { useRouter } from 'next/router'

const DashboardPage: NextPage = () => {
  const router = useRouter()
  console.log('router: ', router)

  return <Dashboard selectedKey="tasks" content={<div>tasks</div>} />
}

export default DashboardPage
