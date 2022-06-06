<template>
    <el-container class="text">
        <span><br/><br/><font face="微软雅黑" size="10" color="#fff">快来测试一下自己的人格类型吧</font></span>
    </el-container>
    <el-container class="littletext">
        <span><font face="微软雅黑" size="4" color="#fff">得到一个具体而准确的对自己的描述，了解你行事方式背后的原因。</font></span>
    </el-container>
    <el-container class="start">
            <el-button
            id="test"
            round
            @click="goTest()"
            color="#88619A">开始测试</el-button>
    </el-container>
    <el-container class="tips">
        <span><font color="#FFF"><br/>这将花费约 10 分钟</font></span>
    </el-container>
        <el-image style="width: 100%; height: 390px;" :src="require('../assets/home.png')" :fit="fill"/>
</template>

<script>
import { useRouter } from 'vue-router';

export default({
    setup() {
        const router = useRouter()

        const goTest=()=> {
            router.push('/test')
        }

        return {
            router,
            goTest,
        }
    },
})
</script>
>

<style scoped lang="less">

    #test {
        color: #fff;
        font-size: 20px;
        height: 70%;
        width: 20%;
        border-radius: 50px
    }

    .text {
        justify-content: center;
        background-color: #4298B4;
        height: 150px;
        padding: 0%;
    }
    .littletext {
        justify-content: center;
        background-color: #4298B4;
        height: 50px;
        padding: 0%;
    }
    .start {
        justify-content: center;
        background-color: #4298B4;
        height: 100px;
        padding: 0%;
    }
    .tips {
        justify-content: center;
        background-color: #4298B4;
        height: 100px;
        padding: 0%;
    }

     .el-button {
         position: relative;
         top: 30px;
     }
</style>