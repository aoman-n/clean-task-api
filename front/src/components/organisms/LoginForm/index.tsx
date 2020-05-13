import React, { useState, useCallback } from 'react'
import { useRouter } from 'next/router'
import { Form, Input, Button } from 'antd'
import { UserOutlined, LockOutlined } from '@ant-design/icons'
import { loginApi } from '~/utils/api'
import { login } from '~/auth'
// import { useForm } from 'react-hook-form'

type FormData = {
  loginName: string
  password: string
}

const initialValues: FormData = {
  loginName: '',
  password: '',
}

const LoginForm = () => {
  const router = useRouter()
  const [data, setData] = useState<FormData>(initialValues)

  const handleOnSubmit = useCallback(async () => {
    console.log('handleOnSubmit value: ', data)
    try {
      const res = await loginApi(data)
      console.log('res: ', res)
      login({ token: res.token })
      router.push('/projects')
    } catch (e) {
      console.log('エラーだよ！')
    }
  }, [data])

  const onChangeInput = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      setData({
        ...data,
        [e.target.name]: e.target.value,
      })
    },
    [data],
  )

  return (
    <Form
      name="normal_login"
      className="login-form"
      initialValues={initialValues}
      onFinish={handleOnSubmit}
    >
      <Form.Item
        name="loginName"
        rules={[
          { required: true, message: 'loginNameを入力してください' },
          { min: 4, message: '4文字以上で入力してください' },
        ]}
      >
        <Input
          name="loginName"
          prefix={<UserOutlined className="site-form-item-icon" />}
          placeholder="LoginName"
          size="large"
          onChange={onChangeInput}
        />
      </Form.Item>
      <Form.Item
        name="password"
        rules={[
          { required: true, message: 'Passwordを入力してください' },
          { min: 6, message: '6文字以上で入力してください' },
        ]}
      >
        <Input
          name="password"
          prefix={<LockOutlined className="site-form-item-icon" />}
          type="password"
          placeholder="Password"
          size="large"
          onChange={onChangeInput}
        />
      </Form.Item>

      <Form.Item>
        <Button
          type="primary"
          htmlType="submit"
          className="login-form-button"
          block
          size="large"
        >
          Log in
        </Button>
        Or <a href="">register now!</a>
      </Form.Item>
    </Form>
  )
}

export default LoginForm
