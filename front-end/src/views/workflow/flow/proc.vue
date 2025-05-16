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
                                    <div>
                                        <p class="text-xl">批复内容：</p>
                                        <a-textarea v-model:value="content" placeholder="请填写批复内容" :rows="4" />
                                        <p class="mt-3">
                                            <a-space>
                                                <a-button type="primary" @click="pass">
                                                    同意
                                                </a-button>
                                                <a-button type="primary" @click="unpass">
                                                    驳回
                                                </a-button>
                                            </a-space>
                                        </p>
                                    </div>
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
const { setPass, setUnPass } = useProc();
const route = useRoute();

const flow_id = route.params.flow_id
const entry_id = route.params.entry_id;
// 获取路由query参数
const query = route.query;
const process_id = query.process_id
const proc_id = query.proc_id
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

const content = ref("")
const huluFormRef = ref()
const onSubmit = async (values) => {

}

const pass = async (values) => {
    await setPass({
        content: content.value,
        process_id: process_id
    })
}
const unpass = async (values) => {
    await setUnPass({
        content: content.value,
        proc_id: proc_id
    })
}
init();

</script>

<style></style>