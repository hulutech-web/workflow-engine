<template>
    <div>
        <n-row :gutter="[24, 1]">
            <n-col :span="12">
                <n-card>
                    <p>部门</p>
                    <n-button type="primary">
                        新增部门
                    </n-button>
                  <vxe-grid ref='xGrid' v-bind="gridOptions" v-on="gridEvent">
                    <template #Manager="{ row }">
                      <div>
                       {{row.Manager?row.Manager.name:''}}
                      </div>
                    </template>
                    <template #Director="{ row }">
                      <div>
                        {{row.Director?row.Director.name:''}}
                      </div>
                    </template>
                    <template #action="{ row }">
                      <div>
                        <n-button type="primary">删除</n-button>
                        <n-button type="primary">编辑</n-button>
                      </div>
                    </template>
                  </vxe-grid>

                    <n-modal  v-model:show="open" width="1000px" title="用户"
                        :bodyStyle="{ height: '800px' }">
                      <n-card style="width:600px">
                        <Emplist @bind="bindins" />
                      </n-card>
                    </n-modal>
                </n-card>
            </n-col>
            <n-col :span="12">
                <n-card>
                    <p>部门</p>
                </n-card>
                <n-card>

                </n-card>
            </n-col>
        </n-row>
    </div>
</template>

<script setup lang="ts">
import { h } from 'vue';
const { loadDepts, setManager, setDirector,gridOptions } = useDept();
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
</script>

<style scoped></style>
