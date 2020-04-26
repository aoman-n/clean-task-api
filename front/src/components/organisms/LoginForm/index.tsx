import React from 'react'
import { Form, Input, Button, Checkbox } from 'antd'
import { UserOutlined, LockOutlined } from '@ant-design/icons'
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
  const handleOnSubmit = (values: any) => {
    console.log('handleOnSubmit value: ', values)
  }

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
          prefix={<UserOutlined className="site-form-item-icon" />}
          placeholder="LoginName"
          size="large"
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
          prefix={<LockOutlined className="site-form-item-icon" />}
          type="password"
          placeholder="Password"
          size="large"
        />
      </Form.Item>
      {/* <Form.Item>
        <Form.Item name="remember" valuePropName="checked" noStyle>
          <Checkbox>Remember me</Checkbox>
        </Form.Item>

        <a className="login-form-forgot" href="">
          Forgot password
        </a>
      </Form.Item> */}

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
