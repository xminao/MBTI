<template>
  <div class="common-layout">
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
            @click="dialog = true" >登录/注册</el-button>

            <el-dialog 
              v-model="dialog" 
              title="登录" 
              center="true"
              destroy-on-close
              @close="closeDialog"
              width=30%>
              <el-form 
              :model="login_form">
                <el-form-item 
                  label="用户名" 
                  :label-width="formLabelWidth"
                  size="large">
                  <el-input 
                    :prefix-icon="User"
                    v-model="login_form.username" 
                    autocomplete="off"
                    style="width:250px;"/>
                </el-form-item>
                <el-form-item label="密　码" :label-width="formLabelWidth" style="font-size: 30px">
                  <el-input 
                    :prefix-icon="Lock"
                    v-model="login_form.passwd" 
                    autocomplete="off"
                    type="password"
                    size="large"
                    style="width:250px;"
                    show-password />
                </el-form-item>
              </el-form>
              <template #footer>
                <span class="dialog-footer">
                  <el-button
                    id="login"
                    type="primary" 
                    @click="login()"
                    color="#88619A"
                    >登录</el-button>
                </span>
              </template>
          </el-dialog>

        </el-header>

        <el-main>
            <router-view/>
        </el-main>
    </el-container>
  </div>
</template>



<script>
import { useRouter } from 'vue-router';
import { getCurrentInstance, reactive, ref } from 'vue';
import { ElMessage } from 'element-plus';
import { User, Lock } from '@element-plus/icons-vue'

export default {
    setup () {

        //路由
        const router = useRouter()
        const {proxy} = getCurrentInstance()

        let dialog = ref(false)

        const formLabelWidth = "55px"

        //登录注册表单
        const login_form = reactive({
            username: '',
            passwd: '',
        })
        const register_form = reactive({
            username: '',
            nickname: '',
            passwd: '',
            gender: '',
        })

        //登录函数
        const login = async()=> {
            if (login_form.username.length == 0) {
                ElMessage({
                message: "用户名不能为空",
                type: 'warning',
                })
                return
            }
            if (login_form.passwd.length == 0) {
                ElMessage({
                message: "密码不能为空",
                type: 'warning',
                })
                return
            }

            const obj = {"username":login_form.username, "password":login_form.passwd}
            const res = await new proxy.$request(proxy.$urls.m().login, obj).post()
            console.log(res)
            ElMessage({
                message: "登录成功，" + res.data.nickname,
                type: 'success',
                center: true,
            })
            localStorage.setItem('token', res.data.jwt_token.access_token)
            close_dialog()
        }
        //注册函数
        const register = async()=> {
            
        }

        let close_dialog=()=> {
            dialog.value = false;
        }

        return { 
            User,
            Lock,
            close_dialog,  
            register_form, 
            login, 
            register }
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
        position: absolute;
        left: 0;
        right: 0;
        top: 0;
        bottom: 0;
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
        position: relative;
        left: 0;
        right: 0;
        top: 0;
        bottom: 0;
        width: 100%;
        height: 100%;
    }

    .el-footer {
        background-color: #477777;
        height: 40px;
        align-items: center;
    }
</style>