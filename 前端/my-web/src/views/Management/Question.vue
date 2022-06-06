<template>
  <el-table :data="result.table" height="100%" stripe style="width: 100%; padding-left: 15px; padding-right: 15px;" :table-layout="fixed">
  <el-table-column fixed id="id" prop="question_id" sortable label="题目序号" width="150"/>
    <el-table-column prop="question_desc" label="题目描述" />
    <el-table-column prop="option_a_desc" label="选项A" />
    <el-table-column prop="option_a_target" sortable  label="倾向" width="100"/>
    <el-table-column prop="option_b_desc" label="选项B" />
    <el-table-column prop="option_b_target" sortable  label="倾向" width="100"/>
  </el-table>
</template>

<script>
import { getCurrentInstance, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'

export default ({
    setup() {
        const router = useRouter()
        const {proxy} = getCurrentInstance()

        let result = reactive({
            table: [],
        })

        console.log(result.table)

        const func=async()=> {
            const listres = await new proxy.$request(proxy.$urls.m().getquestionidlist).get()
            const idList = listres.question_id_list
            for (let i=0; i<idList.length; i++) {
                const obj = {"id":idList[i]}
                const res = await new proxy.$request(proxy.$urls.m().getquestion, obj).get()
                //草啊，不用赋值啊  直接push就不报错了 我日
                result.table.push(res.question_info)
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
            result,
            goHome,
        }
    },
})
</script>

<style scoped lang="less">
</style>

