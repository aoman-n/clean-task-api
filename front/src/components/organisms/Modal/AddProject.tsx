import React from 'react'
import { Modal } from 'antd'

interface AddProjectModal {
  visible: boolean
  hideModal: () => void
  submit: (e: React.MouseEvent<HTMLElement>) => void
}

const AddProject: React.FC<AddProjectModal> = ({
  visible,
  hideModal,
  submit,
}) => {
  return (
    <Modal
      title="プロジェクト作成"
      visible={visible}
      onCancel={hideModal}
      okText="Create"
      onOk={submit}
      cancelText="Cancel"
    >
      <p>Bla bla ...</p>
      <p>Bla bla ...</p>
      <p>Bla bla ...</p>
    </Modal>
  )
}

export default AddProject
