import { NextPage } from 'next'
import Link from 'next/link'
import { AppContext } from '~/components/AppContext'
import { setTasks } from '~/modules/task'
import { Task } from '~/services/model'

const Index: NextPage<{ tasks: Task[] }> = ({ tasks }) => {
  return (
    <div>
      <h3>Index Page.</h3>
      <Link href="/sample">
        <a>Sampleへ</a>
      </Link>
      <h4>Tasks</h4>
      <ul>
        {tasks.map((t) => (
          <li key={t.id}>
            {t.id}: {t.name}
          </li>
        ))}
      </ul>
    </div>
  )
}

Index.getInitialProps = async (ctx: AppContext) => {
  const { store, auth } = ctx
  // console.log('Index Page内 store: ', store)
  // console.log('Index page getInitialProps auth: ', auth)

  const tasks: Task[] = [
    {
      id: 1,
      name: 'task1',
      status: 1,
    },
    {
      id: 2,
      name: 'task2',
      status: 1,
    },
    {
      id: 3,
      name: 'task3',
      status: 1,
    },
  ]

  store.dispatch(setTasks(tasks))

  // const state: RootState = store.getState()
  // console.log('Index Page内 state.project.list: ', state.project.list)

  return { tasks }
}

export default Index
