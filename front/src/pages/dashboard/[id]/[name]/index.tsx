import { NextPage, ExNextPageContext } from 'next'
import Link from 'next/link'
import { Layout, Breadcrumb, Button, Card } from 'antd'
import { ProjectOutlined } from '@ant-design/icons'
import styled from 'styled-components'
import TaskList from '~/components/organisms/TaskList'
import { fetchTasks } from '~/services/api/tasks'
import { redirect } from '~/routes'
import { setTasks } from '~/modules/task'
import { select } from '~/modules/projectModule'

const { Header, Content, Footer } = Layout

const tabList = [
  {
    key: 'task',
    tab: 'タスク',
  },
  {
    key: 'tag',
    tab: 'タグ',
  },
  {
    key: 'member',
    tab: 'メンバー',
  },
]

const Dashboard: NextPage = () => {
  return (
    <>
      <Layout style={{ minHeight: '100vh' }}>
        <Layout>
          <StyledHeader style={{ zIndex: 1, width: '100%' }}>
            <Link href="/projects">
              <StyledButton ghost icon={<ProjectOutlined />}>
                プロジェクトリスト
              </StyledButton>
            </Link>
          </StyledHeader>
          <Content style={{ padding: '0 50px' }}>
            <Breadcrumb style={{ margin: '16px 0' }}>
              <Breadcrumb.Item>Dashboard</Breadcrumb.Item>
              <Breadcrumb.Item>プロジェクト名</Breadcrumb.Item>
            </Breadcrumb>
            <LayoutContent>
              <Card
                tabList={tabList}
                activeTabKey="task"
                onTabChange={() => {}}
                style={{ width: '100%' }}
              >
                <TaskList />
              </Card>
            </LayoutContent>
          </Content>
          <Footer style={{ textAlign: 'center' }}>
            Ant Design ©2018 Created by Ant UED
          </Footer>
        </Layout>
      </Layout>
    </>
  )
}

const StyledHeader = styled(Header)`
  background: #036aa7;
`
const StyledButton = styled(Button)`
  font-weight: bold;
`
const LayoutContent = styled.div`
  background: #fff;
  padding: 24px;
  min-height: 280px;
`

Dashboard.getInitialProps = async (ctx: ExNextPageContext) => {
  if (!ctx.auth.token) {
    redirect(ctx.res, '/login')
    return
  }

  const projectId = Number(ctx.query.id)
  ctx.store.dispatch(select(projectId))

  try {
    const tasks = await fetchTasks(ctx.auth.token, projectId)

    ctx.store.dispatch(setTasks(tasks))
    return
  } catch (e) {
    console.log('fetchTasks error: ', e)
  }
}

export default Dashboard
