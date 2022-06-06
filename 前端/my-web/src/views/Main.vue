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
          <el-menu-item index="/home"><List style="width: 1em; height: 1em; margin-right: 8px" />人格测试</el-menu-item>
          <el-menu-item> </el-menu-item>
          <el-menu-item index="/question"><Menu style="width: 1em; height: 1em; margin-right: 8px" />人格类型</el-menu-item>
          <el-menu-item> </el-menu-item>
          <el-menu-item index="/user"><Promotion style="width: 1em; height: 1em; margin-right: 8px" />个人中心</el-menu-item>
          </el-menu>
          <el-button
            v-show="!login_status"
            id="sign" 
            type="info"
            color="#88619A"
            @click="dialog = true" >登录/注册</el-button>
          <el-button
            v-show="login_status"
            id="logout" 
            type="info"
            color="#66CCCC"
            @click="logout()" > 登出 </el-button>

            <el-dialog 
              v-model="dialog" 
              title="登录界面" 
              center="true"
              destroy-on-close
              width=35%>
              <el-form 
              :model="login_form">
                <el-form-item 
                  label="用户名" 
                  :label-width="formLabelWidth"
                  size="large">
                  <el-input 
                    :prefix-icon="User"
                    placeholder="请输入用户名"
                    v-model="login_form.username" 
                    autocomplete="off"
                    style="width:250px;"/>
                </el-form-item>
                <el-form-item label="密　码" :label-width="formLabelWidth" style="font-size: 30px">
                  <el-input 
                    :prefix-icon="Lock"
                    placeholder="请输入密码"
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
                    color="#00CC99"
                    >登录</el-button>
                <span>　　　</span>
                 <el-button
                    id="register"
                    type="primary"
                    @click="open_reg_dialog()"
                    color="#33CCCC"
                    >注册</el-button>
                </span>
              </template>
          </el-dialog>


            <el-dialog 
              v-model="reg_dialog" 
              title="注册界面" 
              center="true"
              destroy-on-close
              width=35%>
              <el-form 
              :model="login_form">
                <el-form-item 
                  label="用户名" 
                  :label-width="formLabelWidth"
                  size="large">
                  <el-input 
                    :prefix-icon="User"
                    placeholder="请输入用户名"
                    maxlength="8"
                    show-word-limit
                    v-model="register_form.username" 
                    autocomplete="off"
                    style="width:250px;"/>
                </el-form-item>
                <el-form-item 
                  label="昵　称" 
                  :label-width="formLabelWidth" 
                  style="font-size: 30px">
                  <el-input 
                    :prefix-icon="Star"
                    maxlength="8"
                    show-word-limit
                    placeholder="请输入昵称"
                    v-model="register_form.nickname" 
                    autocomplete="off"
                    size="large"
                    style="width:250px;" />
                </el-form-item>
                <el-form-item 
                label="密　码" 
                :label-width="formLabelWidth" 
                style="font-size: 30px">
                  <el-input 
                    :prefix-icon="Lock"
                    placeholder="请输入密码"
                    v-model="register_form.passwd" 
                    autocomplete="off"
                    size="large"
                    style="width:250px;" />
                </el-form-item>
                <el-form-item 
                label="性　别" 
                :label-width="formLabelWidth" 
                style="font-size: 30px">
                <el-select v-model="register_form.gender" class="m-2" placeholder="选择性别" size="large" suffix-icon="CaretBottom">
                    <el-option
                    v-for="item in [{value: '男', label: '♂'}, {value: '女', label: '♀'}]"
                    :key="item.value"
                    :value="item.value"
                    :label="item.value"
                    >
                    <span style="float: left">{{item.value}}</span>
                    <span style="float: right">{{item.label}}</span>
                    </el-option>
                </el-select>
                </el-form-item>
                <el-form-item 
                label="学　号" 
                :label-width="formLabelWidth" 
                style="font-size: 30px">
                  <el-input 
                    :prefix-icon="Postcard"
                    maxlength="10"
                    show-word-limit
                    placeholder="请输入学号"
                    v-model="register_form.studentId" 
                    autocomplete="off"
                    size="large"
                    style="width:250px;" />
                </el-form-item>
                <el-form-item 
                label="姓　名" 
                :label-width="formLabelWidth" 
                style="font-size: 30px">
                  <el-input 
                    :prefix-icon="UserFilled"
                    maxlength="10"
                    show-word-limit
                    placeholder="请输入姓名"
                    v-model="register_form.name" 
                    autocomplete="off"
                    size="large"
                    style="width:250px;" />
                </el-form-item>
              </el-form>
              <template #footer>
                <span class="dialog-footer">
                 <el-button
                    id="register"
                    type="primary" 
                    @click="register()"
                    color="#33CCCC"
                    >注册</el-button>
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
import { User, Lock, Star, CaretBottom, Postcard, UserFilled } from '@element-plus/icons-vue'

export default {
    setup () {

        //路由
        const router = useRouter()
        const {proxy} = getCurrentInstance()

        let dialog = ref(false)
        let reg_dialog = ref(false)

        //判断用户是否登录
        let login_status = ref(false)
        
        if (localStorage.getItem('token')) {
          login_status.value = true
        }

        const formLabelWidth = "90px"

        //登录注册表单
        let login_form = reactive({
            username: '',
            passwd: '',
        })
        const register_form = reactive({
            username: '',
            nickname: '',
            passwd: '',
            gender: '',
            studentId: '',
            name: ''
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
            login_status.value = true
            close_dialog()
        }
        //注册函数
        const register = async()=> {
            if (register_form.username.length == 0) {
                ElMessage({
                message: "用户名不能为空",
                type: 'warning',
                })
                return
            }
            if (register_form.nickname.length == 0) {
                ElMessage({
                message: "昵称不能为空",
                type: 'warning',
                })
                return
            }
            if (register_form.passwd.length == 0) {
                ElMessage({
                message: "密码不能为空",
                type: 'warning',
                })
                return
            }
            if (register_form.gender.length == 0) {
                ElMessage({
                message: "请选择性别",
                type: 'warning',
                })
                return
            }
            if (register_form.studentId.length == 0) {
                ElMessage({
                message: "学号不能为空",
                type: 'warning',
                })
                return
            }
           if (register_form.name.length == 0) {
                ElMessage({
                message: "姓名不能为空",
                type: 'warning',
                })
                return
            }

            const obj = {"username":register_form.username, "nickname":register_form.nickname, "password":register_form.passwd, "gender":register_form.gender, "binding_student_id":register_form.studentId}
            const res = await new proxy.$request(proxy.$urls.m().register, obj).post()
            console.log(res)
            ElMessage({
                message: "注册成功，" + res.data.nickname,
                type: 'success',
                center: true,
            })
            close_reg_dialog()
        }

        let close_dialog=()=> {
            login_form.username = ''
            login_form.passwd = ''
            dialog.value = false
        }

        let close_reg_dialog=()=> {
          register_form.username = ''
          register_form.nickname = ''
          register_form.passwd = ''
          register_form.gender = ''
          reg_dialog.value = false
        }

        let open_reg_dialog=()=> {
            dialog.value = false;
            reg_dialog.value = true;
        }

        let logout=()=> {
          login_status.value = false
          localStorage.setItem('token', '')
           ElMessage({
                message: "成功登出",
                type: 'success',
                center: true,
            })
        }

        //验证规则
        // const validatePass = (rule, value, callback) => {
        //    //  密码只能由大小写英文字母或数字开头，且由大小写英文字母_.组成
        //   const reg = /^[A-Za-z0-9][A-Za-z0-9_.]{5,14}$/
        //   console.log('reg', value.match(reg))
        //   if (!value.match(reg)) {
        //     callback(new Error('密码由字母或数字开头，且只能为字母,数字,下划线及().)'))
        //   } else {
        //     callback()
        //   }
        // }

        return { 
            User,
            Lock,
            CaretBottom,
            Star,
            UserFilled,
            Postcard,
            login_status,
            logout,
            close_dialog, 
            open_reg_dialog,
            router, 
            formLabelWidth, 
            dialog, 
            reg_dialog,
            login_form, 
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
    #sign {
        color: #fff;
        font-size: 20px;
        height: 50px;
        width: 10%;
        border-radius: 50px
    }
    #logout {
        color: #fff;
        font-size: 20px;
        height: 50px;
        width: 10%;
        border-radius: 50px
    }

    #login {
        color: #fff;
        font-size: 20px;
        height: 50px;
        width: 35%;
        border-radius: 50px
    }
    #register {
        color: #fff;
        font-size: 20px;
        height: 50px;
        width:35%;
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
        //background-color: #4298B4;
        background-color: #fff;
        padding: 0%;
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