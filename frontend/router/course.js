import list from '~/pages/course/list'
import info from '~/pages/course/info'
import detail from '~/pages/course/detail'

export default [
  {
    path: '/courses',
    component: list
  },
  {
    path: '/info',
    component: info
  },
  {
    path: '/detail',
    component: detail
  },
]
