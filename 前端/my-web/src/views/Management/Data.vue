<template>
<el-container>
  <el-table :data="result.table" height="100%" stripe style="width: 100%; padding-left: 15px; padding-right: 15px;" :table-layout="fixed">
    <el-table-column fixed type="index" :index="indexMethod" />
    <el-table-column id="id" prop="username" sortable label="用户名"/>
    <el-table-column prop="type" label="测试结果" />
    <el-table-column prop="time" label="测试时间" />
  </el-table>
</el-container>
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
            const listres = await new proxy.$request(proxy.$urls.m().getdatalist).get()
            const datas = listres.data_list
            console.log(datas)
            for (let i=0; i<datas.length; i++) {
                result.table.push(datas[i])
            }
        }

        func()
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
.el-container {
    width: 100%;
}
</style>

