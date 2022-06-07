import axios from 'axios'
import { ElMessage } from 'element-plus'
//创建axios通用配置
//token失效问题
let instance = axios.create({
    responseType:"json",
    timeout: 1000,
    headers:{
        'Authorization':localStorage.getItem('token'),
        'Content-Type':'application/json',
    },
})

//tp拦截，在axios发出请求之前每一个接口携带token去后端校验身份
// instance.interceptors.request.use(
//     config =>{
//         console.log(config.headers)
//         let token = localStorage.getItem('token')                   // 每次发送请求之前判断是否存在token，如果存在，则统⼀在http请求的header都加上token
//         if (token) {
//             config.headers.Authorization = 'Bearer ' + token
//         }
//         return config
//     },
//     err =>{
//         return Promise.reject(err)
//     }
// )

//tp拦截，在axios发出请求之后
instance.interceptors.response.use(
    response =>{
        return response.data
    },
    error =>{
        console.log(error.response)
        if(error.response){
            let errcode = error.response.data.code
            let msg = error.response.data.msg

            switch (errcode) {
                case 100007:
                    ElMessage({
                        message: msg,
                        type: 'warning',
                    })
                    break;
            }
        }
        return Promise.reject(error.response)
    }
)

export default instance