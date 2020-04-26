import React from 'react'
import Link from 'next/link'
import { Layout, Menu, Breadcrumb, Typography } from 'antd'
import {
  // UserOutlined,
  // LaptopOutlined,
  // NotificationOutlined,
  PieChartOutlined,
} from '@ant-design/icons'
import styled from 'styled-components'

// const { SubMenu } = Menu
const { Header, Content, Sider, Footer } = Layout

export const menuKeys = {
  tasks: 'task',
  members: 'member',
  tags: 'tags',
} as const

type MenuKey = keyof typeof menuKeys

interface DashBoard {
  selectedKey: MenuKey
  content: React.ReactNode
  // projectId: number
}

const DashBoard: React.FC<DashBoard> = ({ selectedKey, content }) => {
  return (
    <StyledLayout style={{ minHeight: '100vh' }}>
      <Header className="header">
        <Logo />
        <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['2']}>
          <Menu.Item key="1">nav 1</Menu.Item>
          <Menu.Item key="2">nav 2</Menu.Item>
          <Menu.Item key="3">nav 3</Menu.Item>
        </Menu>
      </Header>
      <Layout>
        <Sider width={260} className="site-layout-background">
          <Menu
            mode="inline"
            selectedKeys={[menuKeys[selectedKey]]}
            style={{ height: '100%', borderRight: 0 }}
          >
            <Menu.Item key={menuKeys.tasks}>
              <Link href={`tasks`}>
                <a>
                  <PieChartOutlined />
                  <span>タスク</span>
                </a>
              </Link>
            </Menu.Item>
            <Menu.Item key={menuKeys.members}>
              <Link href={`members`}>
                <a>
                  <PieChartOutlined />
                  <span>メンバー</span>
                </a>
              </Link>
            </Menu.Item>
            <Menu.Item key={menuKeys.tags}>
              <Link href={`tags`}>
                <a>
                  <PieChartOutlined />
                  <span>タグ</span>
                </a>
              </Link>
            </Menu.Item>
          </Menu>
        </Sider>
        <Layout style={{ padding: '24px 24px 24px' }}>
          {content}
          {/* <Breadcrumb style={{ margin: '16px 0' }}>
            <Breadcrumb.Item>Home</Breadcrumb.Item>
            <Breadcrumb.Item>List</Breadcrumb.Item>
            <Breadcrumb.Item>App</Breadcrumb.Item>
          </Breadcrumb>
          <Content
            className="site-layout-background"
            style={{
              padding: 24,
              margin: 0,
              minHeight: 280,
            }}
          >
            Content
          </Content>
          <Footer style={{ textAlign: 'center' }}>
            Ant Design ©2018 Created by Ant UED
          </Footer> */}
        </Layout>
      </Layout>
    </StyledLayout>
  )
}

const StyledLayout = styled(Layout)`
  #components-layout-demo-top-side-2 .logo {
    width: 120px;
    height: 31px;
    background: rgba(255, 255, 255, 0.2);
    margin: 16px 28px 16px 0;
    float: left;
  }

  .site-layout-background {
    background: #fff;
  }
`

const Logo = styled.div`
  width: 160px;
  height: 31px;
  background: rgba(255, 255, 255, 0.2);
  margin: 16px 28px 16px 0;
  float: left;
`

// const LinkMenu = styled(Menu.Item)``

export default DashBoard
