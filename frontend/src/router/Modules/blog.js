import Layout from '@/views/blog/Layout'

const blogRouter = {
  path: '',
  components: Layout,
  redirect: 'noRedirect',
  name: 'Vaniot',
  meta: {
    title: 'Charts',
    icon: 'chart'
  },
  children: [
    {
      path: '',
      component: () => import('@/views/blog/Home'),
      name: 'Home',
      meta: { title: 'Home', noCache: true }
    },
    {
      path: 'post',
      component: () => import('@/views/blog/Post'),
      name: 'Post',
      meta: { title: 'Post', noCache: true }
    },
    {
      path: 'tags',
      component: () => import('@/views/blog/Tags'),
      name: 'Tags',
      meta: { title: 'Tags', noCache: true }
    },
    {
      path: 'categories',
      component: () => import('@/views/blog/Categories'),
      name: 'Categories',
      meta: { title: 'Categories', noCache: true }
    },
    {
      path: 'timeline',
      component: () => import('@/views/blog/TimeLine'),
      name: 'TimeLine',
      meta: { title: 'TimeLine', noCache: true }
    },
  ]
}

export default blogRouter
