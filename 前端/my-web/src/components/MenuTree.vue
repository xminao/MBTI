<template v-slot:title>
  <div class="menus">
    <p v-for="v in navMenus" :key="v.path">
      <el-menu>
        <el-menu-item v-if="!v.children" :index="v.path" :route="v.path">
          <span>{{ v.title }}</span>
        </el-menu-item> 
        <el-sub-menu v-if="v.children" :key="v.path" :index="v.path">
          <span>{{ v.title }}</span>
          <nav-menu :navMenus="v.children"></nav-menu>
         </el-sub-menu>
      </el-menu>
    </p>
  </div>
</template>
 
<script>

const data = [
    {
        label: '计算机',
        children: [
            {
                label: '2021级',
                children: [
                    {
                        label: '软件工程',
                    }
                ]
            },
            {
                label: '2020级',
                children: [
                    {
                        label: '网络工程',
                    }
                ]
            }
        ]
    },
]
</script>