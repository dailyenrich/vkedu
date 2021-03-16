import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/auth',
    component: () => import('../components/layout/UserLayout.vue'),
    children: [
      {
        path: 'login',
        component: () => import(/* webpackChunkName: "login" */ '../views/Login.vue')
      }
    ]
  },
  {
    path: '/',
    name: 'Home',
    component: () => import('../components/layout/Layout.vue'),
    children: [
      {
        path: '/',
        name: 'Index',
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
      },
      {
        path: 'dashbord',
        name: 'Dashbord',
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
      }
    ]
  },
]

let router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: routes
})

export default router
