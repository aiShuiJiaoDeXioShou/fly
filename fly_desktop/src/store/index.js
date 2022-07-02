import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    // 路由管理,默认显示首页
    router_paths: []
  },
  mutations: {
    ROUTER_ADMIN(state, value) {
      state.router_paths.push(value)
    }
  },
  actions: {
    RouterAdmin(ctx, value) {
      ctx.commit('ROUTER_ADMIN', value)
    }
  },
  getters: {
    GetRouterPaths(state) {
      return state.router_paths
    }
  },
  modules: {

  }
})
