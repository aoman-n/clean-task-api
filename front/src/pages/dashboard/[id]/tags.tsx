import { NextPage } from 'next'
import Dashboard from '~/components/templates/Dashboard'

const DashboardPage: NextPage = () => {
  return <Dashboard selectedKey="tags" content={<div>tags</div>} />
}

export default DashboardPage
