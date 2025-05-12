<template>
    <div>
        <a-row :gutter="[24, 1]">
            <a-col :span="12">
                <a-card>
                    <p>模板</p>
                    <a-button type="primary" :icon="h(PlusOutlined)">
                        添加模板
                    </a-button>
                    <vxe-grid ref='xGrid' v-bind="gridOptions" v-on="gridEvent">
                        <template #action="{ row }">
                            <div>
                                <a-button type="primary">删除</a-button>
                                <a-button type="primary">编辑</a-button>
                                <a-button type="primary" @click="loadTmplForm(row)">表单控件</a-button>
                            </div>
                        </template>
                        <template #dept="{ row }">
                            <div>
                                {{ row.Dept.id == 0 ? "未分配" : row.Dept.dept_name }}
                            </div>
                        </template>
                    </vxe-grid>
                </a-card>
            </a-col>
            <a-col :span="12">
                <a-card>
                    <p>部门</p>
                </a-card>
                <a-card>

                </a-card>
            </a-col>
        </a-row>
    </div>
</template>

<script setup lang="ts">
import { h } from 'vue';
import { PlusOutlined } from '@ant-design/icons-vue';
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