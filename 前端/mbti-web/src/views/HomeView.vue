<template>
  <div class="common-layout">
    <el-container>
      <el-header>
          <div class="logo">
            <img src="../assets/logo.svg" alt="出错了">
          </div>
          <ele-button type="info" @click="login">登录</ele-button>
        </el-header>
      <el-main>
        <div>
          <span>测试</span>
        </div>
        </el-main>
    </el-container>
  </div>
</template>


<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  setup () {
    const login = '登录'

    return {login}
  },
  name: 'HomeView',
});
</script>

<style scoped>
  .common-layout {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
  }

  .logo {
      width: 50%;
  }
  .el-container{
    height: 100%;
  }
  .el-header{
     background-color: #adaaa8;
     display: flex;
     justify-content: space-between;
  }
  .el-main{
    height: 100%;
    background-color: #78754e;
    color: rgb(221, 6, 6);text-align: center;line-height: 0px;
  }
</style>
