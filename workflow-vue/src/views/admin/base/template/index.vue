<template>
    <div>
        <a-row :gutter="[24, 1]">
            <a-col :span="15">
                <a-card class="mb-3">
                    <a-form :model="templateState" ref="templateRef" layout="inline">
                        <a-form-item label="模板名称" name="template_name"
                            :rules="[rulesStore.getRule('template_name') ? rulesStore.getRule('template_name') : { required: false }]">
                            <a-input v-model:value="templateState.template_name"></a-input>
                        </a-form-item>
                        <a-form-item>
                            <a-button type="primary" @click="submitTemplate">添加</a-button>
                        </a-form-item>
                    </a-form>
                    <vxe-grid ref='xGrid' v-bind="gridOptions" v-on="gridEvent">
                        <template #action="{ row }">
                            <div>
                                <a-button-group>
                                    <a-popconfirm title="将同步删除模板字段，确认删除吗？" ok-text="是" cancel-text="点错了"
                                        @confirm="deleteTemplate(row)">
                                        <a-button size="small" danger type="primary">删除</a-button>
                                    </a-popconfirm>
                                    <a-button type="primary" size="small" @click="editTemplate(row)">编辑</a-button>
                                    <a-button type="primary" size="small" @click="loadTmplForm(row)">表单控件</a-button>
                                </a-button-group>
                            </div>
                        </template>
                        <template #dept="{ row }">
                            <div>
                                {{ row.Dept.id == 0 ? "未分配" : row.Dept.dept_name }}
                            </div>
                        </template>
                    </vxe-grid>
                </a-card>
                <a-modal :footer="false" v-model:open="editOpen" width="500px" title="修改模板" centered>
                    <a-form :model="templateUpdateState" ref="templateUptRef">
                        <a-form-item label="模板名称" name="template_name"
                            :rules="[rulesStore.getRule('template_name') ? rulesStore.getRule('template_name') : { required: false }]">
                            <a-input v-model:value="templateUpdateState.template_name"></a-input>
                        </a-form-item>
                        <a-form-item>
                            <a-button type="primary" @click="submitUpdateTemplate">保存</a-button>
                        </a-form-item>
                    </a-form>
                </a-modal>
                <a-card>
                    <a-table bordered :columns="columns" :dataSource="fields">
                        <template #bodyCell="{ column, record }">
                            <template v-if="column.dataIndex === 'action'">
                                <div>
                                    <a-button-group>
                                        <a-button size="small" type="primary" @click="edit(record)">编辑</a-button>
                                        <a-popconfirm title="Title" @confirm="delRecord(record)">
                                            <a-button size="small" type="primary" danger>删除</a-button>
                                        </a-popconfirm>
                                    </a-button-group>
                                </div>
                            </template>
                        </template>
                    </a-table>
                </a-card>
            </a-col>
            <a-col :span="9">
                <a-card>
                    <p>设计字段</p>
                    <Formpart ref="formpartRef" />
                </a-card>
            </a-col>

        </a-row>

        <a-modal v-model:open="open" title="控件配置" centered width="800px" :footer="false">
            <Formpart :id="cid" ref="formpartRef" />
        </a-modal>
    </div>
</template>



<script setup lang="ts">
import Formpart from './formpart.vue'
const { loadTemplates, gridOptions, storeTemplate, deleteTemplate, updateTemplate } = useTemplate();
const { loadTemplateForm, deleteTemplateForm } = useTemplateForm();
const router = useRouter()
const xGrid = ref()
import useRulesStore from '@/store/useRulesStore.ts'
const rulesStore = useRulesStore();
// 模板管理
const templateRef = ref()
const templateState = ref({
    template_name: ""
})
const templateUpdateState = ref({
    id: null,
    template_name: ""
})
const cid = ref()
const submitTemplate = async () => {
    try {
        //先清空一下验证
        // templateRef.value.clearValidate()
        await storeTemplate(templateState.value)
        // templateRef.value.resetFormField();
    } catch (error) {
        templateRef.value.validate()
    }
    // xGrid.value.commitProxy("query")
}
const editOpen = ref(false)
const editTemplate = (row) => {
    // console.log(row)
    editOpen.value = true
    templateUpdateState.value = row
}
const templateUptRef = ref()
const formpartRef = ref()
const submitUpdateTemplate = async () => {
    try {
        //先清空一下验证
        templateUptRef.value.clearValidate()
        await updateTemplate(templateUpdateState.value)
        templateUptRef.value.resetFormField();
    } catch (error) {
        templateUptRef.value.validate()
    }
    xGrid.value.commitProxy("query")
    formpartRef.value.loadTemplateOpts()
    editOpen.value = false;
}

// 模板管理

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
const columns = [
    {
        dataIndex: "id",
        title: "id",
        key: "id"
    },
    {
        dataIndex: "field_name",
        title: "控件名称",
        key: "field_name"
    },
    {
        dataIndex: "field",
        title: "字段名",
        key: "field"
    },
    {
        dataIndex: "field_type",
        title: "字段类型",
        key: "field_type"
    },
    {
        dataIndex: "sort",
        title: "排序",
        key: "sort"
    },
    {
        dataIndex: "template_id",
        title: "模板id",
        key: "template_id"
    },
    {
        dataIndex: "action",
        title: "操作",
        key: "action"
    },

]
const fields = ref([])

const loadTmplForm = async (row) => {
    const data = await loadTemplateForm(row.id)
    fields.value = data
}

const open = ref(false)
const edit = (record) => {
    cid.value = record.id
    open.value = true
}


const templateOpts = ref([])
const loadTemplateOpts = async () => {
    const { data } = await loadTemplates()
    templateOpts.value = data
}

const delRecord = async (row) => {
    await deleteTemplateForm(row.id)
}
</script>

<style scoped></style>