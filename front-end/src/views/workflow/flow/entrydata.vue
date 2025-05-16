<template>
    <div>
        <a-row>
            <a-col :span="8"></a-col>
            <a-col :span="8">
                <div class="p-3">
                    <a-watermark content="已发起">
                        <div>
                            <a-card>
                                <p class="text-2xl font-bold mb-3 text-center">
                                    <span>{{ flow.flow_name }}</span>
                                <div>
                                    <a-tag v-if="flow.Template" color="blue" class="ml-2">
                                        {{ flow.Template.template_name }}
                                    </a-tag>
                                </div>
                                </p>

                                <a-form-item label="标题" :rules="[{ required: true, message: '输入标题' }]"
                                    style="max-width: 600px" v-bind="{
                                        labelCol: { span: 8 },
                                        wrapperCol: { span: 16 },
                                    }">
                                    <a-input placeholder="请输入标题" v-model:value="title" />
                                </a-form-item>
                                <Form :fields="fillFields" @submit="onSubmit" :entryDatas="entryDatas"
                                    ref="huluFormRef">
                                    <div></div>
                                </Form>
                            </a-card>
                        </div>
                    </a-watermark>

                </div>
            </a-col>
            <a-col :span="8"></a-col>

        </a-row>
    </div>

</template>

<script setup lang='ts'>
import { message } from 'ant-design-vue';

const { loadFlowEntryConfig, storeEntry, getEntryData } = useEntry();
const route = useRoute();
const flow_id = route.params.flow_id
const entry_id = route.params.entry_id;
const fillFields = ref([]);
const flow = ref({})
const title = ref("")
const entry = ref({})
const init = async () => {
    if (flow_id) {
        const { data } = await loadFlowEntryConfig(flow_id);
        flow.value = data
        fillFields.value = data.Template.TemplateForms
    }
    if (entry_id) {
        await loadEntryDatas()
    }
}
const entryDatas = ref([])
const loadEntryDatas = async () => {
    const { data } = await getEntryData(entry_id)
   
    entryDatas.value = data.entrydata
    entry.value = data.entry
    title.value = data.entry.title
}

const fillDefaultValue = () => {

}
const huluFormRef = ref()
const onSubmit = async (values) => {
    if (title.value == '') {
        return message.error("请输入标题")
    }
    Object.assign(values, { flow_id: +id, title: title.value })
    try {
        huluFormRef.value.clearValidate()
        await storeEntry(values)
    } catch (error) {
        huluFormRef.value.validate()
    }
}
init();

</script>

<style></style>