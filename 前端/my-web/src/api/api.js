//所有请求接口

const urls = class{
    static m(){
        const user = "http://localhost:8888"
        const register = `${user}/user/register`
        const login = `${user}/user/login`
        const getuserinfo = `${user}/user/getuserinfo`
        const getuserlist = `${user}/user/getuserlist`

        const question = "http://localhost:8887"
        const getquestion = `${question}/question/get`
        const getquestionidlist = `${question}/question/getidlist`

        const university = "http://localhost:8889"
        const getcollegelist =  `${university}/university/getcollegelist`
        const getyearlist =  `${university}/university/getyearlist`
        const getmajorlist =  `${university}/university/getmajorlist`
        const getclasslist =  `${university}/university/getclasslist`
        return {
            register,
            login,
            getuserinfo,
            getuserlist,
            getquestion,
            getquestionidlist,
            getcollegelist,
            getyearlist,
            getmajorlist,
            getclasslist,
        }
    }

}
export default urls

