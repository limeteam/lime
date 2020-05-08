import Vue from 'vue'
import 'normalize.css/normalize.css' // A modern alternative to CSS resets
import 'element-ui/lib/theme-chalk/index.css'
import '@/styles/index.scss' // global css

import Element from 'element-ui'
import App from './App'
import router from './router'
import store from './store'
import Cookies from 'js-cookie'

import i18n from './lang' // Internationalization
import './icons' // icon
import './errorLog' // error log
import './permission' // permission control

import * as VueGoogleMaps from 'vue2-google-maps'
import ECharts from 'vue-echarts' // 在 webpack 环境下指向 components/ECharts.vue

// 手动引入 ECharts 各模块来减小打包体积
import 'echarts/lib/chart/bar'
import 'echarts/lib/component/tooltip'
import 'echarts/lib/chart/line'
import 'echarts/lib/component/polar'

import * as filters from './filters' // global filters

Vue.use(Element, {
  size: Cookies.get('size') || 'medium', // 'medium', // set element-ui default size
  i18n: (key, value) => i18n.t(key, value)
})
Vue.component('v-chart', ECharts)

// 谷歌地图
Vue.use(VueGoogleMaps, {
  load: {
    key: 'AIzaSyCqsXGeGltYxHMlOGyadORfCDU-tg_b2p4',
    language: 'en',
    libraries: 'places,drawing'
  }
})

// register global utility filters.
Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key])
})

// 模拟数据 mockjs
// TODO: 开发用，发版测试需要移除
import '@/api/admin/mock'
import '@/api/lime-admin/mock'

// console.log('process.env', process.env)

Vue.config.productionTip = false
new Vue({
  el: '#app',
  router,
  store,
  i18n,
  render: h => h(App)
})
