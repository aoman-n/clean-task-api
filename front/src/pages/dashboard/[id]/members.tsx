import { NextPage } from 'next'
import Dashboard from '~/components/templates/Dashboard'
import MemberList from '~/components/organisms/MemberList'

const DashboardPage: NextPage = () => {
  return <Dashboard selectedKey="members" content={<MemberList />} />
}

export default DashboardPage
