<template>
    <el-table :data="tableData">
    <el-table-column prop="date" label="Date" width="140" />
    <el-table-column prop="name" label="Name" width="120" />
    <el-table-column prop="address" label="Address" />
    </el-table>
</template>

<script>
import { defineComponent, ref } from 'vue'

export default defineComponent({
    setup() {
        const item = {
            date: '2016-05-02',
            name: 'Tom',
            address: 'No. 189, Grove St, Los Angeles',
        }
        const tableData = ref(Array.from({ length: 20 }).fill(item))
        return {tableData}
    },
})
</script>

<style scoped lang="less">
  #app {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
  }
</style>

