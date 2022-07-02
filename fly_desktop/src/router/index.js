import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/home/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    // path: '/',
    path: '/home',
    name: 'Home',
    component: Home,
    children:[
      {
        path: '/book_sort',
        name: 'BookSort',
        component: () => import('../views/author/BookSort.vue')
      },
      {
        path: '/note_edit',
        name: 'NoteEdit',
        component: () => import('../views/author/NoteEdit.vue')
      }
    ]
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('../views/home/About.vue')
  },
  {
    path: '/home_test',
    // path: '/',
    name: 'HomeTest',
    component: () => import('../views/template_test/HomeTest.vue')
  },
  {
    path: '/',
    name: 'Index',
    component: () => import('../views/home/Index.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/user/Login.vue')
  },
  {
    path: '/raindrop',
    name: 'Raindrop',
    component: () => import('../components/bg/Raindrop.vue')
  }

]

const router = new VueRouter({
  routes
})

export default router
