<template>
    <div v-if="formFields">
        <a-form :model="formState" ref="formRef" v-bind="{
            labelCol: { span: 8 },
            wrapperCol: { span: 16 },
        }" style="max-width: 600px">
            <div v-for="(field, index) in formFields" :key="index" v-if="formFields.length > 0">
                <a-form-item :label="field.field_name" v-if="field['field_type'] == 'text'" :name="field.field"
                    :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
                    <a-input v-model:value="formState[field.field]" style="width:100%" />
                </a-form-item>
                <a-form-item :label="field.field_name" v-if="field['field_type'] == 'select'" :name="field.field"
                    :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
                    <a-select v-model:value="formState[field.field]" style="width:100%;">
                        <a-select-option :value="item" v-for="(item, ind) in field['field_value']" :key="ind">
                            {{ item }}
                        </a-select-option>
                    </a-select>
                </a-form-item>
                <a-form-item :label="field.field_name" v-if="field['field_type'] == 'textarea'" :name="field.field"
                    :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
                    <a-textarea v-model:value="formState[field.field]" style="width:100%" />
                </a-form-item>
                <a-form-item :label="field.field_name" v-if="field['field_type'] == 'number'" :name="field.field"
                    :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
                    <a-input-number v-model:value="formState[field.field]" :min="1" style="width:100%" />
                </a-form-item>
                <a-form-item :label="field.field_name" v-if="field['field_type'] == 'date'" :name="field.field"
                    :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
                    <a-date-picker valueFormat="YYYY-MM-DD" v-model:value="formState[field.field]" style="width:100%" />
                </a-form-item>
                <a-form-item :label="field.field_name" v-if="field['field_type'] == 'radio'" :name="field.field"
                    :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
                    <a-radio-group v-model:value="formState[field.field]" style="width:100%">
                        <a-radio :value="item" v-for="(item, ind) in field['field_value']" :key="ind">
                            {{ item }}
                        </a-radio>
                    </a-radio-group>
                </a-form-item>
                <a-form-item :label="field.field_name" v-if="field['field_type'] == 'checkbox'" :name="field.field"
                    :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
                    <a-checkbox-group v-model:value="formState[field.field]" style="width:100%">
                        <a-checkbox :value="item" v-for="(item, ind) in field['field_value']" :key="ind">
                            {{ item }}
                        </a-checkbox>
                    </a-checkbox-group>

                </a-form-item>

                <a-form-item :label="field.field_name" :key="field.name" v-if="field['field_type'] == 'file'"
                    :name="field.field"
                    :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
                    <!-- accept=".xls, .xlsx, .csv" -->
                    <div @click="tapHandle(field)">
                        <a-upload-dragger v-model:fileList="genFileList" name="file" :multiple="false"
                            action="/api/upload"
                            :headers="{ 'Authorization': `Bearer ${storage.get(CacheEnum.TOKEN_NAME)}` }"
                            @change="handleChange" @drop="handleDrop">
                            <p class="ant-upload-drag-icon">
                                <inbox-outlined></inbox-outlined>
                            </p>
                            <p class="ant-upload-text">点击上传或拖拽上传</p>
                            <p class="ant-upload-hint">
                                仅支持单文件上传，支持格式为excel
                            </p>
                        </a-upload-dragger>
                    </div>
                </a-form-item>

            </div>
            <slot name="default">
                <a-form-item :wrapper-col="{
                    wrapperCol: { span: 16 },
                    offset: 8
                }">
                    <a-button type="primary" @click="submit">提交</a-button>
                </a-form-item>
            </slot>
        </a-form>
    </div>
</template>

<script setup lang='ts'>
import useRulesStore from '@/store/useRulesStore.ts'
import { message } from 'ant-design-vue'

const rulesStore = useRulesStore();
const storage = useStorage()
const props = defineProps({
    fields: {
        type: Array,
        default: []
    },
    entryDatas: {
        type: Array,
        default: []
    }
})

const formFields = ref([])
const formState = ref({

})

const genFileList = ref([])
const initFormState = () => {
    const state = {}
    formFields.value.forEach(field => {
        switch (field.field_type) {
            case 'select':
                state[field.field] = []
                break;
            case 'radio':
                state[field.field] = []
                break;
            case 'checkbox':
                state[field.field] = []
                break;
            case 'slice':
                state[field.field] = []
                break;
            default:
                state[field['field']] = field.field_default_value ?? ""
                break;
        }
    })
    formState.value = state
    
}
watch(() => props.fields, (newVal) => {
    formFields.value = newVal
    initFormState()
})

watch(() => props.entryDatas, (newVal) => {
    if (newVal && newVal.length > 0) {
        // 将数据进行填充
        newVal.forEach(item => {
            console.log(item['field_name'])
            formState.value[item['field_name']] = item.field_value
        })
    }else{
        console.log("is not ok")
    }
})

const emit = defineEmits(["submit"])
const submit = () => {
    emit("submit", formState.value)
}
const currentFile = ref();
const currentUploadTap = ref("")

const tapHandle = (e) => {
    currentUploadTap.value = e.field
}

const handleChange = (info) => {
    const status = info.file.status;
    if (status !== 'uploading') {
        // console.log(info.file, info.fileList);
    }
    if (status === 'done') {
        message.success(`${info.file.name} 上传成功.`);
        console.log(info, 333333)
        let urlArr = info.fileList.map(item => {
            if (isRef(item)) {
                return item.value.response.data
            } else {
                return item.response.data
            }
        })
        formState.value[currentUploadTap.value] = urlArr
        console.log(2222, formState.value)
    } else if (status === 'error') {
        message.error(`${info.file.name} 上传失败.`);
    }
    if (info.file.status === 'done') {
        // 替换当前上传文件列表为最新的上传成功的文件
        currentFile.value = [info.file];
    } else if (info.file.status === 'removed') {
        // 如果文件被移除，则清空当前文件
        currentFile.value = null;
    }
};
function handleDrop(e: DragEvent) {
    console.log(e);
    // 当文件状态变为已上传成功时
}
const formRef = ref()

const validate = () => {
    formRef.value.validate()
}
const clearValidate = () => {
    formRef.value.clearValidate()
}

defineExpose({
    validate,
    clearValidate
})

</script>

<style></style>