<template>
  <div class="common-layout">
    <el-scrollbar>
    <el-container>
      <el-header>
          <div class="logo">
            <img src="../assets/logo.png" alt="logo">
          </div>
          <el-menu 
          router=true
          class="el-menu-demo" 
          mode="horizontal"
          text-color="#000000"
          active-text-color="#4298B4"
          >
          <el-menu-item index="/test"><List style="width: 1em; height: 1em; margin-right: 8px" />人格测试</el-menu-item>
          <el-menu-item> </el-menu-item>
          <el-menu-item index="/type"><Menu style="width: 1em; height: 1em; margin-right: 8px" />人格类型</el-menu-item>
          <el-menu-item> </el-menu-item>
          <el-menu-item index="/login"><Promotion style="width: 1em; height: 1em; margin-right: 8px" />联系我们</el-menu-item>
          </el-menu>
          <el-button 
            id="sign" 
            type="info"
            color="#88619A"
            @click="logindialog = true" >登录 / 注册</el-button>
            <el-dialog 
              v-model="logindialog" 
              title="登录" 
              center="true"
              @close="signin"
              width=40%>
              <el-form :model="form">
                <el-form-item 
                  label="用户名" 
                  :label-width="formLabelWidth"
                  size="large">
                  <el-input 
                    v-model="form.username" 
                    autocomplete="off"
                    style="width:300px;"/>
                </el-form-item>
                <el-form-item label="密码" :label-width="formLabelWidth" style="font-size: 30px">
                  <el-input 
                    v-model="form.passwd" 
                    autocomplete="off"
                    type="password"
                    size="large"
                    style="width:300px;"
                    show-password />
                </el-form-item>
              </el-form>
              <template #footer>
                <span class="dialog-footer">
                  <el-button
                    id="login"
                    type="primary" 
                    @click="signin()"
                    color="#88619A"
                    >登录</el-button>
                </span>
              </template>
          </el-dialog>
        </el-header>

      <el-main>
          <router-view/>
      </el-main>
      <el-footer>Footer</el-footer>
    </el-container>
    </el-scrollbar>
  </div>
</template>



<script>
import { useRouter } from 'vue-router';

export default {
    setup () {
        const router = useRouter()
        return { router }
    },
    //name: 'Main'
}
</script>


<style lang="less" scoped>
    .common-layout {
        position: absolute;
        left: 0;
        right: 0;
        top: 0;
        bottom: 0;
        height: 100%;
    }

    .el-container {
        height: 100%;
    }

    .el-header {
        height: 80px;
        background-color: #fff;
        display: flex;
        justify-content: space-between;
        padding-left: 0.5cm;
        // 居中
        align-items: center;
    }
    #login {
        color: #fff;
        font-size: 20px;
        height: 50px;
        width: 30%;
        border-radius: 50px
    }
    .logo img{
        width: 214.5px;
        height: 44px;
    }
    .el-menu-demo {
       width: 450px;
    }


    .el-main {
        background-color: #4298B4;
        width: 100%;
        height: 100%
    }
    .screen {
        height: 100%;
        width: 100%;
    }

    .el-footer {
        background-color: #477777;
        height: 40px;
    }
</style>