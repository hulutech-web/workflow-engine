<template>
    <div class="p-3">
        <a-card>
            <a-form ref="formRef" :model="flowState" :label-col="{ style: { width: '150px' } }"
                :wrapper-col="{ span: 6 }">
                <a-form-item label="流程编码" name="flow_no"
                    :rules="[rulesStore.getRule('flow_no') ? rulesStore.getRule('flow_no') : { required: false }]">
                    <a-input v-model:value="flowState.flow_no"></a-input>
                </a-form-item>
                <a-form-item label="流程名称" name="flow_name"
                    :rules="[rulesStore.getRule('flow_name') ? rulesStore.getRule('flow_name') : { required: false }]">
                    <a-input v-model:value="flowState.flow_name"></a-input>
                </a-form-item>
                <a-form-item label="模板" name="template_id" 
                    :rules="[rulesStore.getRule('template_id') ? rulesStore.getRule('template_id') : { required: false }]">
                    <a-select v-model:value="flowState.template_id">
                        <a-select-option v-for="(opt, ind) in options.templates" :key="ind" :value="opt.id">
                            {{ opt.template_name }}
                        </a-select-option>
                    </a-select>
                </a-form-item>
                <a-form-item label="流程类型" name="type_id" 
                    :rules="[rulesStore.getRule('type_id') ? rulesStore.getRule('type_id') : { required: false }]">
                    <a-select v-model:value="flowState.type_id">
                        <a-select-option v-for="(opt, ind) in options.flowtypes" :key="ind" :value="opt.id">
                            {{ opt.type_name }}
                        </a-select-option>
                    </a-select>
                </a-form-item>
                <a-form-item>
                    <a-button type="primary" @click="onSubmit">提交</a-button>
                </a-form-item>
            </a-form>
        </a-card>
    </div>
</template>

<script setup lang='ts'>
const route = useRoute()
const id = ref(null)
id.value = route.params.id
const { createFlow, storeFlow, updateFlow, showFlow } = useFlow();
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
    const data = await createFlow();
    console.log(data);

    options.value.templates = data.templates;
    options.value.flowtypes = data.flowtypes;
}
const formRef = ref()
onMounted(async () => {
    init();
    if (id.value != null) {
        const { data } = await showFlow(id.value)
        //将data变成响应式数据并赋值给employeeForm
        flowState.value = data
    }
})
const onSubmit = async () => {
    try {
        //先清空一下验证
        formRef.value.clearValidate()
        if (id.value != null) {
            await updateFlow(flowState.value)

        } else {
            const { data } = await storeFlow(flowState.value)
        }
    } catch (error) {
        formRef.value.validate()
    }
}

</script>

<style></style>