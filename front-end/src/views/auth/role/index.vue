<script setup lang="tsx">
import type { DataTableColumns, FormInst } from 'naive-ui'

import { useBoolean } from '@/hooks'
import {changeRole, deleteRole, fetchRoleList} from '@/service'
import { NButton, NPopconfirm, NSpace, NSwitch, NTag } from 'naive-ui'
import TableModal from './components/TableModal.vue'

const { bool: loading, setTrue: startLoading, setFalse: endLoading } = useBoolean(false)

const initialModel = {
  pageNo: 1,
  pageSize: 10,
}
const model = ref({ ...initialModel })
function handleResetSearch() {
  model.value = { ...initialModel }
}

const formRef = ref<FormInst | null>()
const modalRef = ref()

async function delRole(id: number) {
  await deleteRole({id: id})
  getRoleList()
}

const columns: DataTableColumns<Auth.Role> = [
  {
    title: '名称',
    align: 'center',
    key: 'name',
  },
  {
    title: '成员',
    align: 'center',
    key: 'member',
  },
  {
    title: '是否启用',
    align: 'center',
    key: 'is_disable',
    render: (row) => {
      return (
        <NSwitch
          value={row.is_disable}
          checked-value={0}
          unchecked-value={1}
          onUpdateValue={(value: 0 | 1) =>
            handleUpdateDisabled(row.id!)}
        >
          {{ checked: () => '启用', unchecked: () => '禁用' }}
        </NSwitch>
      )
    },
  },
  {
    title: '操作',
    align: 'center',
    key: 'actions',
    render: (row) => {
      return (
        <NSpace justify="center">
          <NButton
            size="small"
            onClick={() => modalRef.value.openModal('edit', row)}
          >
            编辑
          </NButton>
          <NPopconfirm onPositiveClick={() => delRole(row.id!)}>
            {{
              default: () => '确认删除',
              trigger: () => <NButton size="small" type="error">删除</NButton>,
            }}
          </NPopconfirm>
        </NSpace>
      )
    },
  },
]

const count = ref(0)
const listData = ref<Auth.Role[]>([])

async function getRoleList() {
  startLoading()
  const res =  await fetchRoleList(model.value)
  if (res.code === 200) {
    listData.value = res.data.lists
    count.value = res.data.count
  } else {
    window.$message.error(res.message)
  }
  endLoading()
}

onMounted(() => {
  getRoleList()
})

function changePage(page: number, size: number) {
  model.value.pageNo = page
  model.value.pageSize = size
  getRoleList()
}

async function handleUpdateDisabled(id: number){
  await changeRole({id: id})
  getRoleList()
}

</script>

<template>
  <n-flex>
    <NSpace vertical class="flex-1">
      <n-card>
        <n-form ref="formRef" :model="model" label-placement="left" inline :show-feedback="false">
          <n-flex>
            <n-form-item label="姓名" path="condition_1">
              <n-input v-model:value="model.condition_1" placeholder="请输入" />
            </n-form-item>
            <n-form-item label="性别" path="condition_2">
              <n-input v-model:value="model.condition_2" placeholder="请输入" />
            </n-form-item>
            <n-flex class="ml-auto">
              <NButton type="primary" @click="getRoleList">
                <template #icon>
                  <icon-park-outline-search />
                </template>
                搜索
              </NButton>
              <NButton strong secondary @click="handleResetSearch">
                <template #icon>
                  <icon-park-outline-redo />
                </template>
                重置
              </NButton>
            </n-flex>
          </n-flex>
        </n-form>
      </n-card>

      <n-card class="flex-1">
        <template #header>
          <NButton type="primary" @click="modalRef.openModal('add')">
            <template #icon>
              <icon-park-outline-add-one />
            </template>
            新建角色
          </NButton>
        </template>
        <NSpace vertical>
          <n-data-table :columns="columns" :data="listData" :loading="loading" />
          <Pagination :count="count" @change="changePage" />
        </NSpace>

        <TableModal ref="modalRef" modal-name="角色" />
      </n-card>
    </NSpace>
  </n-flex>
</template>
