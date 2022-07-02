import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/home/Home.vue'
import store from '../store'

Vue.use(VueRouter)

const routes = [
  {
    path: '/home',
    name: '主页',
    component: Home,
    children: [
      {
        path: '/',
        name: '起点排行',
        component: () => import('../views/author/BookSort.vue')
      },
      {
        path: '/note_edit',
        name: '编辑器',
        component: () => import('../views/author/NoteEdit.vue')
      }
    ]
  },
  {
    path: '/',
    name: '首页',
    component: () => import('../views/home/Index.vue'),
    children: [

    ]
  },
  {
    path: '/login',
    name: '登入 or 注册',
    component: () => import('../views/user/Login.vue')
  },
]

const router = new VueRouter({
  routes
})

// 创建一个路由守卫
router.beforeEach((to, form, next) => {
  // 捕获navigate(to, next)的运行异常
  navigate(to, form, next)
})

// 导航
function navigate(to, form, next) {

  console.log('to', to);
  // 当前路由对象
  let self = getRouterPaths(to)

  // 获取已经进入的路由
  let paths = store.state.router_paths

  // 判断该属性是否跟to.path相同,是否已经包含了该路由
  for (let i = 0; i < paths.length; i++) {
    // 如果包含了该路由，则不添加路由组
    if (paths[i].path === self.path) {
      // 将该路由后面的路由全部删除
      paths.splice(i + 1, paths.length - i - 1)
      next()
      return
    }
  }

  // 获取当前路由的父类
  let parent = self.parent

  // 如果不为空比较form.parent是否与to.parent相同
  // 判断将要去往的路由父类是否为空
  let form_self = getRouterPaths(form)
  if (parent && form_self)
    if (form_self.parent)
      if (form_self.parent.path === parent.path) {
        let i = getPathIndex(paths, form_self.parent.path)
        // 删除包括i的所有路由
        paths.splice(i, paths.length - i - 1)
      }

  // 判断父类是否为空,且paths是否存在
  if (parent && paths.filter(item => item.path === parent.path).length === 0) {
    // 如果不为空,且paths不存在该路由，则将父类添加到路由组中
    store.dispatch('RouterAdmin', parent)
  }

  store.dispatch('RouterAdmin', self)
  next();
}

// 获取路由纯洁版
function getRouterPaths(to) {
  // 获取匹配的路由对象
  let matcheds = to.matched
  // 获取当前的路由对象matched的最后一个参数
  let last = matcheds[matcheds.length - 1]
  // 返回当前对象
  return last
}

// 获取指定path所在的坐标
function getPathIndex(paths, path) {
  let index = -1
  for (let i = 0; i < paths.length; i++) {
    // 如果包含了该路由，则不添加路由组
    if (paths[i].path === path) {
      index = i
      return
    }
  }

  return index
}

export default router
