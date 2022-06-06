<template>
    <el-container class="result">
        <el-aside>
            <el-image style="width: 350px; height: 350px;" :src="require('../assets/enfj-protagonist.png')" :fit="fill"/>
        </el-aside>
        <el-main>
            <div class="content">
            <p id="name">{{type.name}}</p>
            <p id="label">{{type.label}}</p>
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

        let type = reactive({
            name: '',
            label: '',
        })

        const types = {
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

        let result = route.params.type
        type.label = result
        type.name = types[result]

        return {
            router,
            type,
            goTest,
        }
    },
})
</script>
>

<style scoped lang="less">
    .result {
        justify-content: center;
        background-color: #D6ECE3;
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
        color: #99C26C;
        font-family: 微软雅黑;
        font-weight: bold;  
    }

    #label {
        font-size: 60px;
        color: #B8C645;
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