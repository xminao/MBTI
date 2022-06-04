import axios from 'axios'
import { Base64 } from 'js-base64'
import { ElMessageBox } from 'element-plus'
//创建axios通用配置
let instance = axios.create({
    responseType:"json",
    headers:{
        'Content-Type':'application/json',
        // 'Access-Control-Allow-Origin': '*',
        // 'Access-Control-Allow-Methods':'get, post, put'
    },


})
//对token进行加密
function baseFun(){ 
    const token = localStorage.getItem('token')                     //从本地缓存里取出token
    const base64 = Base64.encode(token + ':')                         //对token进行加密
    return 'Basic ' + base64
}
//tp拦截，在axios发出请求之前每一个接口携带token去后端校验身份
instance.interceptors.request.use(
    config =>{
        console.log(config.headers)
        let token = localStorage.getItem('token')                   // 每次发送请求之前判断是否存在token，如果存在，则统⼀在http请求的header都加上token
        if (token) {
            config.headers.Authorization = baseFun()
        }
 
        return config
    },
    err =>{
        return Promise.reject(err)
    }
)
//tp拦截，在axios发出请求之后
instance.interceptors.response.use(
    response =>{
        return response
    },
    error =>{
        if(error.response){
            let ERRS = error.response.status                                            //这个是状态码
            let MSG = error.response.data.msg.msg
            let errdata = ERRS == 401 ? MSG :'服务器发生错误'                           //判断返回的是不是等于401，如果是，返回服务器发生错误
            switch(ERRS){
                case 401:                                                               //没权限
                    ElMessageBox.alert(errdata,'提示',{
                        confirmButtonText: '好的',
                        type:"error"
                    })
                    // .then(res=>{
                        
                    // })
                    // .catch(err=>{

                    // })
                    break;
            }
        }
        return Promise.reject(error.response.data)
    }
)
export default instance