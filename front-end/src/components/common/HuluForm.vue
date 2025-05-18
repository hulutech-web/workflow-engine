<template>
  <div v-if="formFields">
    <n-form :model="formState" ref="formRef" v-bind="{
            labelCol: { span: 8 },
            wrapperCol: { span: 16 },
        }" style="max-width: 600px">
      <div v-for="(field, index) in formFields" :key="index" v-if="formFields.length > 0">
        <n-form-item :label="field.field_name" v-if="field['field_type'] == 'text'" :name="field.field"
                     :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
          <n-input v-model:value="formState[field.field]" style="width:100%"/>
        </n-form-item>
        <n-form-item :label="field.field_name" v-if="field['field_type'] == 'select'" :name="field.field"
                     :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
          {{formState[field.field_value]}}
          <n-select v-model:value="formState[field.field]" style="width:100%;" :options="field.field_value.map(item=>{
                      return {
                        label:item,
                        value:item
                      }
                    })">
          </n-select>
        </n-form-item>
        <n-form-item :label="field.field_name" v-if="field['field_type'] == 'textarea'" :name="field.field"
                     :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
          <n-input type="textarea" v-model:value="formState[field.field]" style="width:100%"/>
        </n-form-item>
        <n-form-item :label="field.field_name" v-if="field['field_type'] == 'number'" :name="field.field"
                     :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
          <n-input-number v-model:value="formState[field.field]" :value="formState[field.field]?formState[field.field]:0" :min="1" style="width:100%"/>
        </n-form-item>
        <n-form-item :label="field.field_name" v-if="field['field_type'] == 'date'" :name="field.field"
                     :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
          <n-date-picker clearable format="YYYY-MM-DD" v-model:value="formState[field.field]" style="width:100%"/>
        </n-form-item>
        <n-form-item :label="field.field_name" v-if="field['field_type'] == 'radio'" :name="field.field"
                     :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
          <n-radio-group v-model:value="formState[field.field]" style="width:100%">
            <n-radio :value="item" v-for="(item, ind) in field['field_value']" :key="ind">
              {{ item }}
            </n-radio>
          </n-radio-group>
        </n-form-item>
        <n-form-item :label="field.field_name" v-if="field['field_type'] == 'checkbox'" :name="field.field"
                     :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
          <n-checkbox-group v-model:value="formState[field.field]" style="width:100%">
            <n-space item-style="display: flex;">
              <n-checkbox :value="item" v-for="(item, ind) in field['field_value']" :key="ind">
                {{ item }}
              </n-checkbox>
            </n-space>
          </n-checkbox-group>

        </n-form-item>

        <n-form-item :label="field.field_name" :key="field.name" v-if="field['field_type'] == 'file'"
                     :name="field.field"
                     :rules="[rulesStore.getRule(field.field) ? rulesStore.getRule(field.field) : { required: false }]">
          <!-- accept=".xls, .xlsx, .csv" -->
          <div @click="tapHandle(field)">
            <n-upload
              multiple
              directory-dnd
              action="https://www.mocky.io/v2/5e4bafc63100007100d8b70f"
              :max="5"
            >
              <n-upload-dragger>
                <div style="margin-bottom: 12px">
                  <n-icon size="48" :depth="3">
                  </n-icon>
                </div>
                <n-text style="font-size: 16px">
                  点击或者拖动文件到该区域来上传
                </n-text>
                <n-p depth="3" style="margin: 8px 0 0 0">
                  请不要上传敏感数据，比如你的银行卡号和密码，信用卡号有效期和安全码
                </n-p>
              </n-upload-dragger>
            </n-upload>
          </div>
        </n-form-item>

      </div>
      <slot name="default">
        <n-form-item :wrapper-col="{
                    wrapperCol: { span: 16 },
                    offset: 8
                }">
          <n-button type="primary" @click="submit">提交</n-button>
        </n-form-item>
      </slot>
    </n-form>
  </div>
</template>

<script setup lang='ts'>
import useRulesStore from '@/store/useRulesStore.ts'
import {useMessage} from "naive-ui";

const message = useMessage();

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
const formState = ref({})

const genFileList = ref([])
const initFormState = () => {
  const state = {}
  formFields.value.forEach(field => {
    switch (field.field_type) {
      case 'number':
        state[field.field] = field.field_default_value==''?0:field.field_default_value
        break;
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
  console.log(formState.value)

}
watch(() => props.fields, (newVal) => {
  formFields.value = newVal
  console.log(newVal)
  initFormState()
})

watch(() => props.entryDatas, (newVal) => {
  if (newVal && newVal.length > 0) {
    // 将数据进行填充
    newVal.forEach(item => {
      console.log(item['field_name'])
      formState.value[item['field_name']] = item.field_value
    })
  } else {
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
