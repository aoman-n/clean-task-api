import { NextPage } from 'next'
import Link from 'next/link'
import { AppContext } from '~/components/AppContext'
import { setTasks, Task } from '~/modules/projectModule'
import { RootState } from '~/modules/rootState'

const Index: NextPage = () => {
  return (
    <div>
      <h3>Index Page.</h3>
      <Link href="/sample">
        <a>Sampleへ</a>
      </Link>
    </div>
  )
}

Index.getInitialProps = async (ctx: AppContext) => {
  const { store, auth } = ctx
  console.log('Index Page内 store: ', store)
  console.log('Index page getInitialProps auth: ', auth)

  const tasks: Task[] = [
    {
      id: 1,
      title: 'task1',
      description: 'task1 description',
    },
    {
      id: 2,
      title: 'task2',
      description: 'task2 description',
    },
    {
      id: 3,
      title: 'task3',
      description: 'task3 description',
    },
  ]

  store.dispatch(setTasks(tasks))

  const state: RootState = store.getState()
  console.log('Index Page内 state.project.list: ', state.project.list)

  return {}
}

export default Index
