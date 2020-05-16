import React, { useState, useCallback } from 'react'
import Link from 'next/link'
import { PageHeader, Button, List } from 'antd'
import { PlusOutlined } from '@ant-design/icons'
import styled from 'styled-components'
import { AddProject } from '~/components/organisms/Modal'
import { Project } from '~/services/model'
import { fetchHello, fetchCookie } from '~/utils/api'

interface ProjectListProps {
  projects: Project[]
}

const useHello = () => {
  const handleHello = async () => {
    try {
      const res = await fetchHello()

      console.log('res: ', res)
    } catch (e) {
      console.log('err: ', e)
    }
  }

  return { handleHello }
}

const useCookie = () => {
  const handleCookie = async () => {
    try {
      const res = await fetchCookie()

      console.log('res: ', res)
    } catch (e) {
      console.log('err: ', e)
    }
  }

  return { handleCookie }
}

const ProjectList: React.FC<ProjectListProps> = ({ projects }) => {
  const [visible, setVisible] = useState(false)
  const { handleHello } = useHello()
  const { handleCookie } = useCookie()

  const showModal = useCallback(() => {
    setVisible(true)
  }, [])

  const hideModal = useCallback(() => {
    setVisible(false)
  }, [])

  return (
    <Component>
      <Button onClick={handleHello}>Cookie Sample</Button>
      <Button onClick={handleCookie}>BFF へ</Button>
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
