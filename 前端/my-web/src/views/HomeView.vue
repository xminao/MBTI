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
        <router-view></router-view>
      </el-main>
    </el-container>
    </el-scrollbar>
  </div>
</template>


<script>
import { useRouter } from 'vue-router';
import { reactive, ref, getCurrentInstance } from 'vue';
import { ElMessage } from 'element-plus';
import axios from "axios"
import qs from "qs"

export default {
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

    const {proxy} = getCurrentInstance()

    const closeDialog = () =>{
      logindialog = false
    }

    //async await成对
    const signin = async()=> {
      if (form.username.length == 0) {
        ElMessage({
          message: "用户名不能为空",
          type: 'warning',
        })
        return
      }
      if (form.passwd.length == 0) {
        ElMessage({
          message: "密码不能为空",
          type: 'warning',
        })
        return
      }

      const obj = {"username":form.username, "password":form.passwd}
      const res = await new proxy.$request(proxy.$urls.m().login, obj).post()
      console.log(res)
      localStorage.setItem('token', res.data.jwt_token.access_token)

      const obj_info = {"username":form.username} 
      const info_res = await new proxy.$request(proxy.$urls.m().getuserinfo, obj_info).get()
      console.log(info_res.data)

      // var login_url = "http://localhost:8888/user/login";
      // var login_data = {}
      // axios({
      //   method: 'post',
      //   url: login_url,
      //   data: obj,
      //   headers: {
      //     'Content-Type':'application/json'
      //   }
      // })
      // .then(res => {
      //   login_data = res.data
      // })

      // console.log(login_data)

      // //草，他妈的，这里get不能传json，草他妈的
      // const obj_info = {"username":form.username}
      // var info_url = "http://localhost:8888/user/getuserinfo";
      // var info_data = {}
      // axios({
      //   method: 'get',
      //   url: info_url,
      //   //get用params post用data
      //   params: obj_info,
      //   headers: {
      //     'Content-Type':'application/json'
      //   }
      //         }).then(res => {
      //             info_data = res.data
      // })
      // console.log(info_data)
      if (msg != "登录成功") { //不是登录成功
        ElMessage({
          message: msg,
          type: 'warning',
        })
      } else { //登录成功，跳转
        ElMessage({
          message: msg,
          type: 'success',
        })
       // localStorage.setItem('token', res.data.jwt_token.access_token)
        //const info_obj = {"username":form.username}
        //.log(info_obj)
        //const info = await new proxy.$request(proxy.$urls.m().getuserinfo, {"username": "xminao2"}).modeget()
        //console.log(info)
        //延迟三秒刷新
        setTimeout(()=>{
          router.push({
            path: '/type'
         })
        }, 2000)
        setTimeout(()=>{
          location.reload()
        }, 2000)
        }
    }
    
    return {login, onSubmit, form, signin, logindialog, formLabelWidth}
  },
  
  name: 'HomeView',
};
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
    background-color: #4298B4;
    color: rgb(221, 6, 6);text-align: center;line-height: 0px;
  }
</style>
