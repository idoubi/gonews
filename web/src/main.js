import Vue from 'vue'
import App from './App.vue'
import { createStore } from './store'
import { createRouter } from './router'

const store = createStore()
const router = createRouter()

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
