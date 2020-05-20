import { NextPage, ExNextPageContext } from 'next'
import Entrance from '~/components/templates/Entrance'
import ProjectList from '~/components/organisms/ProjectList'
import { fetchProjects } from '~/services/api/projects'
import { redirect } from '~/routes'
import { Project } from '~/services/model'
import { setProjects } from '~/modules/project'

const ProjectsPage: NextPage<{ projects: Project[] }> = ({ projects }) => {
  return <Entrance content={<ProjectList projects={projects} />} />
}

ProjectsPage.getInitialProps = async (ctx: ExNextPageContext) => {
  if (!ctx.auth.jwt) {
    redirect(ctx.res, '/login')
    return { projects: [] }
  }

  try {
    const projects = await fetchProjects(ctx.auth.jwt)
    ctx.store.dispatch(setProjects(projects))
    return { projects }
  } catch (e) {
    console.log('fetchProjects error: ', e)
    redirect(ctx.res, '/login')
    return { projects: [] as Project[] }
  }
}

export default ProjectsPage
