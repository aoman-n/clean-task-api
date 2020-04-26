import React from 'react'
import styled from 'styled-components'

interface Entrance {
  content: React.ReactNode
}

const Entrance: React.FC<Entrance> = ({ content }) => {
  return (
    <Layout>
      <Content>
        {/* <Header>Task App</Header> */}
        <Body>{content}</Body>
      </Content>
    </Layout>
  )
}

const Layout = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f0f2f5;
  height: 100vh;
`
const Content = styled.div`
  width: 340px;
`
const Header = styled.div`
  padding: 0 24px;
  font-size: 24px;
  font-weight: 700;
  text-align: center;
`
const Body = styled.div`
  padding: 30px 0;
`

export default Entrance
