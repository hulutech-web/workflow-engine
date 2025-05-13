<template>
  <div class="p-3">
    <n-card>
      <n-form ref="formRef" :model="flowState" :label-col="{ style: { width: '150px' } }"
              :wrapper-col="{ span: 6 }">
        <n-form-item label="流程编码" name="flow_no"
                     :rules="[rulesStore.getRule('flow_no') ? rulesStore.getRule('flow_no') : { required: false }]">
          <n-input v-model:value="flowState.flow_no"></n-input>
        </n-form-item>
        <n-form-item label="流程名称" name="flow_name"
                     :rules="[rulesStore.getRule('flow_name') ? rulesStore.getRule('flow_name') : { required: false }]">
          <n-input v-model:value="flowState.flow_name"></n-input>
        </n-form-item>
        <n-form-item label="模板" name="template_id"
                     :rules="[rulesStore.getRule('template_id') ? rulesStore.getRule('template_id') : { required: false }]">
          <n-select v-model:value="flowState.template_id" :options="options.templates" placeholder="请选择模板">

          </n-select>

        </n-form-item>
        <n-form-item label="流程类型" name="type_id"
                     :rules="[rulesStore.getRule('type_id') ? rulesStore.getRule('type_id') : { required: false }]">
          <n-select v-model:value="flowState.type_id" :options="options.flowtypes" placeholder="请选择类型">
          </n-select>
        </n-form-item>
        <n-form-item>
          <n-button type="primary" @click="onSubmit">提交</n-button>
        </n-form-item>
      </n-form>
    </n-card>
  </div>
</template>

<script setup lang='ts'>
const route = useRoute()
const id = ref(null)
id.value = route.params.id
const {createFlow, storeFlow, updateFlow, showFlow} = useFlow();
import useRulesStore from '@/store/useRulesStore.ts'

const rulesStore = useRulesStore();
const flowState = ref({
  id: null,
  flow_no: "",
  flow_name: "",
  template_id: null,
  type_id: null,
})
const options = ref({
  templates: [],
  flowtypes: []
})
const init = async () => {
  const {data} = await createFlow();
  options.value.templates = data.templates.map(item => {
    return {
      label: item.template_name,
      value: item.id
    }
  });
  options.value.flowtypes = data.flowtypes.map(item => {
    return {
      label: item.type_name,
      value: item.id
    }
  });
}
const formRef = ref()
onMounted(async () => {
  init();
  if (id.value != null) {
    const {data} = await showFlow(id.value)
    flowState.value = data
  }
})

const onSubmit = async (e) => {
  e.preventDefault()
  formRef.value?.validate(async (errors) => {
    if (!errors) {
      if (flowState.value.id) {
        await updateFlow(flowState.value)
      }else{
        await storeFlow(flowState.value)
      }
      formRef.value?.restoreValidation()
    } else {
      console.log('errors', errors)
    }
  })
}

</script>

<style></style>
