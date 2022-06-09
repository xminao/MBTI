<template>
<el-container>
  <el-table :data="result.table" height="100%" stripe style="width: 100%; padding-left: 15px; padding-right: 15px;" :table-layout="fixed">
    <el-table-column fixed type="index" :index="indexMethod" />
    <el-table-column id="id" prop="username" sortable label="用户名" width="150"/>
    <el-table-column prop="type" label="测试结果" width="100"/>
    <el-table-column prop="time" label="测试时间" />
    <el-table-column label="其它操作">
      <template #default="scope">
        <el-button link type="primary" size="small" @click="handleClick(scope.row)"
          >测试细则</el-button>
        </template>
     </el-table-column>
  </el-table>
  <el-dialog v-model="dialogTableVisible" title="题目细则">
    <el-table :data="options.datas[0]">
      <el-table-column property="Id" label="题号"  />
      <el-table-column property="Option" label="选项" />
    </el-table>
  </el-dialog>
</el-container>
</template>

<script>
import { getCurrentInstance, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'

export default ({
    setup() {
        const router = useRouter()
        const {proxy} = getCurrentInstance()

        let dialogTableVisible = ref(false)

        const indexMethod=(index)=>{
            return index
        }

        let result = reactive({
            table: [],
        })

        let options = reactive({
            datas: [],
        })

        const func=async()=> {
            indexMethod(0)
            const listres = await new proxy.$request(proxy.$urls.m().getdatalist).get()
            const datas = listres.data_list
            for (let i=0; i<datas.length; i++) {
                result.table.push(datas[i])
            }
            console.log(result.table)
        }

        func()
        const goHome=()=> {
            router.push('/home')
        }

        const handleClick=(parm)=> {
            for (let i=0; i<result.table.length; i++) {
                if (result.table[i]['time'] == parm.time) {
                    dialogTableVisible.value = true
                    options.datas.push(JSON.parse(result.table[i]['selection']))
                }
            }
        }

        

        return {
            //getQuesList,
            //getQuesCount,
            indexMethod,
            handleClick,
            result,
            goHome,
            options,
            dialogTableVisible,
        }
    },
})
</script>

<style scoped lang="less">
.el-container {
    width: 100%;
}
</style>

