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
          <el-menu-item index="/type" route="/type"><Menu style="width: 1em; height: 1em; margin-right: 8px" />人格类型</el-menu-item>
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
              title="登录/注册" 
              center="true"
              width=40%>
              <el-form :model="form">
                <el-form-item 
                  label="用户名" 
                  :label-width="formLabelWidth"
                  size="large">
                </el-form-item>
                <el-input 
                    v-model="form.username" 
                    autocomplete="off"
                    style="width:300px;"/>
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
                    @click="logindialog = false"
                    color="#88619A"
                    >登录</el-button>
                </span>
              </template>
          </el-dialog>
        </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
    </el-scrollbar>
  </div>
</template>


<script lang="ts">
import { useRouter } from 'vue-router';
import { defineComponent, reactive, ref } from 'vue';

export default defineComponent({
  setup () {
    const login = '登录'
    const router = useRouter()
    const onSubmit = () => {
      router.push('/login')
    }

    const logindialog = ref(false)
    const formLabelWidth = '90px'

    const form = reactive({
      username: '',
      passwd: '',
      region: '',
      date1: '',
      date2: '',
      delivery: false,
      type: [],
      resource: '',
      desc: '',
    })

    return {login, onSubmit, form, logindialog, formLabelWidth}
  },
  
  name: 'HomeView',
});
</script>

<style scoped lang="less">
  
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
  .common-layout {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
  }
  .el-container{
    height: 100%;
  }
  .el-header{
     height: 90px;
     background-color: #fff;
     display: flex;
     justify-content: space-between;
     padding-left: 0.5cm;
     align-items: center;
     color: #fff;
     font-size: 20px;
     > #sign {
        color: #fff;
        font-size: 20px;
        height: 60%;
        width: 13%;
        border-radius: 50px
     }
     > .el-menu-demo {
       width: 450px;
     }
  }
  .el-main{
    height: 100%;
    background-color: #008B8B;
    color: rgb(221, 6, 6);text-align: center;line-height: 0px;
  }
</style>
