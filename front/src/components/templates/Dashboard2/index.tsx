import React from 'react'
import Link from 'next/link'
import { Layout, Breadcrumb, Button } from 'antd'
import { ProjectOutlined } from '@ant-design/icons'
import styled from 'styled-components'

const { Header, Content } = Layout

interface Dashboard2Props {
  MainComponent: React.ReactNode
  FooterComponent: React.ReactNode
}

const Dashboard2: React.FC<Dashboard2Props> = ({
  MainComponent,
  FooterComponent,
}) => {
  return (
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
          <LayoutContent>{MainComponent}</LayoutContent>
        </Content>
        {FooterComponent}
      </Layout>
    </Layout>
  )
}

const StyledHeader = styled(Header)`
  background: #036aa7;
`
const LayoutContent = styled.div`
  background: #fff;
  padding: 24px;
  min-height: 280px;
`
const StyledButton = styled(Button)`
  font-weight: bold;
`

export default Dashboard2
