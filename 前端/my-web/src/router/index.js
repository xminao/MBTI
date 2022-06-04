import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Test from '../components/Test.vue'
import Type from '../components/Type.vue'
import Login from '../components/Login.vue'
import { ElMessageBox } from 'element-plus'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
    redirect: '/test',
    children: [{ path: '/test', component: Test, meta: { requiresAuth: false}}, 
               { path: '/type', component: Type, meta: { requiresAuth: true}}, 
               { path: '/login', component: Login, meta: { requiresAuth: true}}], 
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})


//路由守卫
router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth == true) {
    if (localStorage.getItem('token') != undefined) {
      next()
    } else {
      ElMessageBox.alert('没有权限访问，请登录','出错啦',{
        confirmButtonText: '好的',
        type:"error",
        center: true,
      })
      router.push('/test')
    }
  } else {
    // 放过去
    next()
  }
})

export default router
