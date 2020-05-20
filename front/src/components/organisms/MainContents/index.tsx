import React, { useState } from 'react'
import { Card } from 'antd'
import TaskList from '~/components/organisms/TaskList'
import TagList from '~/components/organisms/TagList'

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

const contentList = {
  task: <TaskList />,
  tag: <TagList />,
  member: <p>Members</p>,
}

type TabKey = keyof typeof contentList

const Component: React.FC = () => {
  const [keyState, setKeyState] = useState<{ key: TabKey }>({ key: 'task' })
  const onTagChange = (key: string) => {
    setKeyState({ key: key as TabKey })
  }

  return (
    <Card
      tabList={tabList}
      activeTabKey={keyState.key}
      onTabChange={(key) => onTagChange(key)}
      style={{ width: '100%' }}
    >
      {contentList[keyState.key]}
    </Card>
  )
}

export default Component
