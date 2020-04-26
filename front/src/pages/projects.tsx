import { NextPage, ExNextPageContext } from 'next'
import ProjectList from '~/components/organisms/ProjectList'

const IndexPage: NextPage = () => {
  return <ProjectList />
}

IndexPage.getInitialProps = async (ctx: ExNextPageContext) => {
  // 必要な情報をfetchする
  console.log('ctx: ', ctx)
}

export default IndexPage
