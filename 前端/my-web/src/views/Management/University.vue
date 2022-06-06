<template>
    <el-container>
        <el-aside>
              <!-- <el-menu
                default-active="aaa"
                class="el-menu-vertical-demo"
                :collapse="isCollapse"
                unique-opened
                @open='index="aaa"'
                @close="handleClose">

                    <template v-for="item in list.data" :key="item">
                        <el-sub-menu :index="item.id">
                            <template #title>{{item.label}}</template>
                                <template v-for="item in list.data" :key="item">
                                    <el-sub-menu :index="item.id">
                                        <template #title>{{item.label}}</template>
                                    </el-sub-menu>
                                </template>
                        </el-sub-menu>
                    </template>



                    <mean-treee :data="item"></mean-treee>

                </el-menu> -->
                <el-tree 
                    :data="list.data" 
                    :props="defaultProps"
                    accordion
                    :default-expand-all="false"
                    @node-click="handleNodeClick" />
        </el-aside>

        <el-main>
            <el-table :data="result.table" height="100%" stripe style="width: 100%; padding-left: 15px; padding-right: 15px;" :table-layout="fixed">
            <el-table-column fixed id="id" prop="question_id" sortable label="题目序号" width="150"/>
                <el-table-column prop="question_desc" label="题目描述" />
                <el-table-column prop="option_a_desc" label="选项A" />
                <el-table-column prop="option_a_target" sortable  label="倾向" width="100"/>
                <el-table-column prop="option_b_desc" label="选项B" />
                <el-table-column prop="option_b_target" sortable  label="倾向" width="100"/>
            </el-table>
        </el-main>
    </el-container>
</template>

<script>
import { getCurrentInstance, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'

export default ({
    setup() {
        const router = useRouter()
        const {proxy} = getCurrentInstance()
        const index = ref(0)

        const defaultProps = {  
            children: "children",  //"children"内的每个对象解析为一个子项;
            label: "label",  //所有"label"所在的对象解析为一个父项
        };

        const handleNodeClick=(data)=>{
            router.push(data.url);  //本例中我们拿取url属性做跳转;
        };

        let result = reactive({
            college_table: [],
            year_table: [],
            major_table: [],
            class_table: [],
        })

        const func=async()=> {
            let listres = await new proxy.$request(proxy.$urls.m().getcollegelist).get()
            let datas = listres.college_list
            for (let i=0; i<datas.length; i++) {
                result.college_table.push(datas[i])

                //获取该学院的年级列表
                const yearobj = {"college_id": datas[i].college_id}
                const yearres = await new proxy.$request(proxy.$urls.m().getyearlist, yearobj).get()
                //该学院有年级
                if (yearres.year_list != null) {
                    result.year_table.push({[datas[i].college_id]: yearres.year_list})
                }
            }
        }

        func()
        

        const goHome=()=> {
            router.push('/home')
        }

        let list = reactive({
            data: [],
        })

        const test=async()=> {
            //获取学院列表，放进第一层
            const listres = await new proxy.$request(proxy.$urls.m().getcollegelist).get()
            const datas = listres.college_list
            for (let i=0; i<datas.length; i++) {
                list.data.push({
                    "id": datas[i].college_id,
                    "label": datas[i].college_name,
                    "children": []
                })
                // 获取该专业年级列表，放进第二层
                const yearobj = {"college_id": datas[i].college_id}
                const yearres = await new proxy.$request(proxy.$urls.m().getyearlist, yearobj).get()
                // 如果该专业存在年级
                if (yearres.year_list != null) {
                    //遍历第一层，找到该插入的地方
                    for (let j=0; j<list.data.length; j++) {
                        //如果找到学院，则插入年级信息
                        if (datas[i].college_id == list.data[j].id) {
                            console.log(yearres.year_list.length)
                            for (let k=0; k<yearres.year_list.length; k++) {
                                list.data[j].children.push({
                                    "id": yearres.year_list[k].year_id,
                                    "label": yearres.year_list[k].year_name,
                                    "children": []
                                })
                            }
                            // data[j].children.push({
                            //     "id": yearres.year_list.id
                            // })
                        }
                        //console.log(data[j].id)
                    //    if (data[i].id == yearres.year_list.year_id) {
                    //        console.log("哈哈哈哈")
                    //    }
                    }
                }
            }

            console.log(list.data)
        }

        test()


        return {
            list,
            result,
            goHome,
            index,
            defaultProps,
            handleNodeClick,
        }
    },
})
</script>

<style scoped lang="less">

.el-container {
    width: 100%;
    height: 100%;
    padding-left: 15px;
    padding-right: 15px;
    background-color: #fff;
}

.el-aside {
    width: 20%;
    background-color: #FFCCCC;
}


</style>

