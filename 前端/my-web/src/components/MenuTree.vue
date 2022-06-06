<!-- <template v-for="item in a" :key="item">
    <el-sub-menu :index="item.college_id">
        <template #title>{{item.college_name}}</template>
    </el-sub-menu>
</template> -->
<template>
    <el-sub-menu>哈哈哈哈</el-sub-menu>
</template>

<script>
import { getCurrentInstance, reactive, ref } from 'vue'

export default {

  name: "tree-list",
  //接受传参
  emits: ['year'],
  props: {
      "data": Array
  },
  setup(props, context) {


      let a = props.data
      const {proxy} = getCurrentInstance()

    const func=async()=> {
        let listres = await new proxy.$request(proxy.$urls.m().getyearlist, {"college_id": 1}).get()
        let datas = listres.college_list
        context.emit(datas)
    } 

    func()

      return {
        a
      }
  }
}
</script>
