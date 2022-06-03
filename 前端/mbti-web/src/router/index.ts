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

export default router
