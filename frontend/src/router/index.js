import Vue from 'vue'
import Router from 'vue-router'
import blogRouter from "./Modules/blog";
import adminRouter from "./Modules/admin";


Vue.use(Router)

// import blogRouter from './Modules/blog'
export const constantRoutes = [
  {
    path: '/',
    component: () => import('@/views/blog/Layout'),
    hidden: true,
    name: 'app',
    meta: { title: 'vaniot' }
  },
  // {
  //   path: '/404',
  //   component: () => import('@/views/error-page/404'),
  //   hidden: true
  // },
  // {
  //   path: '/401',
  //   component: () => import('@/views/error-page/401'),
  //   hidden: true
  // },
  // // 404 page must be placed at the end !!!
  // { path: '*', redirect: '/404', hidden: true }
]

export const asyncRoutes = [
  blogRouter,
  adminRouter
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
