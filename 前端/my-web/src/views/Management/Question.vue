<template>
  <el-radio-group v-model="tableLayout">
    <el-button class="add_button" type="primary" @click="handleAdd()">增加题目</el-button>
  </el-radio-group>
  <el-table v-loading="loading" :default-sort="{ prop: 'question_id', order: 'ascending' }" :data="result.table" height="100%" stripe style="width: 100%; padding-left: 15px; padding-right: 15px;" :table-layout="fixed">
 <el-table-column fixed id="id" prop="question_id" label="题目序号" width="100"/>
    <el-table-column prop="question_desc" label="题目描述" />
    <el-table-column prop="option_a_desc" label="选项A" />
    <el-table-column prop="option_a_target"  label="倾向" width="100"/>
    <el-table-column prop="option_b_desc" label="选项B" />
    <el-table-column prop="option_b_target"  label="倾向" width="100"/>
        <el-table-column fixed="right" label="操作" width="150">
      <template #default="scope">
        <el-button type="primary" size="small" @click="handleEdit(scope.row)"
          >编辑</el-button>
        <el-button type="danger" size="small" @click="handleDelete(scope.row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>

    <el-dialog 
        v-model="edit_dialog" 
        title="编辑题目" 
        center="true"
        destroy-on-close
        width=35%>
        <el-form 
        :model="login_form">
        <el-form-item 
            label="题目描述" 
            :label-width="formLabelWidth"
            size="large">
            <el-input 
            :prefix-icon="User"
            v-model="edit.question_desc"
            show-word-limit
            autocomplete="off"
            style="width:350px;"/>
        </el-form-item>
        <el-form-item 
            label="选项A描述" 
            :label-width="formLabelWidth"
            size="large">
            <el-input 
            :prefix-icon="User"
            v-model="edit.option_a_desc"
            show-word-limit
            autocomplete="off"
            style="width:350px;"/>
        </el-form-item>
        <el-form-item 
        label="选项A倾向" 
        :label-width="formLabelWidth" 
        style="font-size: 30px">
            <el-select v-model="edit.option_a_target" size="large" suffix-icon="CaretBottom">
                <el-option
                    v-for="item in targets"
                    :key="item"
                    :value="item"
                    >
                </el-option>
            </el-select>
        </el-form-item>
        <el-form-item 
            label="选项B描述" 
            :label-width="formLabelWidth"
            size="large">
            <el-input 
            :prefix-icon="User"
            v-model="edit.option_b_desc"
            show-word-limit
            autocomplete="off"
            style="width:350px;"/>
        </el-form-item>
        <el-form-item 
        label="选项B倾向" 
        :label-width="formLabelWidth" 
        style="font-size: 30px">
            <el-select v-model="edit.option_b_target" size="large" suffix-icon="CaretBottom">
            <el-option
                v-for="item in targets"
                :key="item"
                :value="item"
            >
            </el-option>
            </el-select>
        </el-form-item>

        </el-form>
        <template #footer>
        <span class="dialog-footer">
            <el-button
            class="edit_button"
            id="register"
            type="primary" 
            @click="save_edit()"
            color="#CC0033"
            >保存</el-button>
            <el-button
            class="edit_button"
            id="register"
            type="primary" 
            @click="close_editdialog()"
            color="#CCCCCC"
            >取消</el-button>
        </span>
        </template>
    </el-dialog>

    <el-dialog 
        v-model="add_dialog" 
        title="新增题目" 
        center="true"
        destroy-on-close
        width=35%>
        <el-form 
        :model="login_form">
        <el-form-item 
            label="题目描述" 
            :label-width="formLabelWidth"
            size="large">
            <el-input 
            :prefix-icon="User"
            v-model="edit.question_desc"
            show-word-limit
            autocomplete="off"
            style="width:350px;"/>
        </el-form-item>
        <el-form-item 
            label="选项A描述" 
            :label-width="formLabelWidth"
            size="large">
            <el-input 
            :prefix-icon="User"
            v-model="edit.option_a_desc"
            show-word-limit
            autocomplete="off"
            style="width:350px;"/>
        </el-form-item>
        <el-form-item 
        label="选项A倾向" 
        :label-width="formLabelWidth" 
        style="font-size: 30px">
            <el-select v-model="edit.option_a_target" size="large" suffix-icon="CaretBottom">
                <el-option
                    v-for="item in A"
                    :key="item"
                    :value="item"
                    >
                </el-option>
            </el-select>
        </el-form-item>
        <el-form-item 
            label="选项B描述" 
            :label-width="formLabelWidth"
            size="large">
            <el-input 
            :prefix-icon="User"
            v-model="edit.option_b_desc"
            show-word-limit
            autocomplete="off"
            style="width:350px;"/>
        </el-form-item>
        <el-form-item 
        label="选项B倾向" 
        :label-width="formLabelWidth" 
        style="font-size: 30px">
            <el-select v-model="edit.option_b_target" size="large" suffix-icon="CaretBottom">
            <el-option
                v-for="item in B"
                :key="item"
                :value="item"
            >
            </el-option>
            </el-select>
        </el-form-item>

        </el-form>
        <template #footer>
        <span class="dialog-footer">
            <el-button
            class="edit_button"
            id="register"
            type="primary" 
            @click="save_edit()"
            color="#CC0033"
            >保存</el-button>
            <el-button
            class="edit_button"
            id="register"
            type="primary" 
            @click="close_editdialog()"
            color="#CCCCCC"
            >取消</el-button>
        </span>
        </template>
    </el-dialog>
</template>

<script>
import { getCurrentInstance, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

export default ({
    setup() {
        const router = useRouter()
        const {proxy} = getCurrentInstance()
        const edit_dialog = ref(false)
        const add_dialog = ref(false)
        const formLabelWidth  = "90px"
        const loading = ref(true)

        let result = reactive({
            table: [],
        })

        const func=async()=> {
            const listres = await new proxy.$request(proxy.$urls.m().getquestionidlist).get()
            const idList = listres.question_id_list
            for (let i=0; i<idList.length; i++) {
                const obj = {"id":idList[i]}
                const res = await new proxy.$request(proxy.$urls.m().getquestion, obj).get()
                //草啊，不用赋值啊  直接push就不报错了 我日
                result.table.push(res.question_info)
            }

            setTimeout(3000)
            loading.value = false
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
        
        const edit = reactive({
            question_id: 0,
            question_desc: '',
            option_a_desc: '',
            option_a_target: '',
            option_b_desc: '',
            option_b_target: '',
        })
        const handleEdit = (index) => {
            edit_dialog.value = true
            edit.question_id = index.question_id
            edit.question_desc = index.question_desc
            edit.option_a_desc = index.option_a_desc
            edit.option_a_target = index.option_a_target
            edit.option_b_desc = index.option_b_desc
            edit.option_b_target = index.option_b_targets
        }
        const handleDelete = async(index) => {
            const obj = {"id":index.question_id}
            const res = await new proxy.$request(proxy.$urls.m().deletequestion, obj).post()
            console.log(res)
            ElMessage({
                message: "删除成功",
                type: 'success',
            })
        }

        const handleAdd=()=> {
            add_dialog.value = true
        }

        const close_editdialog=()=> {
            edit_dialog.value = false
        }

        const save_edit=async()=> {
            const listres = await new proxy.$request(proxy.$urls.m().editquestion, edit).post()
            location.reload();
            ElMessage({
                message: listres.msg,
                type: 'success',
                center: true,
            })
            edit_dialog.value = false
        }

        const targets = ['E','I','S','N','T','F','J','P']
        return {
            targets,
            edit,
            loading,
            add_dialog,
            handleEdit,
            handleAdd,
            save_edit,
            close_editdialog,
            formLabelWidth,
            edit_dialog,
            handleDelete,
            result,
            goHome,
        }
    },
})
</script>

<style scoped lang="less">
.edit_button {
        color: #fff;
        font-size: 20px;
        height: 50px;
        width: 25%;
        border-radius: 50px
}

.el-radio-group {
    padding-left: 20px;
}


</style>

