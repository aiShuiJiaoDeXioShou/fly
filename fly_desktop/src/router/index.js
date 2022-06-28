import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/home/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    // path: '/home',
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
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/home/About.vue')
  },
  {
    path: '/home_test',
    // path: '/',
    name: 'HomeTest',
    component: () => import('../views/template_test/HomeTest.vue')
  }

]

const router = new VueRouter({
  routes
})

export default router
