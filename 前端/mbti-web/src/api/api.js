//所有请求接口

const url = "http://localhost:8888"

const urls = class{
    static m(){
        const register = `${url}/user/register`
        const login = `${url}/user/login`

        return {
            register,
            login
        }
    }

}
export default urls

