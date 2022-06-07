<template>
  <el-radio-group v-model="tableLayout">
     <el-tree-select v-model="value" :data="list.data" :render-after-expand="false" />
  </el-radio-group>
  <el-table :data="result.table" height="100%" stripe style="width: 100%; padding-left: 15px; padding-right: 15px;" :table-layout="fixed">
    <el-table-column fixed type="index" :index="indexMethod" />
    <el-table-column id="id" prop="student_id" sortable label="学号"/>
    <el-table-column prop="student_name" label="姓名" />
    <el-table-column prop="college_name" label="学院" />
    <el-table-column prop="year_name"  label="年级"/>
    <el-table-column prop="major_name" label="专业" />
    <el-table-column prop="class_name" label="班级" />
  </el-table>
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
            //router.push(data.url);

        };

        //const map = new Map(Object.entries(mapObj));

        let result = reactive({
            college_table: [],
            year_table: [],
            major_table: [],
            class_table: [],
        })

        const handleCollege=(parm)=> {
            let data = []

        }

        const func=async()=> {
            const collegeres = await new proxy.$request(proxy.$urls.m().getcollegelist).get()
            let datas = collegeres.college_list
            for (let i=0; i<datas.length; i++) {
                datas = collegeres.college_list
                result.college_table.push(datas[i])

                //获取该学院的年级列表
                const yearobj = {"college_id": datas[i].college_id}
                const yearres = await new proxy.$request(proxy.$urls.m().getyearlist, yearobj).get()
                //该学院有年级
                if (yearres.year_list != null) {
                    datas = yearres.year_list
                    for (let k=0; k<yearres.year_list.length; k++) {
                        var str = toString(datas[k].college_id)
                        if (result.year_table.length == 0) {
                            result.year_table.push({str : datas[k]})
                        } else {
                            for (let h=0; h<result.year_table.length; h++) {
                                //if (result.year_table[h])
                            }
                        }
                    }
                }
            }
            console.log(result.year_table.length)
        }

        //func()
        

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
                // 获取该学院年级列表，放进第二层
                const yearobj = {"college_id": datas[i].college_id}
                const yearres = await new proxy.$request(proxy.$urls.m().getyearlist, yearobj).get()
                // 如果该学院存在年级
                if (yearres.year_list != null) {
                    for (let j=0; j<list.data.length; j++) {
                        //如果找到学院，则插入年级信息
                        if (datas[i].college_id == list.data[j].id) {
                            //console.log(yearres.year_list.length)
                            for (let k=0; k<yearres.year_list.length; k++) {
                                list.data[j].children.push({
                                    "id": yearres.year_list[k].year_id,
                                    "label": yearres.year_list[k].year_name,
                                    "children": []
                                })
                                // 获取该年级的专业列表，放进第三层
                                const majorobj = {"year_id": yearres.year_list[k].year_id}
                                const majorres = await new proxy.$request(proxy.$urls.m().getmajorlist, majorobj).get()
                                //如果该年级存在专业
                                if (majorres.major_list != null) {
                                    for (let l=0; l<majorres.major_list.length; l++) {
                                        list.data[j].children[k].children.push({
                                            "id": majorres.major_list[l].major_id,
                                            "label": majorres.major_list[l].major_name,
                                            "children": []
                                        })
                                        // 获取该专业的班级列表，放进第四层
                                        const classobj = {"major_id": majorres.major_list[l].major_id}
                                        const classres = await new proxy.$request(proxy.$urls.m().getclasslist, classobj).get()
                                        if (classres.class_list != null) {
                                            for (let n=0; n<classres.class_list.length; n++) {
                                                list.data[j].children[k].children[l].children.push({
                                                    "id": classres.class_list[n].class_id,
                                                    "label": classres.class_list[n].class_name,
                                                    "children": []
                                                })
                                            }
                                        }
                                    }
                               }
                            }
                        }
                    }
                }
            }

            console.log(list.data)
        }

        test()

        const getH=(param)=> {
            return true
        }

        return {
            getH,
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
    background-color: #000;
}


</style>

