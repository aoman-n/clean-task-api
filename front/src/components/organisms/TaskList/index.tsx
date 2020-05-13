import React from 'react'
import { useSelector } from 'react-redux'
import { List, Tag } from 'antd'
import AddTask from './AddTask'
import { Task } from '~/services/model'
import { RootState } from '~/modules/rootState'

const TaskListContainer: React.FC = () => {
  const { tasks } = useSelector((state: RootState) => state.task)

  return <TaskList tasks={tasks} />
}

const TaskList: React.FC<{ tasks: Task[] }> = ({ tasks }) => {
  return (
    <List
      itemLayout="horizontal"
      header={<AddTask />}
      dataSource={tasks}
      renderItem={(item) => (
        <List.Item>
          <List.Item.Meta
            title={<p>{item.name}</p>}
            description={
              <>
                <Tag color="magenta">magenta</Tag>
                <Tag color="red">red</Tag>
              </>
            }
          />
        </List.Item>
      )}
    />
  )
}

export default TaskListContainer
