import { createRouter, createWebHistory } from 'vue-router'
//import HomeView from '../views/HomeView.vue'
import Main from '../views/Main.vue'
import Home from '../views/Home.vue'
import Question from '../views/Management/Question.vue'
import User from '../views/Management/User.vue'
import Management from '../views/Management/Management.vue'
import Test from '../views/Test.vue'
import Result from '../views/Result.vue'
import University from '../views/Management/University.vue'
import Data from '../views/Management/Data.vue'
import Pie from '../views/Management/Pie.vue'
import { ElMessage } from 'element-plus';
import { encrypt, decrypt } from '@/utils/jsencrypt'

const routes = [
  {
    path: '/',
    name: 'main',
    component: Main,
    redirect: '/home',
    children: [{ path: '/home', name: 'home', component: Home, meta: { requiresAuth: false}}, 
               //{ path: '/question', component: Question, meta: { requiresAuth: true}}, 
               { path: '/test', name: 'test',component: Test, meta: { requiresAuth: true}},
               { path: '/management', 
                 component: Management, 
                 meta: { requiresAuth: true,
                         roles:['admin']},
                 children: [
                  { path: '/question', component: Question, meta: { requiresAuth: true}},
                  { path: '/user', component: User, meta: { requiresAuth: true}},
                  { path: '/pie', component: Pie, meta: { requiresAuth: true}},
                  { path: '/data', component: Data, meta: { requiresAuth: true}},
                  { path: '/university', component: University, meta: { requiresAuth: true}},
                 ]},
               { path: '/result', name: 'result', component: Result, meta: { requiresAuth: true}}, 
               ], 
  },
  // {
  //   path: '/management',
  //   component: Management,
  //   children: [
  //     { path: '/question', component: Question, meta: { requiresAuth: true}},
  //   ]
  // }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})


//路由守卫
router.beforeEach((to, from, next) => {
  if (to.path == '/management') {
    console.log(localStorage.getItem('user'))
    if (decrypt(localStorage.getItem('user')) != 'admin') {
      ElMessage({
        message: "没有权限进行访问",
        type: 'warning',
        })
      router.push('/home')
    } else {
      next()
    }
  }
  if (to.meta.requiresAuth == true) {
    if (localStorage.getItem('token')) {
      next()
    } else {
      ElMessage({
        message: "请登录后操作",
        type: 'warning',
        })
      router.push('/home')
    }
  } else {
    // 放过去
    next()
  }
})

export default router
