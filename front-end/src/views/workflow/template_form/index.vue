<template>
    <div>
        <n-grid :gutter="[24, 1]">
            <n-gi :span="12">
                <n-card>
                    <p>模板</p>
                    <n-button type="primary">
                        添加模板
                    </n-button>
                    <vxe-grid ref='xGrid' v-bind="gridOptions" v-on="gridEvent">
                        <template #action="{ row }">
                            <div>
                                <n-button type="primary">删除</n-button>
                                <n-button type="primary">编辑</n-button>
                                <n-button type="primary" @click="loadTmplForm(row)">表单控件</n-button>
                            </div>
                        </template>
                        <template #dept="{ row }">
                            <div>
                                {{ row.Dept.id == 0 ? "未分配" : row.Dept.dept_name }}
                            </div>
                        </template>
                    </vxe-grid>
                </n-card>
            </n-gi>
            <n-gi :span="12">
                <n-card>
                    <p>部门</p>
                </n-card>
                <n-card>

                </n-card>
            </n-gi>
        </n-grid>
    </div>
</template>

<script setup lang="ts">
const { loadTemplates, gridOptions } = useTemplate();
const { loadTemplateForm } = useTemplateForm();
const router = useRouter()
const xGrid = ref()
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
const loadTmplForm = async (row) => {
    console.log(33)
    // await loadTemplateForm(row.id)
}
</script>

<style scoped></style>