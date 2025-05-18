<template>
  <div>
    <n-row>
      <n-col :span="8"></n-col>
      <n-col :span="8">
        <div class="p-3">
          <n-card>
            <div class="text-xl font-bold mb-3 text-center">
              <span>{{ flow.flow_name }}</span>
              <div>
                <n-tag v-if="flow.Template" type="success" class="ml-2">
                  {{ flow.Template.template_name }}
                </n-tag>
              </div>
            </div>

            <n-form-item label="审批标题" :rules="[{ required: true, message: '输入标题' }]" style="max-width: 600px"
                         v-bind="{
                                labelCol: { span: 8 },
                                wrapperCol: { span: 16 },
                            }">
              <n-input placeholder="请输入标题" v-model:value="title"/>
            </n-form-item>
            <HuluForm :fields="fillFields" @submit="onSubmit" ref="huluFormRef"></HuluForm>
          </n-card>
        </div>
      </n-col>
      <n-col :span="8"></n-col>

    </n-row>
  </div>

</template>

<script setup lang='ts'>
import {useMessage} from "naive-ui"

const message = useMessage()
const {loadFlowEntryConfig, storeEntry} = useEntry();
const route = useRoute();
const id = route.params.id;
console.log("id", id)
const fillFields = ref([]);
const flow = ref({})
const title = ref("")
const init = async () => {
  if (id) {
    const {data} = await loadFlowEntryConfig(id);
    flow.value = data
    fillFields.value = data.Template.TemplateForms
    // console.log(fillFields.value)
  }
}

const huluFormRef = ref()
const onSubmit = async (values) => {
  if (title.value == '') {
    return message.error("请输入标题")
  }
  Object.assign(values, {flow_id: +id, title: title.value})
  try {
    huluFormRef.value.clearValidate()
    console.log(values)
    await storeEntry(values)
  } catch (error) {
    huluFormRef.value.validate()
  }
}
init();

</script>

<style></style>
