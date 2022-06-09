<template>
<el-container>
  <div id="charts">
  </div>
  <el-button @click="getinfo()">点击查看数据</el-button>
  </el-container>
</template>

<script>
import * as echarts from 'echarts'
import { getCurrentInstance, reactive } from 'vue'
export default {
    
    setup() {
        const {proxy} = getCurrentInstance()
        const getinfo=async()=> {
            const listres = await new proxy.$request(proxy.$urls.m().getdatalist).get()
            const datas = listres.data_list

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
                title: { text: "各测试结果占比" },
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
        }
    }
}
</script>

<style>
 #charts {
     width: 800px;
     height: 800px;
 }
</style>