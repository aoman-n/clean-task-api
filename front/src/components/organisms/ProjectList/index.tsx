import React, { useState, useCallback } from 'react'
import Link from 'next/link'
import { AddProject } from '~/components/organisms/Modal'
import { PageHeader, Button, Descriptions, List, Avatar } from 'antd'
import { PlusOutlined } from '@ant-design/icons'
import styled from 'styled-components'

const data = [
  {
    id: 1,
    title: 'プロジェクト1',
    description: 'Description 1111111111111',
  },
  {
    id: 2,
    title: 'プロジェクト222',
    description: 'Description 2222222222222',
  },
  {
    id: 3,
    title: 'プロジェクト3',
    description: 'Description 3333333333333333',
  },
  {
    id: 4,
    title: 'プロジェクト4',
    description: 'Description 444444444444444',
  },
]

const ProjectList: React.FC = () => {
  const [visible, setVisible] = useState(false)

  const showModal = useCallback(() => {
    setVisible(true)
  }, [])

  const hideModal = useCallback(() => {
    setVisible(false)
  }, [])

  return (
    <Component>
      <AddProject visible={visible} hideModal={hideModal} submit={() => {}} />
      <PageHeader
        ghost={false}
        title="Project"
        extra={[
          <Button
            key="1"
            type="primary"
            icon={<PlusOutlined />}
            onClick={showModal}
          >
            新規作成
          </Button>,
        ]}
      ></PageHeader>
      <List
        itemLayout="horizontal"
        dataSource={data}
        renderItem={item => (
          <List.Item style={{ paddingLeft: '24px', paddingRight: '24px' }}>
            <Link href={`/dashboard/${item.id}/tasks`}>
              <a>
                <List.Item.Meta
                  title={item.title}
                  description={item.description}
                />
              </a>
            </Link>
          </List.Item>
        )}
      />
    </Component>
  )
}

const Component = styled.div`
  background-color: #f5f5f5;
  padding: 24px;
`

export default ProjectList
