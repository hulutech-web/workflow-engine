<script setup lang="ts">
import { useBoolean } from '@/hooks'
import {addTenant, editTenant, fetchAllRoutes} from '@/service'

interface Props {
  modalName?: string
}

interface Option {
  label: string
  key: number
  menuType: string | number
  children?: Option[]
}

const {
  modalName = '',
} = defineProps<Props>()

const emit = defineEmits<{
  open: []
  close: []
}>()

const { bool: modalVisible, setTrue: showModal, setFalse: hiddenModal } = useBoolean(false)

const { bool: submitLoading, setTrue: startLoading, setFalse: endLoading } = useBoolean(false)
const menuOptions = shallowRef<Option[]>([])

const formDefault: Auth.TenantReq = {
  id: 0,
  name: '',
  address: '',
  phone: '',
  email: '',
  domain: '',
  logo: '',
  description: '',
  is_disable: 0,
  expired_at: 0,
  menus: [],
}
const formModel = ref<Auth.TenantReq>({ ...formDefault })

type ModalType = 'add' | 'view' | 'edit'
const modalType = shallowRef<ModalType>('add')
const modalTitle = computed(() => {
  const titleMap: Record<ModalType, string> = {
    add: '添加',
    view: '查看',
    edit: '编辑',
  }
  return `${titleMap[modalType.value]}${modalName}`
})

async function openModal(type: ModalType = 'add', data: any) {
  emit('open')
  modalType.value = type
  showModal()
  formRef.value?.resetFields()
  const handlers = {
    async add() {
      formModel.value = { ...formDefault }
    },
    async view() {
      if (!data)
        return
      formModel.value = { ...data }
    },
    async edit() {
      if (!data)
        return
      await getMenuOptions()
      formModel.value = { ...data }
    },
  }
  await handlers[type]()
}

function closeModal() {
  hiddenModal()
  endLoading()
  emit('close')
}

defineExpose({
  openModal,
})

const formRef = ref()
async function submitModal() {
  const handlers = {
    async add() {
      return new Promise(async (resolve) => {
        const res = await addTenant(formModel.value)
        if (res.code === 200) {
          window.$message.success('租户添加成功')
          resolve(true)
        } else {
          resolve(true)
        }
      })
    },
    async edit() {
      return new Promise(async (resolve) => {
        const res = await editTenant(formModel.value)
        if (res.code === 200) {
          window.$message.success('租户编辑成功')
          resolve(true)
        } else {
          resolve(true)
        }
      })
    },
    async view() {
      return true
    },
  }
  await formRef.value?.validate()
  startLoading()
  await handlers[modalType.value]() && closeModal()
}

const rules = {
  userName: {
    required: true,
    message: '请输入用户名',
    trigger: 'blur',
  },
}


async function getMenuOptions() {
  menuOptions.value = []
  const { data } = await fetchAllRoutes()
  const options: Option[] = []

// 递归查找子节点函数
  const findChildren = (parentId: number): Option[] => {
    return data
      .filter(item => item.pid === parentId) // 匹配父节点ID
      .map(child => ({
        label: child.title,
        key: child.id,
        menuType: child.menuType,
        children: findChildren(child.id) // 递归查找嵌套子节点
      }))
  }

// 只处理顶级菜单 (pid=0)
  data.forEach((item) => {
    if (item.pid === 0) {
      const option: Option = {
        label: item.title,
        key: item.id,
        menuType: item.menuType,
        children: findChildren(item.id) // 初始化递归查找
      }
      options.push(option)
    }
  })
  menuOptions.value = options
}

const selected = ref<number[]>([])

function handleMenuChange(value: number[]) {
  formModel.value.menus = value
}
</script>

<template>
  <n-modal
    v-model:show="modalVisible"
    :mask-closable="false"
    preset="card"
    :title="modalTitle"
    class="w-700px"
    :segmented="{
      content: true,
      action: true,
    }"
  >
    <n-form ref="formRef" :rules="rules" label-placement="left" :model="formModel" :label-width="100" :disabled="modalType === 'view'">
      <n-grid :cols="2" :x-gap="18">
       <n-form-item-grid-item :span="1" label="租户名称" path="name">
          <n-input v-model:value="formModel.name" placeholder="请输入租户名称" />
        </n-form-item-grid-item>
        <n-form-item-grid-item :span="1" label="租户地址" path="address">
          <n-input v-model:value="formModel.address" placeholder="请输入租户地址" />
          </n-form-item-grid-item>
        <n-form-item-grid-item :span="1" label="联系电话" path="phone">
          <n-input v-model:value="formModel.phone" placeholder="请输入联系电话" />
        </n-form-item-grid-item>
        <n-form-item-grid-item :span="1" label="邮箱" path="email">
          <n-input v-model:value="formModel.email" placeholder="请输入邮箱" />
        </n-form-item-grid-item>
        <n-form-item-grid-item :span="1" label="域名" path="domain">
          <n-input v-model:value="formModel.domain" placeholder="请输入域名" />
        </n-form-item-grid-item>
        <n-form-item-grid-item :span="1" label="描述" path="description">
          <n-input v-model:value="formModel.description" placeholder="请输入描述" />
        </n-form-item-grid-item>
        <n-form-item-grid-item :span="1" label="到期时间" path="expired_at">
          <n-date-picker
            v-model:value="formModel.expired_at"
            type="datetime"
            placeholder="请选择到期时间"
          />
        </n-form-item-grid-item>
<!--        <n-form-item-grid-item :span="1" label="Logo" path="logo">-->
<!--          <n-upload-->
<!--            v-model:value="formModel.logo"-->
<!--            :default-file-list="[formModel.logo]"-->
<!--            :before-upload="beforeUpload"-->
<!--            :max-size="10 * 1024"-->
<!--            :multiple="false"-->
<!--            :show-file-list="false"-->
<!--            :action="''"-->
<!--            :disabled="modalType === 'view'"-->
<!--          >-->
<!--            <template #default="{ file, onRemove }">-->
<!--              <div class="upload-btn">-->
<!--                <n-icon :size="24" :type="file ? 'file-done' : 'upload-cloud'"></n-icon>-->
<!--                <div class="upload-text">-->
<!--                  <span v-if="!file">上传Logo</span>-->
<!--                  <span v-else>{{ file.name }}</span>-->
<!--                </div>-->
<!--              </div>-->
<!--            </template>-->
<!--          </n-upload>-->
<!--        </n-form-item-grid-item>  -->
        <n-form-item-grid-item :span="1" label="是否启用" path="is_disable">
          <n-switch
            v-model:value="formModel.is_disable"
            :checked-value="0" :unchecked-value="1"
          >
            <template #checked>
              启用
            </template>
            <template #unchecked>
              禁用
            </template>
          </n-switch>
        </n-form-item-grid-item>
        <n-form-item-grid-item :span="2" label="菜单权限" path="menus">
          <n-tree-select
            multiple
            checkable
            filterable
            cascade
            :clear-filter-after-select="false"
            :options="menuOptions"
            v-model:value="selected"
            @update-value="handleMenuChange"
            clearable
          />
        </n-form-item-grid-item>
      </n-grid>
    </n-form>
    <template #action>
      <n-space justify="center">
        <n-button @click="closeModal">
          取消
        </n-button>
        <n-button type="primary" :loading="submitLoading" @click="submitModal">
          提交
        </n-button>
      </n-space>
    </template>
  </n-modal>
</template>
