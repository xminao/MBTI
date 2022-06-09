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
import { ElMessageBox } from 'element-plus'

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
    if (localStorage.getItem('user') != 'xminao') {
      ElMessageBox.alert('没有权限访问，请登录','出错啦',{
        confirmButtonText: '好的',
        type:"error",
        center: true,
      })
      router.push('/home')
    }
  }
  if (to.meta.requiresAuth == true) {
    if (localStorage.getItem('token')) {
      next()
    } else {
      ElMessageBox.alert('没有权限访问，请登录','出错啦',{
        confirmButtonText: '好的',
        type:"error",
        center: true,
      })
      router.push('/home')
    }
  } else {
    // 放过去
    next()
  }
})

export default router
