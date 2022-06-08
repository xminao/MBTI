<template>
    <el-container 
        class="result"
        >
        <el-aside>
            <el-image style="width: 350px; height: 350px;" :src="attri.imgsrc" :fit="fill"/>
        </el-aside>
        <el-main>
            <div class="content">
            <p id="name">{{attri.name}}</p>
            <p id="label">{{attri.label}}</p>
            </div>
        </el-main>
    </el-container>
</template>

<script>
import { useRouter, useRoute } from 'vue-router';
import { reactive, getCurrentInstance } from 'vue'

export default({
    setup() {
        const router = useRouter()
        const route = useRoute()
        const {proxy} = getCurrentInstance()

        const goTest=()=> {
            router.push('/test')
        }

        let attri = reactive({
            name: '',
            label: '',
            imgsrc: '',
            bcolor: '',
            name_color: '',
            label_color: ''
        })

        const types = {
            INTJ: '建筑师',
            INTP: '逻辑学家',
            ENTJ: '指挥官',
            ENTP: '辩论家',
            ENFJ: '主人公',
            INFJ: '倡导者',
            INFP: '调停者',
            ENFP: '竞选者',
            ISTJ: '物流师',
            ISFJ: '守卫者',
            ESTJ: '总经理',
            ESFJ: '执政官',
            ISTP: '鉴赏家',
            ISFP: '探险家',
            ESTP: '企业家',
            ESFP: '表演者',
        }

        const colors = {
            analysts: {
                bg: '#E7DFEA',
                name: '#584164',
                label: '#90759F',
            },
            diplomats: {
                bg: '#D6ECE3',
                name: '#35785D',
                label: '#99C26C',
            },
            sentinels: {
                bg: '#D9EAF0',
                name: '#369395',
                label: '#71CACC',
            },
            explorers: {
                bg: '#F9EED7',
                name: '#BE8F00',
                label: '#F4D65D',
            },
        }

        const test = {
            analysts: {
                analysts: ['INTJ', 'INTP', 'ENTJ', 'ENTP'],
                colors: {
                    bg: '#E7DFEA',
                    name: '#584164',
                    label: '#90759F',
                },
            },
        }

        const groups = {
            analysts: ['INTJ', 'INTP', 'ENTJ', 'ENTP'],
            diplomats: ['INFJ', 'INFP', 'ENFJ', 'ENFP'],
            sentinels: ['ISTJ', 'ISFJ', 'ESTJ', 'ESFJ'],
            explorers: ['ISTP', 'ISFP', 'ESTP', 'ESFP'],
        }

        const getGroup=(type)=>{
            let i
            let j
            for (i in groups) {
                for (j=0; j<4; j++) {
                    if (groups[i][j] == type) {
                        return i
                    }
                }
            }
            return
        }


        const getTypeRecord=async()=> {
            const res = await new proxy.$request(proxy.$urls.m().getlatestdata).get()
            let result = res.type
            let group = getGroup(result)
            attri.label = result
            attri.name = types[result]
            attri.imgsrc = require('../assets/' + result + '.png')
            // 结果的界面的颜色
            attri.bcolor = colors[group].bg
            attri.name_color = colors[group].name
            attri.label_color = colors[group].label
        }

        getTypeRecord()

        // let group = getGroup(result)
        // attri.label = result
        // attri.name = types[result]
        // attri.imgsrc = require('../assets/' + result + '.png')
        // // 结果的界面的颜色
        // let bcolor = colors[group].bg
        // let name_color = colors[group].name
        // let label_color = colors[group].label

        return {
            router,
            attri,
            goTest,
        }
    },
})
</script>
>

<style scoped lang="less">
    .result {
        justify-content: center;
        background-color: v-bind("attri.bcolor");
        height: 100%;
        width: 100%;
        padding-left: 15%;
        padding-right: 15%;
        //padding: 5%;
    }

    .el-aside {
        width: 50%;
    }

    .contentmenu {
        width: 40%;
        height: 100%;
    }

    .el-image {
        position: relative;
        top: 20%;
        left: 20%;
    }

    .el-main {
        width: 50%;
        height: 100%;
    }

    .content {
        position: relative;
        text-align: center;
        top: 20%;
    }

    #name {
        font-size: 60px;
        color: v-bind("attri.name_color");
        font-family: 幼圆;
        font-weight: bold;  
    }

    #label {
        font-size: 60px;
        color: v-bind("attri.label_color");
        font-family: 华文琥珀;
        font-weight: bold; 
        line-height: 30px
    }

    // .box-card {
    //     background-color: #fff;
    //     position: relative;
    //     top: 10%;
    //     height: 400px;
    //     width: 500px;
    //     border-radius: 40px
    // }
</style>