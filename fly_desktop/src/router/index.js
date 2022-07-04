import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/home/Home.vue'
import navigate from './navigate'

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
      },
      {
        path: '/read_write',
        name: '文件读写',
        component: () => import('../views/setting/ReadWrite.vue')
      },
      {
        path: '/class_box',
        name: '分类管理',
        component: () => import('../views/kp_core/ClassBox.vue')
      },
      {
        path: '/class_stl',
        name: '分类模板',
        component: () => import('../views/kp_core/ClassSTL.vue')
      },
      {
        path: '/class_export',
        name: '分类导出',
        component: () => import('../views/kp_core/ClassExport.vue')
      },
      {
        path: '/class_import',
        name: '分类导入',
        component: () => import('../views/kp_core/ClassImport.vue')
      },
      {
        path: '/qq_assistant',
        name: 'QQ助手',
        component: () => import('../views/kp_core/QQAssistant.vue')
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

export default router
