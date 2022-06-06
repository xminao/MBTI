import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import 'element-plus/dist/index.css'
import MenuTree from './components/MenuTree.vue'

const app = createApp(App)
app.use(router)
app.use(ElementPlus)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.component('menu-tree', MenuTree)

//请求地址
import urls from '@/api/api.js'
app.config.globalProperties.$urls = urls
//请求方法
import request from '@/api/request.js'
app.config.globalProperties.$request = request

app.mount('#app')
