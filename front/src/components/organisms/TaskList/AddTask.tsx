import React, { useState } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { useRouter } from 'next/router'
import { Modal, Form, Input, Button } from 'antd'
import { LoadingOutlined } from '@ant-design/icons'
import { useForm } from 'react-hook-form'
import { RHFInput } from 'react-hook-form-input'
import { RootState } from '~/modules/rootState'
import { addTask } from '~/modules/task'
import { postTask } from '~/services/api/tasks'

const useVisible = () => {
  const [visible, setVisible] = useState(false)

  const handleOpen = () => {
    setVisible(true)
  }

  const handleClose = () => {
    setVisible(false)
  }

  return { visible, handleOpen, handleClose }
}

interface FormData {
  name: string
}

const useCreateTask = (closeCallback: () => void) => {
  const [posting, setPosting] = useState(false)
  const { handleSubmit, register, errors, reset, control } = useForm<FormData>()
  const query = useRouter()
  const projectId = String(query.query.id)
  const { token } = useSelector((state: RootState) => state.auth)
  const dispatch = useDispatch()

  const onSubmit = async (data: FormData) => {
    console.log({ data })
    setPosting(true)
    try {
      const task = await postTask(token, projectId, data)
      console.log('作成task: ', task)
      dispatch(addTask(task))
      reset()
      closeCallback()
    } catch (e) {
      console.log('error: ', e)
    }
    setPosting(false)
  }

  return {
    handleSubmit: handleSubmit(onSubmit),
    errors,
    register,
    posting,
    control,
  }
}

const AddTask: React.FC = () => {
  const { visible, handleOpen, handleClose } = useVisible()
  const { handleSubmit, errors, posting, register } = useCreateTask(handleClose)
  console.log({ errors })

  return (
    <>
      <Button type="primary" onClick={handleOpen}>
        タスク作成
      </Button>
      <Modal
        title="タスク作成"
        visible={visible}
        onOk={handleSubmit}
        onCancel={handleClose}
        okText={
          posting ? <LoadingOutlined style={{ fontSize: 24 }} spin /> : '作成'
        }
        cancelText="キャンセル"
      >
        <Form layout="vertical">
          <Form.Item
            label="タスク名"
            hasFeedback={!!errors.name}
            validateStatus={errors.name ? 'error' : 'success'}
            help={errors.name && 'タスク名を入力してください'}
          >
            <RHFInput
              as={Input}
              rules={{ required: true }}
              name="name"
              register={register}
              defaultValue=""
            />
          </Form.Item>
        </Form>
      </Modal>
    </>
  )
}

export default AddTask
