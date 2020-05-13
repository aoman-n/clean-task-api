import { NextPage, ExNextPageContext } from 'next'
import Entrance from '~/components/templates/Entrance'
import ProjectList from '~/components/organisms/ProjectList'
import { fetchProjects } from '~/services/api/projects'
import { redirect } from '~/routes'
import { Project } from '~/services/model'

const ProjectsPage: NextPage<{ projects: Project[] }> = ({ projects }) => {
  return <Entrance content={<ProjectList projects={projects} />} />
}

ProjectsPage.getInitialProps = async (ctx: ExNextPageContext) => {
  if (!ctx.auth.token) {
    redirect(ctx.res, '/login')
    return { projects: [] }
  }

  try {
    const projects = await fetchProjects(ctx.auth.token)
    return { projects }
  } catch (e) {
    console.log('fetchProjects error: ', e)
    return { projects: [] as Project[] }
  }
}

export default ProjectsPage
