<template>
    <div>
        <a-card title="批量任务">
            <a-row :gutter="[32, 8]">
                <a-col :span="8">
                    <p>
                        清单
                    </p>
                    <a-table :columns="columns" :data-source="data">
                        <template #headerCell="{ column }">
                            <template v-if="column.key === 'name'">
                                <span>
                                    <smile-outlined />
                                    Name
                                </span>
                            </template>
                        </template>

                        <template #bodyCell="{ column, record }">
                            <template v-if="column.key === 'name'">
                                <a>
                                    {{ record.name }}
                                </a>
                            </template>
                            <template v-else-if="column.key === 'phone'">
                                <span>
                                    {{ record.phone }}
                                </span>
                            </template>
                            <template v-else-if="column.key === 'state'">
                                <a-badge status="success" v-if="record.state == 1" />
                                <a-badge status="processing" v-else />
                            </template>

                            <template v-else-if="column.key === 'action'">
                                <a v-if="record.state == 1">待命</a>
                                <a v-else>开始</a>
                            </template>
                        </template>
                    </a-table>
                </a-col>
                <a-col :span="16">
                    <p>操作</p>
                    <a-steps v-model:current="current" :items="[
                        {
                            title: '情况一',
                            description: '基础信息',
                        },
                        {
                            title: '情况二',
                            description: '资费感知',
                        },
                        {
                            title: '情况三',
                            description: '服务感知',
                        },
                        {
                            title: '情况四',
                            description: '售后感知',
                        },
                    ]">
                        <template #progressDot="{ index, status, prefixCls }">
                            <a-popover>
                                <template #content>
                                    <span>step {{ index }} status: {{ status }}</span>
                                </template>
                                <span :class="`${prefixCls}-icon-dot`" />
                            </a-popover>
                        </template>
                    </a-steps>

                    <a-card title="信息采集">
                        <a-form :model="formState" name="validate_other" @finishFailed="onFinishFailed"
                            @finish="onFinish">
                            <a-form-item name="select" label="年龄结构" has-feedback
                                :rules="[{ required: true, message: 'Please select your country!' }]">
                                <a-select v-model:value="formState.select" placeholder="Please select a country">
                                    <a-select-option value="china">中年</a-select-option>
                                    <a-select-option value="usa">青年</a-select-option>
                                    <a-select-option value="usa">老年</a-select-option>
                                </a-select>
                            </a-form-item>

                            <a-form-item name="select-multiple" label="套餐组成"
                                :rules="[{ required: true, message: 'Please select your favourite colors!', type: 'array' }]">
                                <a-select v-model:value="formState['select-multiple']" mode="multiple"
                                    placeholder="Please select favourite colors">
                                    <a-select-option value="red">单宽</a-select-option>
                                    <a-select-option value="green">单卡</a-select-option>
                                    <a-select-option value="blue">融合</a-select-option>
                                </a-select>
                            </a-form-item>

                            <a-form-item label="在网时间">
                                <a-form-item name="input-number" no-style>
                                    <a-input-number v-model:value="formState['input-number']" :min="1" :max="10" addon-after="年" />
                                </a-form-item>
                            </a-form-item>

                            <a-form-item name="switch" label="总体感知">
                                <a-switch v-model:checked="formState.switch" checked-children="好" un-checked-children="不好" />
                            </a-form-item>

                            <a-form-item name="slider" label="服务等级">
                                <a-slider v-model:value="formState.slider" :marks="{
                                    0: 'A',
                                    20: 'B',
                                    40: 'C',
                                    60: 'D',
                                    80: 'E',
                                    100: 'F',
                                }" />
                            </a-form-item>

                            <a-form-item name="radio-group" label="满意度">
                                <a-radio-group v-model:value="formState['radio-group']">
                                    <a-radio value="a">一般</a-radio>
                                    <a-radio value="b">良好</a-radio>
                                    <a-radio value="c">满意</a-radio>
                                </a-radio-group>
                            </a-form-item>

                            
                            

                            <a-form-item name="rate" label="评分">
                                <a-rate v-model:value="formState.rate" allow-half />
                            </a-form-item>

                            <a-form-item name="upload" label="Upload" extra="其他材料">
                                <a-upload v-model:fileList="formState.upload" name="logo" action="/upload.do"
                                    list-type="picture">
                                    <a-button>
                                        <template #icon>
                                            <UploadOutlined />
                                        </template>
                                        Click to upload
                                    </a-button>
                                </a-upload>
                            </a-form-item>

                            <a-form-item label="Dragger">
                                <a-form-item name="dragger" no-style>
                                    <a-upload-dragger v-model:fileList="formState.dragger" name="files"
                                        action="/upload.do">
                                        <p class="ant-upload-drag-icon">
                                            <InboxOutlined />
                                        </p>
                                        <p class="ant-upload-text">Click or drag file to this area to upload</p>
                                        <p class="ant-upload-hint">Support for a single or bulk upload.</p>
                                    </a-upload-dragger>
                                </a-form-item>
                            </a-form-item>

                            <a-form-item>
                                <a-space>

                                    <a-button type="primary" html-type="submit">保存</a-button>
                                    <a-button type="primary" html-type="submit">下一步</a-button>
                                </a-space>
                            </a-form-item>
                        </a-form>
                    </a-card>

                </a-col>
            </a-row>
        </a-card>
    </div>
</template>

<script setup lang='ts'>
import { SmileOutlined, DownOutlined } from '@ant-design/icons-vue';
const columns = [
    {
        name: '姓名',
        dataIndex: 'name',
        key: 'name',
    },
    {
        title: '年龄',
        dataIndex: 'age',
        key: 'age',
    },
    {
        title: '地址',
        dataIndex: 'address',
        key: 'address',
    },
    {
        title: '电话',
        key: 'phone',
        dataIndex: 'phone',
    },
    {
        title: '状态',
        key: 'state',
        dataIndex: 'state',
    },
    {
        title: '操作',
        key: 'action',
    },
];

const data = [
    {
        key: '1',
        name: 'John Brown',
        age: 32,
        state: 1,
        address: 'New York No. 1 Lake Park',
        phone: '12345xxxx',
    },
    {
        key: '2',
        name: 'Jim Green',
        age: 42,
        state: 1,

        address: 'London No. 1 Lake Park',
        phone: '12345xxxx',

    },
    {
        key: '3',
        name: 'Joe Black',
        age: 32,
        state: 0,
        address: 'Sidney No. 1 Lake Park',
        phone: '123456789',
    },
];

const current = ref<number>(0);
import { reactive } from 'vue';
import { UploadOutlined, InboxOutlined } from '@ant-design/icons-vue';

const formItemLayout = {
    labelCol: { span: 6 },
    wrapperCol: { span: 14 },
};

const formState = reactive<Record<string, any>>({
    'input-number': 3,
    'checkbox-group': ['A', 'B'],
    rate: 3.5,
});
const onFinish = (values: any) => {
    console.log('Success:', values);
};

const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
};
</script>

<style></style>