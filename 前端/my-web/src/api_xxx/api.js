import http from './http.js'

// 获取验证码，这个函数指的是使用GET请求请求目标服务器的
export function getUserinfo(params) {
    return http.get("/user/getuserinfo", params)
}  
// 验证用户名\密码，这里指的是使用POST方法请求Login接口，并携带params对象作为参数。
export function Login(params) {
    return http.post("/user/login", params)
}
