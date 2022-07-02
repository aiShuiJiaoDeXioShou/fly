import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/element.js'
import 'bulma/css/bulma.css'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
// 引入electron组件
import electronPlugin from './plugins/electron'
Vue.use(electronPlugin)

Vue.use(mavonEditor)
Vue.config.productionTip = false
// 引入animate.css 库
import 'animate.css';
// 这个是背景特效
import VueParticles from '@/lib/vue-particles'
Vue.use(VueParticles)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
