<template>
    <div>
        <a-card title="用户管理">
            <template #extra>
                <div>
                    <a-button type="primary" @click="addEmp">创建用户</a-button>
                </div>
            </template>
            <vxe-grid ref='xGrid' v-bind="gridOptions" v-on="gridEvent">
                <template #action="{ row }">
                    <div>
                        <a-button type="primary">删除</a-button>
                        <a-button type="primary">编辑</a-button>
                        
                    </div>
                </template>
                <template #dept="{ row }">
                    <div>
                        {{ row.Dept.id == 0 ? "未分配" : row.Dept.dept_name }}
                    </div>
                </template>
            </vxe-grid>
        </a-card>
    </div>
</template>

<script setup lang="ts">
const { gridOptions } = useUser()
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

const addEmp = () => {
    router.push({ path: "/admin/manage/groupcourse/crt" })
}
const formRef = ref()

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

const formModalRef = ref()

</script>

<style scoped></style>