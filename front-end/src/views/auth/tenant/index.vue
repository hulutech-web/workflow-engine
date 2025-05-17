<script setup lang="tsx">
import type { DataTableColumns, FormInst } from 'naive-ui'

import { useBoolean } from '@/hooks'
import { deleteTenant, fetchTenantList} from '@/service'
import { NButton, NPopconfirm, NSpace } from 'naive-ui'
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

async function delTenant(id: number) {
  await deleteTenant({id: id})
  getTenantList()
}

const columns: DataTableColumns<Auth.Tenant> = [
  {
    title: '名称',
    align: 'center',
    key: 'name',
  },
  {
    title: '域名',
    align: 'center',
    key: 'domain',
  },
  {
    title: '邮箱',
    align: 'center',
    key: 'email',
  },
  {
    title: '描述',
    align: 'center',
    key: 'description',
  },
  {
    title: '创建时间',
    align: 'center',
    key: 'created_at',
  },
  {
    title: '操作',
    align: 'center',
    key: 'actions',
    render: (row) => {
      return (
        <NSpace justify="center">
          <NButton
            v-if="row.id !== 1"
            size="small"
            onClick={() => modalRef.value.openModal('edit', row)}
          >
            编辑
          </NButton>
          <NPopconfirm onPositiveClick={() => delTenant(row.id!)}>
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
const listData = ref<Auth.Tenant[]>([])

async function getTenantList() {
  startLoading()
  const res =  await fetchTenantList(model.value)
  if (res.code === 200) {
    listData.value = res.data.lists
    count.value = res.data.count
  } else {
    window.$message.error(res.message)
  }
  endLoading()
}

onMounted(() => {
  getTenantList()
})

function changePage(page: number, size: number) {
  model.value.pageNo = page
  model.value.pageSize = size
  getTenantList()
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
              <NButton type="primary" @click="getTenantList">
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
            新建租户
          </NButton>
        </template>
        <NSpace vertical>
          <n-data-table :columns="columns" :data="listData" :loading="loading" />
          <Pagination :count="count" @change="changePage" />
        </NSpace>
        <PermModal ref="permRef" modal-name="权限" />
        <TableModal ref="modalRef" modal-name="租户" />
      </n-card>
    </NSpace>
  </n-flex>
</template>
