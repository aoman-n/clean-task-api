import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useRouter } from 'next/router'
import { Tag, Spin } from 'antd'
import { Tag as ITag } from '~/services/model'
import { fetchTags } from '~/services/api/projects'
import { RootState } from '~/modules/rootState'
import { setTags } from '~/modules/project'

const useFetchTags = () => {
  const [isFetching, setIsFetching] = useState(false)
  const dispatch = useDispatch()
  const { query } = useRouter()
  const projectId = String(query.id)
  const { jwt } = useSelector((state: RootState) => state.auth)

  useEffect(() => {
    setIsFetching(true)
    const fetchData = async () => {
      try {
        const res = await fetchTags(projectId, jwt)
        setIsFetching(false)
        dispatch(setTags(res))
      } catch (e) {
        console.log('tag fetch error: ', e)
      }
    }
    fetchData()
  }, [jwt, projectId, dispatch])

  return { isFetching }
}

const useSelectTags = () => {
  const { tags } = useSelector((state: RootState) => state.project)

  return { tags }
}

// Container Component
const TagListContainer: React.FC = () => {
  const { isFetching } = useFetchTags()
  const { tags } = useSelectTags()

  return <TagList tags={tags} isFetching={isFetching} />
}

interface TagListModalProps {
  tags: ITag[]
  isFetching: boolean
}

// Presentation Component
const TagList: React.FC<TagListModalProps> = ({ tags, isFetching }) => {
  console.log({ isFetching })

  return (
    <div>
      {isFetching ? (
        <Spin />
      ) : (
        tags.map((tag) => (
          <Tag color={tag.color} key={tag.id}>
            {tag.name}
          </Tag>
        ))
      )}
    </div>
  )
}

export default TagListContainer
