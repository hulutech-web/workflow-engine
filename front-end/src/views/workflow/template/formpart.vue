<template>

  <n-form :model="formField" ref="formEditRef">
    <n-form-item label="模板" name="template_id" v-if="!tid"
                 :rules="[rulesStore.getRule('template_id') ? rulesStore.getRule('template_id') : { required: false }]">
      <n-select v-model:value="formField.template_id" :options="options.t_option" placeholder="请选择模板">

      </n-select>
    </n-form-item>
    <n-form-item label="控件名称" name="field_name"
                 :rules="[rulesStore.getRule('field_name') ? rulesStore.getRule('field_name') : { required: false }]">
      <n-input v-model:value="formField.field_name"></n-input>
    </n-form-item>
    <n-form-item label="控件类型" name="field_type"
                 :rules="[rulesStore.getRule('field_type') ? rulesStore.getRule('field_type') : { required: false }]">
      <n-select v-model:value="formField.field_type" :options="options.t_option">

      </n-select>
    </n-form-item>
    <n-form-item label="字段名(英文)" name="field"
                 :rules="[rulesStore.getRule('field') ? rulesStore.getRule('field') : { required: false }]">
      <n-input v-model:value="formField.field"></n-input>
    </n-form-item>
    <n-form-item label="控件选项" name="field_value">
        <n-dynamic-tags :closable="true" type="success" v-model:value="formField.field_value">
        </n-dynamic-tags>
    </n-form-item>
    <n-form-item label="控件默认值" name="default_value">
      <n-input v-model:value="formField.default_value"></n-input>
    </n-form-item>
    <n-form-item label="排序" name="sort">
      <n-input-number v-model:value="formField.sort"></n-input-number>
    </n-form-item>
    <n-form-item label="规则" name="rules">
      <HuluRules v-model="formField.field_rules" />
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="saveField">保存</n-button>
    </n-form-item>
  </n-form>
</template>

<script setup lang='ts'>

import useRulesStore from '@/store/useRulesStore.ts'
import { watch } from 'vue';
const { storeTemplateForm, updateTemplateForm, showTemplateForm } = useTemplateForm();
const { loadTemplates } = useTemplate();
const rulesStore = useRulesStore();
const tid = ref()
const formEditRef = ref()
const inputRef = ref();
const props = defineProps({
  id: {
    type: Number,
    default: null
  }
})

const options = ref({
  f_option:[
    {
      label: '文本框',
      value: 'text'
    },
    {
      label: '数字输入',
      value: 'number'
    },
    {
      label: '文本域',
      value: 'textarea'
    },
    {
      label: '下拉框',
      value: 'select'
    },
    {
      label: '单选框',
      value: 'radio'
    },
    {
      label: '复选框',
      value: 'checkbox'
    },
    {
      label: '日期框',
      value: 'date'
    },
    {
      label: '文件',
      value: 'file'
    }


  ],
  t_option:[

  ]
})
const formField = ref({
  id: null,
  field: "",
  field_default_value: "",
  field_name: "",
  field_type: "",
  field_value: [],
  sort: 0,
  template_id: "",
  field_rules: []
})
const loadTemp = async () => {
  const { data } = await showTemplateForm(+tid.value)
  formField.value = data
}
const loadTemplateOpts = async () => {
  const { data } = await loadTemplates()
  templateOpts.value = data.data
  options.value.t_option = data.data.map(item => {
    return {
      label: item.template_name,
      value: item.id
    }
  })
}

watch(() => props.id, async (val) => {

  if (val) {
    tid.value = val
    await loadTemp()
  } else {
    await loadTemplateOpts()
  }
}, {
  deep: true,
  immediate: true
})

const templateOpts = ref([])


const state = reactive({
  tags: [],
  inputVisible: false,
  inputValue: '',
});


const showInput = () => {
  state.inputVisible = true;
  nextTick(() => {
    inputRef.value.focus();
  });
};

const handleEditInputConfirm = () => {
  const inputValue = state.inputValue;
  let tags = state.tags;
  if (inputValue && tags.indexOf(inputValue) === -1) {
    tags = [...tags, inputValue];
  }
  Object.assign(state, {
    tags,
    inputVisible: false,
    inputValue: '',
  });
  formField.value.field_value = tags
};




const handleClose = (removedTag: string) => {
  const tags = state.tags.filter(tag => tag !== removedTag);
  state.tags = tags;
  formField.value.field_value = tags
};



// 控件
const saveField = async () => {
  console.log(formField.value)
  try {
    //先清空一下验证
    formEditRef.value.clearValidate()

    if (tid.value) {
      await updateTemplateForm(formField.value)
      formEditRef.value.resetFields();
    } else {
      await storeTemplateForm(formField.value)
      formEditRef.value.resetFields();
    }

  } catch (error) {
    formEditRef.value.validate()
  }
}
defineExpose({
  loadTemplateOpts
})
</script>

<style></style>
