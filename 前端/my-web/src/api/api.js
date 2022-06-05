//所有请求接口

const urls = class{
    static m(){
        const user = "http://localhost:8888"
        const register = `${user}/user/register`
        const login = `${user}/user/login`
        const getuserinfo = `${user}/user/getuserinfo`

        const question = "http://localhost:8887"
        const getquestion = `${question}/question/get`
        const getquestionidlist = `${question}/question/getidlist`
        return {
            register,
            login,
            getuserinfo,
            getquestion,
            getquestionidlist,
        }
    }

}
export default urls

