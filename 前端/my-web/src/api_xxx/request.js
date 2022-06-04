import axios from "axios";
import qs from "qs";
 
/**
 * 封装请求方式
 */
const request =
{
    /**
     * @name: 封装axios get方法
     * @desc: 描述
     * @param url 请求连接
     * @param params 请求参数
     */
    get(url, params) 
    {
        return new Promise((resolve, reject) => {
            axios
                .get(url, {
                    params: params
                })
                .then(res => {
                    resolve(res.data);
                })
                .catch(err => {
                    reject(err);
                });
        });
    },
 
    /**
     * @name: 封装axios post方法
     * @desc: 描述
     * @param url 请求连接
     * @param params 请求参数
     */
    post(url, params) 
    {
        return new Promise((resolve, reject) => {
            axios
                .post(url, qs.stringify(params))
                .then(res => {
                    resolve(res.data);
                })
                .catch(err => {
                    reject(err);
                });
        });
    },
}
export default request;
