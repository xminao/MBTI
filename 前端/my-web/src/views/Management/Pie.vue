<template>
<el-container class="select">
  <el-radio-group>
     <el-tree-select 
     v-model="value" 
     accordion 
     :data="list.data" 
     placeholder="请选择"
     :render-after-expand="false"
     @node-click="getNodeValue"
    check-strictly/>
  </el-radio-group>
</el-container>
<el-container v-model="loading" class="chart">
<div id="charts"/>
</el-container>
<!-- <el-container>
  <div id="charts">
  </div>
  <el-button @click="getinfo()">点击查看数据</el-button>
  </el-container> -->
</template>

<script>
import { getCurrentInstance, reactive, ref } from 'vue'
import * as echarts from 'echarts'
import { useRouter } from 'vue-router'

export default ({
    setup() {
        const router = useRouter()
        const {proxy} = getCurrentInstance()
        const index = ref(0)
        const value = ref('')
        const loading = ref(false)

            var option;

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
                    "value": datas[i].college_id + '',
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
                                    "value": datas[i].college_id + '-' + yearres.year_list[k].year_id,
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
                                            "value": datas[i].college_id + '-' + yearres.year_list[k].year_id + '-' + majorres.major_list[l].major_id,
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
                                                    "value": datas[i].college_id + '-' + yearres.year_list[k].year_id + '-' + majorres.major_list[l].major_id + '-' + classres.class_list[n].class_id,
                                                    "label": classres.class_list[n].class_name,
                                                    //"children": []
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
        }

        test()

        const getH=(param)=> {
            return true
        }

        let table = reactive({
            data: []
        })

        let data = []

        const getStuList=async(node, id)=> {
            const listres = await new proxy.$request(proxy.$urls.m().getstudentlist).get()
            if (listres.student_list != null) {
                for (let i=0; i<listres.student_list.length; i++) {
                    const temp = listres.student_list
                    data.push(temp[i])
                    data[i]["college_name"] = JSON.parse(temp[i].college).college_name
                    data[i]["year_name"] = JSON.parse(listres.student_list[i].year).year_name
                    data[i]["major_name"] = JSON.parse(listres.student_list[i].major).major_name
                    data[i]["class_name"] = JSON.parse(listres.student_list[i].class).class_name
                   //console.log(data)
                }
                if (node == 1) { //学院
                    for (let k=0; k<data.length; k++) {
                        if (JSON.parse(data[k].college).college_id != id) {
                            data.splice(k, 1);
                            k--;
                        }
                    }
                } else if (node == 3) { //年级
                    for (let k=0; k<data.length; k++) {
                        if (JSON.parse(data[k].year).year_id != id) {
                            data.splice(k, 1);
                            k--;
                        }
                    }
                } else if (node == 5) { //专业
                    for (let k=0; k<data.length; k++) {
                        if (JSON.parse(data[k].major).major_id != id) {
                            data.splice(k, 1);
                            k--;
                        }
                    }
                } else if (node == 7) { //班级
                    for (let k=0; k<data.length; k++) {
                        if (JSON.parse(data[k].class).class_id != id) {
                            data.splice(k, 1);
                            k--;
                        }
                    }
                }
            }
           
            const datalistres = await new proxy.$request(proxy.$urls.m().getdatalist).get()
            console.log(datalistres)
            let datas = datalistres.data_list //测试结果
            const userlistres = await new proxy.$request(proxy.$urls.m().getuserlist).get()
            const userlist = userlistres.userlist //用户表
            console.log(userlist)
            console.log(data)

            for (let i=0; i<datas.length; i++) {
                let flag = false
                for (let j=0; j<data.length; j++) {
                    for (let k=0; k<userlist.length; k++) {
                        // console.log(data[j].student_id)
                        // console.log(datas[i].username)
                        // console.log(userlist[k].username)
                        if ((datas[i].username == userlist[k].username) && (userlist[k].binding_student_id == data[j].student_id)) {
                            flag = true
                            break
                        }
                        //if(datas[i].use)
                        // if (datas[i].username == userlist[k].username && userlist[k].binding_student_id == data[j].student_id)
                        //     flag = true
                        //     break;
                    }
                }
                if (flag == false) {
                    datas.splice(i, 1);
                    i--;
                }
            }

        let myChart = echarts.init(document.getElementById("charts"));

        option = {
            legend: {},
            tooltip: {
            trigger: 'axis',
            showContent: false
            },
            dataset: {
                source: [
                    // ['product', '2012', '2013', '2014', '2015', '2016', '2017'],
                    // ['Milk Tea', 56.5, 82.1, 88.7, 70.1, 53.4, 85.1],
                    // ['Matcha Latte', 51.1, 51.4, 55.1, 53.3, 73.8, 68.7],
                    // ['Cheese Cocoa', 40.1, 62.2, 69.5, 36.4, 45.2, 32.5],
                    // ['Walnut Brownie', 25.2, 37.1, 41.2, 18, 33.9, 49.1],
                    // ['Walnut', 25.2, 37.1, 41.2, 18, 33.9, 49.1],
                    // [' Brownie', 25.2, 37.1, 41.2, 18, 33.9, 49.1]
                    //['product', '2022-06-08', '2022-06-09', '2022-06-10', '2022-06-11'],
                    ['date'],
                    ['INTJ'],
                    ['INTP'],
                    ['ENTJ'],
                    ['ENTP'],
                    ['INFJ'],
                    ['INFP'],
                    ['ENFJ'],
                    ['ENFP'],
                    ['ISTJ'],
                    ['ISFJ'],
                    ['ESTJ'],
                    ['ESFJ'],
                    ['ISTP'],
                    ['ISFP'],
                    ['ESTP'],
                    ['ESFP'],
                ]
            },

            xAxis: { type: 'category' },
            yAxis: { gridIndex: 0 },
            grid: { top: '55%' },
            series: [
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'line',
                smooth: true,
                seriesLayoutBy: 'row',
                emphasis: { focus: 'series' }
            },
            {
                type: 'pie',
                id: 'pie',
                radius: '30%',
                center: ['50%', '25%'],
                emphasis: {
                focus: 'self'
                },
                label: {
                formatter: '{b}: {@2022-06-10} ({d}%)'
                },
                encode: {
                itemName: 'date',
                value: '2022-06-10',
                tooltip: '2022-06-10'
                }
            }
            ]
        };

        for (let i=0; i<datas.length; i++) {
            let date = (datas[i].time).split(' ')[0]
            let flag = false
            for (let k=1 ;k<option.dataset.source[0].length; k++) {
                if (option.dataset.source[0][k] == date) {
                    flag = true
                }
            }
            if (flag == false) {
                option.dataset.source[0].push(date)
                for (let k=1; k<option.dataset.source.length; k++) {
                    option.dataset.source[k].push(0)
                }
            }
        }

        //二维遍历，横竖
        for (let i=0; i<datas.length; i++) {
            let date = (datas[i].time).split(' ')[0]
            //寻找是哪个人格
            for (let j=1; j<option.dataset.source.length; j++) {
                if (datas[i].type == option.dataset.source[j][0]) {
                    //寻找日期
                    for (let k=1; k<option.dataset.source[0].length; k++) {
                        if (date == option.dataset.source[0][k]) {
                            option.dataset.source[j][k]++
                        }
                    }
                }
            }
        }
        

        /******************* */

        myChart.on('updateAxisPointer', function (event) {
            const xAxisInfo = event.axesInfo[0];
            if (xAxisInfo) {
            const dimension = xAxisInfo.value + 1;
            myChart.setOption({
                series: {
                id: 'pie',
                label: {
                    formatter: '{b}: {@[' + dimension + ']} ({d}%)'
                },
                encode: {
                    value: dimension,
                    tooltip: dimension
                }
                }
            });
            }
        });
        myChart.setOption(option);
        loading.value(true)
        }


        const getNodeValue=(val)=> {
            data = []
            getStuList(val.value.length, val.id)
        }

        const getinfo=async()=> {

            const listres = await new proxy.$request(proxy.$urls.m().getdatalist).get()
            const datas = listres.data_list

            //console.log(data)
            for (let i=0; i<data.length; i++) {
                console.log(data[i])
            }

            let rest = reactive({
                INTJ: 0,
                INTP: 0,
                ENTJ: 0,
                ENTP: 0,
                ENFJ: 0,
                INFJ: 0,
                INFP: 0,
                ENFP: 0,
                ISTJ: 0,
                ISFJ: 0,
                ESTJ: 0,
                ESFJ: 0,
                ISTP: 0,
                ISFP: 0,
                ESTP: 0,
                ESFP: 0,
            })

            for (let i=0; i<datas.length; i++) {

                switch (datas[i].type) {
                    case 'INTJ':
                        rest.INTJ++
                        break;
                    case 'INTP':
                        rest.INTP++
                        break;
                    case 'ENTJ':
                        rest.ENTJ++
                        break;
                    case 'ENTP':
                        rest.ENTP++
                        break;
                    case 'ENFJ':
                        rest.ENFJ++
                        break;
                    case 'INFJ':
                        rest.INFJ++
                        break;
                    case 'INFP':
                        rest.INFP++
                        break;
                    case 'ENFP':
                        rest.ENFP++
                        break;
                    case 'ISTJ':
                        rest.ISTJ++
                        break;
                    case 'ISFJ':
                        rest.ISFJ++
                        break;
                    case 'ESTJ':
                        rest.ESTJ++
                        break;
                    case 'ESFJ':
                        rest.ESFJ++
                        break;
                    case 'ISTP':
                        rest.ISTP++
                        break;
                    case 'ISFP':
                        rest.ISFP++
                        break;
                    case 'ESTP':
                        rest.ESTP++
                        break;
                    case 'ESFP':
                        rest.ESFP++
                        break;
                    default:
                        break;
                }

            }

            let myChart = echarts.init(document.getElementById("charts"));
            // 绘制图表
            myChart.setOption({
                // title: { text: "测试结果图" },
                series: [
                    {
                        type: 'pie',
                        data: [
                            {
                                value: rest.ENFJ,
                                name: 'ENFJ'
                            },
                            {
                                value: rest.ENFP,
                                name: 'ENFP'
                            },
                            {
                                value: rest.ENTJ,
                                name: 'ENTJ'
                            },
                            {
                                value: rest.ENTP,
                                name: 'ENTP'
                            },
                            {
                                value: rest.ESFJ,
                                name: 'ESFJ'
                            },
                            {
                                value: rest.ESFP,
                                name: 'ESFP'
                            },
                            {
                               value: rest.ESTJ,
                                name: 'ESTJ'
                            },
                            {
                                value: rest.ESTP,
                                name: 'ESTP'
                            },
                            {
                               value: rest.INFJ,
                                name: 'INFJ'
                            },
                            {
                                value: rest.INFP,
                                name: 'INFP'
                            },
                            {
                                value: rest.INTJ,
                                name: 'INTJ'
                            },
                            {
                                value: rest.INTP,
                                name: 'INTP'
                            },
                            {
                                value: rest.ISFJ,
                                name: 'ISFJ'
                            },
                            {
                                value: rest.ISFP,
                                name: 'ISFP'
                            },
                            {
                                value: rest.ISTJ,
                                name: 'ISTJ'
                            },
                            {
                                value: rest.ISTP,
                                name: 'ISTP'
                            }
                        ],
                        roseType: 'area'
                    }
                ]
            });
            window.onresize = function () {//自适应大小
                myChart.resize();
            };
        }

        return {

            getinfo,
            getNodeValue,
            getH,
            list,
            result,
            goHome,
            index,
            value,
            table,
            defaultProps,
            handleNodeClick,
        }
    },
})
</script>

<style scoped lang="less">

// .el-container {
//     width: 100%;
//     height: 100%;
//     background-color: #fff;
// }

.el-table {
    padding-top: 5px;
}

.el-radio-group {
    padding-left: 15px;
    width: 20%;
}

 #charts {
     width: 800px;
     height: 800px;
 }

.select {
    width: 100%;
    height: 10%;
    background-color: #fff;
}
.chart {
    width: 100%;
    padding-left: 5%;
    height: 80%;
    background-color: #fff;
}


</style>

