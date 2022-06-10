<template>
<el-container>
    <div id="charts">
    </div>
    <el-button @click="getinfo()">点击查看数据</el-button>
</el-container>
</template>

<script>
import { getCurrentInstance, reactive, ref } from 'vue'
import * as echarts from 'echarts'
import { useRouter } from 'vue-router'

export default {
    setup() {
    var option;
    const {proxy} = getCurrentInstance()
    const getinfo=async()=> {
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
                formatter: '{b}: {@2012} ({d}%)'
                },
                encode: {
                itemName: 'product',
                value: '2012',
                tooltip: '2012'
                }
            }
            ]
        };

            /*************** */
        const listres = await new proxy.$request(proxy.$urls.m().getdatalist).get()
        const datas = listres.data_list

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
}

option && myChart.setOption(option);

    return {
        getinfo,
    }
    }
}
</script>

<style>
.el-container {
    width: 100%;
    height: 100%
}
 #charts {
     width: 100%;
     height: 800px;
 }
</style>