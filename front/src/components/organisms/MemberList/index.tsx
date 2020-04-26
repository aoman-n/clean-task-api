import React, { useState, useCallback } from 'react'
import { List, Avatar, Layout, Button, Modal } from 'antd'
import { PlusOutlined } from '@ant-design/icons'
import styled from 'styled-components'

const { Content, Header } = Layout

const data = [
  {
    loginName: 'aohiro01',
    displayName: 'hiroaki aoba',
    role: 'admin',
  },
  {
    loginName: 'aohiro01',
    displayName: 'hiroaki aoba',
    role: 'write',
  },
  {
    loginName: 'aohiro01',
    displayName: 'hiroaki aoba',
    role: 'write',
  },
  {
    loginName: 'aohiro01',
    displayName: 'hiroaki aoba',
    role: 'write',
  },
]

const MemberList: React.FC = () => {
  const [visble, setVisble] = useState(false)

  const showModal = useCallback(() => {
    setVisble(true)
  }, [])

  const hideModal = useCallback(() => {
    setVisble(false)
  }, [])

  return (
    <>
      <Modal
        title="メンバー追加"
        visible={visble}
        onOk={hideModal}
        onCancel={hideModal}
        okText="ok"
        cancelText="キャンセル"
      >
        <p>Bla bla ...</p>
        <p>Bla bla ...</p>
        <p>Bla bla ...</p>
      </Modal>
      <Header
        // className="site-layout-background"
        style={{ padding: '0 24px 0 24px' }}
      >
        <StyledHeader>
          <h3>メンバー</h3>
          <StyledButton ghost icon={<PlusOutlined />} onClick={showModal}>
            追加
          </StyledButton>
        </StyledHeader>
      </Header>
      <Content
        className="site-layout-background"
        style={{
          padding: 24,
          flex: 'none',
          // margin: 24,
          // minHeight: 280,
        }}
      >
        <List
          // header={<div>Header</div>}
          // footer={<div>Footer</div>}
          // itemLayout="horizontal"
          dataSource={data}
          renderItem={item => (
            <List.Item
              actions={[
                <a key="list-loadmore-edit">edit</a>,
                <a key="list-loadmore-more">more</a>,
              ]}
            >
              <List.Item.Meta
                avatar={
                  <Avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />
                }
                title={<a href="https://ant.design">@{item.loginName}</a>}
                description={item.displayName}
              />
              <div>{item.role}</div>
            </List.Item>
          )}
        />
      </Content>
    </>
  )
}

const StyledHeader = styled.div`
  height: 100%;
  display: flex;
  align-items: center;
`
const StyledButton = styled(Button)`
  margin-left: auto;
`

export default MemberList
