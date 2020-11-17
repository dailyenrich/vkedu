import Vue from 'vue'
import Router from 'vue-router'

import index from '~/router/index'
import course from '~/router/course'

Vue.use(Router)

export function createRouter() {
  return new Router({
    mode: 'history',
    routes: [
      index,
      ...course
    ]
  })
}
