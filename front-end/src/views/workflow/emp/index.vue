<template>
    <div>
        <a-card title="员工管理">
            <template #extra>
                <div>
                    <a-button type="primary" @click="addEmp">创建员工</a-button>
                </div>
            </template>
            <vxe-grid ref='xGrid' v-bind="gridOptions" v-on="gridEvent">
                <template #action="{ row }">
                    <div>
                        <a-button type="primary">删除</a-button>
                        <a-button type="primary">编辑</a-button>
                        <a-button type="primary" @click="bind(row)">绑定用户</a-button>
                    </div>
                </template>
                <template #dept="{ row }">
                    <div>
                        {{ row.Dept.id == 0 ? "未分配" : row.Dept.dept_name }}
                    </div>
                </template>
            </vxe-grid>
            <a-modal :footer="false" v-model:open="open" width="1000px" title="用户" centered
                :bodyStyle="{ height: '800px' }">
                <Userlist @bind="bindins" />
            </a-modal>
        </a-card>
    </div>
</template>

<script setup lang="ts">
import { message } from 'ant-design-vue';


const { gridOptions, bindUser } = useEmp()
const router = useRouter()
const rulesStore = useRulesStore()

// TABLE
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
const open = ref(false)
const addEmp = () => {
    router.push({ path: "/admin/manage/groupcourse/crt" })
}
const formRef = ref()

const state = ref({
    user_id: null,
    emp_id: null
})
const bindins = async (val) => {
    state.value.user_id = val.user_id
    if (state.value.user_id == null || state.value.emp_id == null) {
        return message.error("为选中用户")
    }
    await bindUser(state.value)
}


const onSubmit = async () => {
    try {
        //先清空一下验证
        formRef.value.clearValidate()
    } catch (error) {
        formRef.value
            .validate()
            .then(() => {

            })
            .catch(error => {
                if (error) {
                    return;
                }
            });
    }
}


const bind = (val) => {
    console.log(val)
    open.value = true
    state.value.emp_id = val.id
}
</script>

<style scoped></style>