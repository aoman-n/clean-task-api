import React from 'react'
import { Tag } from 'antd'
import { Tag as ITag } from '~/services/model'

const useFetchTags = () => {}

interface TagListModalProps {
  tags: ITag[]
}

const TagList: React.FC<TagListModalProps> = ({ tags }) => {
  return (
    <div>
      {tags.map((tag) => (
        <Tag color={tag.color} key={tag.id}>
          {tag.name}
        </Tag>
      ))}
    </div>
  )
}

export default TagList
