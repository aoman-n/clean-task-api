import React, { useState, useCallback } from 'react'
import Link from 'next/link'
import { PageHeader, Button, Descriptions, List, Avatar } from 'antd'
import { PlusOutlined } from '@ant-design/icons'
import styled from 'styled-components'
import { AddProject } from '~/components/organisms/Modal'
import { Project } from '~/services/model'

interface ProjectListProps {
  projects: Project[]
}

const ProjectList: React.FC<ProjectListProps> = ({ projects }) => {
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
        dataSource={projects}
        renderItem={(item) => (
          <List.Item style={{ paddingLeft: '24px', paddingRight: '24px' }}>
            {/* <Link href={`/dashboard/${item.id}/tasks`}> */}
            <Link href={`/dashboard/${item.id}/${item.title}`}>
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
