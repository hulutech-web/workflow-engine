<template>
    <div>
        <h1>用户列表</h1>
        <vxe-grid ref='xGrid' v-bind="gridOptions" v-on="gridEvent">
            <template #action="{ row }">
                <div>
                    <a-button type="primary" @click="bindUser(row)">绑定</a-button>
                </div>
            </template>
            <template #dept="{ row }">
                <div>
                    {{ row.Dept.id == 0 ? "未分配" : row.Dept.dept_name }}
                </div>
            </template>
        </vxe-grid>
    </div>
</template>

<script setup lang='ts'>
const xGrid = ref()
const { gridOptions } = useUser()
const emit = defineEmits(['bind'])
const gridEvent: VxeGridListeners<RowVO> = {
    proxyQuery() {
        console.log('数据代理查询事件')
        const grid = xGrid.value
        // 获取表格中的数据
        const data = grid.getTableData().fullData
    },
    proxyDelete() {
        console.log('数据代理删除事件')
    },
    proxySave() {
        console.log('数据代理保存事件')
    }
}

const bindUser = (user) => {
    emit('bind', {user_id:user.id})
}


</script>

<style></style>