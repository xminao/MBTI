<template>
    <el-container>
        <el-button
        id="test"
        round
        color="#88619A">
            开始测试
        </el-button>
    </el-container>
        <el-image style="width: 100%; height: 408px;" :src="require('../assets/homepage.png')" :fit="fill"/>
</template>

<style scoped lang="less">
    #left {
        background-color: aliceblue;
    }
    #right {
        background-color: aquamarine;
    }

    #test {
        color: #fff;
        font-size: 20px;
        height: 60%;
        width: 13%;
        border-radius: 50px
    }

    .el-container {
        justify-content: center;
        height: 100px;
     }

     .el-button {
         position: relative;
         top: 20px;
     }
</style>