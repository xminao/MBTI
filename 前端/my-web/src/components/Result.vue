<template>
    <el-container 
        class="result"
        >
        <el-aside>
            <el-image style="width: 350px; height: 350px;" :src="imgsrc" :fit="fill"/>
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
import { reactive } from 'vue'

export default({
    setup() {
        const router = useRouter()
        const route = useRoute()

        const goTest=()=> {
            router.push('/test')
        }

        let attri = reactive({
            name: '',
            label: '',
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
                images: {
                }
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

        let result = route.params.type
        let group = getGroup(result)
        attri.label = result
        attri.name = types[result]
        // 结果的界面的颜色
        let bcolor = colors[group].bg
        let name_color = colors[group].name
        let label_color = colors[group].label

        let imgsrc = require('../assets/enfj.png')

        return {
            router,
            bcolor,
            name_color,
            label_color,
            attri,
            goTest,
            imgsrc,
        }
    },
})
</script>
>

<style scoped lang="less">
    .result {
        justify-content: center;
        background-color: v-bind(bcolor);
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
        color: v-bind(name_color);
        font-family: 微软雅黑;
        font-weight: bold;  
    }

    #label {
        font-size: 60px;
        color: v-bind(label_color);
        font-family: 华文琥珀;
        font-weight: bold;  
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