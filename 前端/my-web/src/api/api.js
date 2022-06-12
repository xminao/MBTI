//所有请求接口

const server = "http://43.138.137.222"
//const server = "http://localhost"
const urls = class{
    static m(){

        const user = `${server}:8882`
        const register = `${user}/user/register`
        const login = `${user}/user/login`
        const getuserinfo = `${user}/user/getuserinfo`
        const getuserlist = `${user}/user/getuserlist`
        const verifytoken = `${user}/user/verifytoken`

        const question = `${server}:8887`
        const getquestion = `${question}/question/get`
        const getquestionidlist = `${question}/question/getidlist`
        const editquestion = `${question}/question/edit`
        const deletequestion = `${question}/question/delete`

        const university = `${server}:8889`
        const getcollegelist =  `${university}/university/getcollegelist`
        const getyearlist =  `${university}/university/getyearlist`
        const getmajorlist =  `${university}/university/getmajorlist`
        const getclasslist =  `${university}/university/getclasslist`
        const getstudentlist = `${university}/university/getstudentlist`
        const getstudentinfo = `${university}/university/getstudentinfo`

        const data = `${server}:8886`
        const addtestdata =  `${data}/data/add`
        const getdatalist = `${data}/data/getlist`
        const getlatestdata = `${data}/data/get`
        return {
            register,
            login,
            getuserinfo,
            verifytoken,
            getuserlist,
            getquestion,
            getquestionidlist,
            deletequestion,
            getcollegelist,
            getyearlist,
            getmajorlist,
            getclasslist,
            getstudentlist,
            editquestion,
            addtestdata,
            getlatestdata,
            getdatalist,
            getstudentinfo,
        }
    }

}
export default urls

