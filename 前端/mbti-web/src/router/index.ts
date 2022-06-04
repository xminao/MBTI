import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Test from '../components/Test.vue'
import Type from '../components/Type.vue'
import Login from '../components/Login.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
    redirect: '/test',
    children: [{ path: '/test', component: Test}, 
               { path: '/type', component: Type},
              { path: '/login', component: Login}]
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})


// router.beforeEach((to) => {
//   // 根据token判断是否登录
//   let token = localStorage.getItem('Token');
//   // 有token但是访问的是登录页 => 强制去首页
//   if (token && to.path === '/test')
//       return "/";
//   // 没有token但是访问的是首页 => 强制去登录页
//   else if (!token && to.path !== '/login')
//       return "/";
// });

export default router
