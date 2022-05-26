import { createApp } from 'vue'

/* 引入ElementPlus */ 
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import App from './App.vue'

/* 通过createApp创建一个新的应用示例 */
const app = createApp(App)

app.mount('#app')
app.use(ElementPlus) //使用ElementPlus
