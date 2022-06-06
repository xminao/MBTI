<template>
  <el-table :data="result.table" height="100%" stripe style="width: 100%; padding-left: 15px; padding-right: 15px;" :table-layout="fixed">
    <el-table-column fixed type="index" :index="indexMethod" />
    <el-table-column id="id" prop="username" sortable label="用户名"/>
    <el-table-column prop="nickname" label="昵称" />
    <el-table-column prop="password" label="密码" />
    <el-table-column prop="gender"  label="性别"/>
    <el-table-column prop="auth_group" label="用户组" />
    <el-table-column prop="binding_student_id" label="绑定学号" />
  </el-table>
</template>

<script>
import { getCurrentInstance, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'

export default ({
    setup() {
        const router = useRouter()
        const {proxy} = getCurrentInstance()

        const indexMethod=(index)=>{
            return index
        }

        let result = reactive({
            table: [],
        })

        const func=async()=> {
            indexMethod(0)
            const listres = await new proxy.$request(proxy.$urls.m().getuserlist).get()
            const datas = listres.userlist
            console.log(datas)
            for (let i=0; i<datas.length; i++) {
                result.table.push(datas[i])
            }
        }

        func()
        
        // //递归
        // const getQuesList=async(idlist, i)=> {
            
        //     console.log("哈哈哈哈")
        //     if (i == idlist.length) {
        //         return
        //     }

        //     const obj = {"id":idlist[i]}
        //     const res = await new proxy.$request(proxy.$urls.m().getquestion, obj).get()
        //     tableData.value = tableData.value.push(res.question_info)
        //     getQuesList(idlist, i+1)
        // }

        // //递归
        // const getQuesCount=async()=> {
        //     const listres = await new proxy.$request(proxy.$urls.m().getquestionidlist).get()
        //     const idList = listres.question_id_list
        //     return idList
        // }
        // getQuesCount().then(res=>{
        //     getQuesList(res, 0)
        // })

        const goHome=()=> {
            router.push('/home')
        }

        return {
            //getQuesList,
            //getQuesCount,
            indexMethod,
            result,
            goHome,
        }
    },
})
</script>

<style scoped lang="less">
</style>

