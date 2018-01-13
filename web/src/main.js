import Vue from 'vue'
import App from './App.vue'
import ProgressBar from './components/ProgressBar.vue'
import { createStore } from './store'
import { createRouter } from './router'

const bar = Vue.prototype.$bar = new Vue(ProgressBar).$mount()
document.body.appendChild(bar.$el)

const store = createStore()
const router = createRouter()

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
