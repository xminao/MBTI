<template>
    <el-container>
        <el-aside>
              <el-menu
                default-active="aaa"
                class="el-menu-vertical-demo"
                :collapse="isCollapse"
                router=true
                unique-opened>
                    <el-sub-menu index="question">
                        <template #title>题库管理</template>
                        <el-menu-item index="/question">题库列表</el-menu-item>
                    </el-sub-menu>

                    <el-sub-menu index="management">
                        <template #title>用户管理</template>
                        <el-menu-item index="/user">用户列表</el-menu-item>
                    </el-sub-menu>

                    <el-sub-menu index="studentmanagement">
                        <template #title>校园管理</template>
                        <el-menu-item index="/university">学生列表</el-menu-item>
                    </el-sub-menu>

                    <el-sub-menu index="datas">
                        <template #title>数据统计</template>
                    </el-sub-menu>
                </el-menu>
        </el-aside>

        <el-main>
            <router-view/>
        </el-main>
    </el-container>
</template>

<script>
import { useRouter } from 'vue-router'

export default ({
    setup() {
        const router = useRouter()

        const goQues=()=> {
            router.push('/question')
        }

        return {
            goQues,
            router
        }
    },
})
</script>

<style scoped lang="less">

.el-container {
    width: 100%;
    height: 100%;
    background-color: #fff;
}

.el-aside {
    padding-top: 50px;
    width: 20%;
    background-color: #fff;
}

.el-main {
    width: 100%;
}


</style>