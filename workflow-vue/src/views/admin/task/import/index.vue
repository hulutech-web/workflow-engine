<template>
    <a-card title="导入数据">
        <a-row :gutter="[32, 8]">
            <a-col :span="8">
                <a-upload-dragger v-model:fileList="fileList" name="file" :multiple="true"
                    action="https://www.mocky.io/v2/5cc8019d300000980a055e76" @change="handleChange" @drop="handleDrop">
                    <p class="ant-upload-drag-icon">
                        <inbox-outlined></inbox-outlined>
                    </p>
                    <p class="ant-upload-text">拖拽文件</p>
                    <p class="ant-upload-hint">
                        仅支持EXCEL、CSV格式
                    </p>
                </a-upload-dragger>
            </a-col>
            <a-col :span="16">
                <a-card>
                    <vxe-grid ref='xGrid' v-bind="gridOptions" v-on="gridEvent">
                        <template #action="{ row }">
                            <div>
                                <a-button type="primary">删除</a-button>
                                <a-button type="primary">编辑</a-button>

                            </div>
                        </template>
                        <template #dept="{ row }">
                            <div>
                                {{ row.Dept.id == 0 ? "未分配" : row.Dept.dept_name }}
                            </div>
                        </template>
                    </vxe-grid>
                </a-card>

            </a-col>
        </a-row>
    </a-card>
</template>

<script setup lang='ts'>
import { InboxOutlined } from '@ant-design/icons-vue';
import { message } from 'ant-design-vue';
import type { UploadChangeParam } from 'ant-design-vue';
const fileList = ref([]);
const handleChange = (info: UploadChangeParam) => {
    const status = info.file.status;
    if (status !== 'uploading') {
        console.log(info.file, info.fileList);
    }
    if (status === 'done') {
        message.success(`${info.file.name} file uploaded successfully.`);
    } else if (status === 'error') {
        message.error(`${info.file.name} file upload failed.`);
    }
};
function handleDrop(e: DragEvent) {
    console.log(e);
}
const { gridOptions } = useUser()
const xGrid = ref()
const gridEvent: VxeGridListeners<RowVO> = {
    proxyQuery() {
        console.log('数据代理查询事件')
        const grid = xGrid.value
        // 获取表格中的数据
        const data = grid.getTableData().fullData
    },
    proxyDelete() {
        console.log('数据代理删除事件')
    },
    proxySave() {
        console.log('数据代理保存事件')
    }
}
</script>

<style></style>