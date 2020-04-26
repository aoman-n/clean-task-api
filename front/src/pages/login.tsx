import { NextPage } from 'next'
import Entrance from '~/components/templates/Entrance'
import LoginForm from '~/components/organisms/LoginForm'

const Login: NextPage = () => {
  return <Entrance content={<LoginForm />} />
}

export default Login
