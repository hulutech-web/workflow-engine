<script setup lang='ts'>
import useRulesStore from '@/store/useRulesStore.ts'
import { watch } from 'vue'

const props = defineProps({
  id: {
    type: Number,
    default: null,
  },
})
const { storeTemplateForm, updateTemplateForm, showTemplateForm } = useTemplateForm()
const { loadTemplates } = useTemplate()
const rulesStore = useRulesStore()
const tid = ref()
const formEditRef = ref()
const inputRef = ref()
const formField = ref({
  id: null,
  field: '',
  field_default_value: '',
  field_name: '',
  field_type: '',
  field_value: [],
  sort: 0,
  template_id: '',
  field_rules: [],
})
async function loadTemp() {
  const data = await showTemplateForm(+tid.value)
  formField.value = data
}
async function loadTemplateOpts() {
  const { data } = await loadTemplates()
  templateOpts.value = data
}

watch(() => props.id, async (val) => {
  if (val) {
    tid.value = val
    await loadTemp()
  }
  else {
    await loadTemplateOpts()
  }
}, {
  deep: true,
  immediate: true,
})

const templateOpts = ref([])

const state = reactive({
  tags: [],
  inputVisible: false,
  inputValue: '',
})

function showInput() {
  state.inputVisible = true
  nextTick(() => {
    inputRef.value.focus()
  })
}

function handleEditInputConfirm() {
  const inputValue = state.inputValue
  let tags = state.tags
  if (inputValue && !tags.indexOf(inputValue) === -1) {
    tags = [...tags, inputValue]
  }
  Object.assign(state, {
    tags,
    inputVisible: false,
    inputValue: '',
  })
  formField.value.field_value = tags
}

function handleClose(removedTag: string) {
  const tags = state.tags.filter(tag => tag !== removedTag)
  state.tags = tags
  formField.value.field_value = tags
}

// 控件
async function saveField() {
  try {
    // 先清空一下验证
    formEditRef.value.clearValidate()

    if (tid.value) {
      await updateTemplateForm(formField.value)
      formEditRef.value.resetFields()
    }
    else {
      await storeTemplateForm(formField.value)
      formEditRef.value.resetFields()
    }
  }
  catch (error) {
    formEditRef.value.validate()
  }
}
defineExpose({
  loadTemplateOpts,
})
</script>

<template>
  <a-form ref="formEditRef" :model="formField">
    <a-form-item
      v-if="!tid" label="模板" name="template_id"
      :rules="[rulesStore.getRule('template_id') ? rulesStore.getRule('template_id') : { required: false }]"
    >
      <a-select v-model:value="formField.template_id">
        <a-select-option v-for="(t, m) in templateOpts" :key="m" :value="t.id">
          {{ t.template_name }}
        </a-select-option>
      </a-select>
    </a-form-item>
    <a-form-item
      label="控件名称" name="field_name"
      :rules="[rulesStore.getRule('field_name') ? rulesStore.getRule('field_name') : { required: false }]"
    >
      <a-input v-model:value="formField.field_name" />
    </a-form-item>
    <a-form-item
      label="控件类型" name="field_type"
      :rules="[rulesStore.getRule('field_type') ? rulesStore.getRule('field_type') : { required: false }]"
    >
      <a-select v-model:value="formField.field_type">
        <a-select-option value="text">
          文本框
        </a-select-option>
        <a-select-option value="number">
          数字输入
        </a-select-option>
        <a-select-option value="textarea">
          文本域
        </a-select-option>
        <a-select-option value="select">
          下拉框
        </a-select-option>
        <a-select-option value="radio">
          单选框
        </a-select-option>
        <a-select-option value="checkbox">
          复选框
        </a-select-option>
        <a-select-option value="date">
          日期框
        </a-select-option>
        <a-select-option value="file">
          文件
        </a-select-option>
      </a-select>
    </a-form-item>
    <a-form-item
      label="字段名(英文)" name="field"
      :rules="[rulesStore.getRule('field') ? rulesStore.getRule('field') : { required: false }]"
    >
      <a-input v-model:value="formField.field" />
    </a-form-item>
    <a-form-item label="控件选项" name="field_value">
      <div>
        <template v-for="(tag, index) in formField.field_value" :key="index">
          <a-tooltip :title="tag">
            <a-tag :closable="true" color="processing" @close="handleClose(tag)">
              {{ tag }}
            </a-tag>
          </a-tooltip>
        </template>
        <a-input
          v-if="state.inputVisible" ref="inputRef" v-model:value="state.inputValue" type="text"
          size="small" :style="{ width: '78px' }" @keyup.enter="handleEditInputConfirm"
        />
        <a-tag v-else style="background: #fff; border-style: dashed" @click="showInput">
          <plus-outlined />
          添加选项
        </a-tag>
      </div>
    </a-form-item>
    <a-form-item label="控件默认值" name="default_value">
      <a-input v-model:value="formField.default_value" />
    </a-form-item>
    <a-form-item label="排序" name="sort">
      <a-input-number v-model:value="formField.sort" />
    </a-form-item>
    <a-form-item label="规则" name="rules">
      <HuluRules v-model="formField.field_rules" />
    </a-form-item>
    <a-form-item>
      <a-button type="primary" @click="saveField">
        保存
      </a-button>
    </a-form-item>
  </a-form>
</template>

<template>
<a-form ref="formEditRef" :model="formField">
    <a-form-item
label="模板" v-if="!tid" name="template_id"
                 :rules="[rulesStore.getRule('template_id') ? rulesStore.getRule('template_id') : { required: false }]"
>
      <a-select v-model:value="formField.template_id">
        <a-select-option v-for="(t, m) in templateOpts" :value="t.id" :key="m">
          {{ t.template_name }}
        </a-select-option>
      </a-select>
    </a-form-item>
    <a-form-item
label="控件名称" name="field_name"
                 :rules="[rulesStore.getRule('field_name') ? rulesStore.getRule('field_name') : { required: false }]"
>
      <a-input v-model:value="formField.field_name" />
    </a-form-item>
    <a-form-item
label="控件类型" name="field_type"
                 :rules="[rulesStore.getRule('field_type') ? rulesStore.getRule('field_type') : { required: false }]"
>
      <a-select v-model:value="formField.field_type">
        <a-select-option value="text">
文本框
</a-select-option>
        <a-select-option value="number">
数字输入
</a-select-option>
        <a-select-option value="textarea">
文本域
</a-select-option>
        <a-select-option value="select">
下拉框
</a-select-option>
        <a-select-option value="radio">
单选框
</a-select-option>
        <a-select-option value="checkbox">
复选框
</a-select-option>
        <a-select-option value="date">
日期框
</a-select-option>
        <a-select-option value="file">
文件
</a-select-option>
      </a-select>
    </a-form-item>
    <a-form-item
label="字段名(英文)" name="field"
                 :rules="[rulesStore.getRule('field') ? rulesStore.getRule('field') : { required: false }]"
>
      <a-input v-model:value="formField.field" />
    </a-form-item>
    <a-form-item label="控件选项" name="field_value">
      <div>
        <template v-for="(tag, index) in formField.field_value" :key="index">
          <a-tooltip :title="tag">
            <a-tag :closable="true" color="processing" @close="handleClose(tag)">
              {{ tag }}
            </a-tag>
          </a-tooltip>
        </template>
        <a-input
ref="inputRef" v-model:value="state.inputValue" v-if="state.inputVisible" type="text"
                 size="small" :style="{ width: '78px' }" @keyup.enter="handleEditInputConfirm"
/>
        <a-tag v-else style="background: #fff; border-style: dashed" @click="showInput">
          <plus-outlined />
          添加选项
        </a-tag>
      </div>
</a-form-item>
    <a-form-item label="控件默认值" name="default_value">
      <a-input v-model:value="formField.default_value" />
    </a-form-item>
    <a-form-item label="排序" name="sort">
      <a-input-number v-model:value="formField.sort" />
    </a-form-item>
    <a-form-item label="规则" name="rules">
      <HuluRules v-model="formField.field_rules" />
    </a-form-item>
    <a-form-item>
      <a-button type="primary" @click="saveField">
保存
</a-button>
    </a-form-item>
  </a-form>
</template>.includes(inputValue)) {
    tags = [...tags, inputValue]
  }
  Object.assign(state, {
    tags,
    inputVisible: false,
    inputValue: '',
  })
  formField.value.field_value = tags
}

const handleClose = (removedTag: string) => {
  const tags = state.tags.filter(tag => tag !== removedTag)
  state.tags = tags
  formField.value.field_value = tags
}

// 控件
const saveField = async () => {
  try {
    // 先清空一下验证
    formEditRef.value.clearValidate()

    if (tid.value) {
      await updateTemplateForm(formField.value)
      formEditRef.value.resetFields()
      } else {
      await storeTemplateForm(formField.value)
      formEditRef.value.resetFields()
      }

  }
 catch (error) {
    formEditRef.value.validate()
  }
}
defineExpose({
  loadTemplateOpts,
})
</script>

<style></style>

<style></style>
