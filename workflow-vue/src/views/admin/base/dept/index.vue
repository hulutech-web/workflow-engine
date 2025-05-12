<template>
    <div>
        <a-row :gutter="[24, 1]">
            <a-col :span="12">
                <a-card>
                    <p>部门</p>
                    <a-button type="primary" :icon="h(PlusOutlined)">
                        新增部门
                    </a-button>
                    <a-table bordered :columns="columns" :dataSource="depts" :pagination="false">
                        <template #bodyCell="{ column, text, record }">
                            <template v-if="column.dataIndex === 'html'">
                                <span class="text-blue-500 text-xl mr-3">
                                    {{ record.html }}
                                </span>
                                {{ record.dept_name }}
                            </template>
                            <template v-if="column.dataIndex === 'Manager'">
                                {{ record.Manager ? record.Manager.name : '' }}

                                <a-button type="dashed" size="small" @click="bindManager(record)" danger>设置经理</a-button>
                            </template>
                            <template v-if="column.dataIndex === 'Director'">
                                {{ record.Director ? record.Director.name : '' }}

                                <a-button type="primary" size="small" @click="bindDirector(record)">设置主管</a-button>
                            </template>
                            <template v-if="column.dataIndex === 'action'">
                                <a-space>
                                    <a-button type="primary">修改</a-button>
                                    <a-button type="dashed" danger>删除</a-button>
                                </a-space>
                            </template>
                        </template>
                    </a-table>

                    <a-modal :footer="false" v-model:open="open" width="1000px" title="用户" centered @close="onClose"
                        :bodyStyle="{ height: '800px' }">
                        <Emplist @bind="bindins" />
                    </a-modal>
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
import { ConsoleSqlOutlined, PlusOutlined } from '@ant-design/icons-vue';
const { loadDepts, setManager, setDirector } = useDept();
const depts = ref([])
const columns = [
    {
        title: '层级',
        dataIndex: 'html',
        key: 'html',
    },
    {
        title: '经理',
        dataIndex: 'Manager',
        key: 'Manager',
    },
    {
        title: '主管',
        dataIndex: 'Director',
        key: 'Director',
    },
    {
        title: '操作',
        dataIndex: 'action',
        key: 'action',
    },


]
const init = async () => {
    let d_Data = await loadDepts()
    depts.value = d_Data.data
}
init()
const open = ref(false)
const bindDirector = (record: any) => {
    open.value = true
    console.log(record)
    directorState.value.dept_id = record.id
}
const bindManager = (record: any) => {
    open.value = true
    managerState.value.dept_id = record.id
}
const managerState = ref({
    dept_id: null,
    manager_id: null
})
const directorState = ref({
    director_id: null,
    dept_id: null
})
const onClose = () => {
    managerState.value.dept_id = null
    directorState.value.dept_id = null
    open.value = false
}

const bindins = async (val) => {
    if (directorState.value.dept_id) {
        directorState.value.director_id = val.emp_id
        await setDirector(directorState.value)
    }else{
        managerState.value.manager_id = val.emp_id
        await setManager(managerState.value)
    }
}
</script>

<style scoped></style>