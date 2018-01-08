import Vue from 'vue'
import Router from 'vue-router'

import ItemListView from '../views/ItemList.vue'

Vue.use(Router)


export function createRouter () {
  return new Router({
    mode: 'history',
    fallback: false,
    scrollBehavior: () => ({ y: 0 }),
    routes: [
      { path: '/news/:page(\\d+)?', component: ItemListView },
      { path: '/', redirect: '/news' }
    ]
  })
}