<template>
    <div class="container">
        <div>
            <div class="content-box">
                <vue-cropper :autoCrop="true" ref="cropper" img="/charging.png">
                </vue-cropper>
            </div>
            <a-space>
                <a-button :loading="loading" type="primary" @click="click">获取截图</a-button>
                <a-button type="primary" danger ghost @click="updateCut">保存截图</a-button>
            </a-space>
        </div>
        <div class="preview-box">
            <p class="pre-title">预览</p>
            <div class="pre-box">
                <img :src="cutImg" alt="" mode="widthFix">
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { createVNode, h } from 'vue'
import { Modal } from 'ant-design-vue';
import { ExclamationCircleOutlined } from '@ant-design/icons-vue';
const props = defineProps({
    modelValue: {
        type: String,
        default: "",
    },
})

const emit = defineEmits(["update:modelValue"])


const cutImg = ref("")
const cmpImg = ref("")
// 监听props重点的modelValue
watch(
    () => props.modelValue,
    (newProps) => {
        if (newProps !== "") {
            cmpImg.value = newProps
        }
    }, {
    immediate: true
}
);


const cropper = ref()
const loading = ref(false)
const click = () => {
    loading.value = true;
    cropper.value.getCropData((data) => {
        open(data)
    })
}
const updateCut = ()=>{
    
}
const open = (img) => {
    Modal.confirm({
        title: '确定裁切吗？',
        // 居中
        centered: true,
        // 设置宽高
        width: 400,
        // 设置遮罩层
        maskClosable: false,
        icon: createVNode(ExclamationCircleOutlined),
        content: h(
            'img', { style: 'width:100%', src: img, mode: 'aspectFill' }
        ),
        okText: '确认',
        cancelText: '取消',
        onOk() {
            cutImg.value = img
            loading.value = false
            emit("update:modelValue", cutImg.value)
        },
        onCancel() {
            loading.value = false
        }
    });
}

</script>

<style lang="scss" scoped>
.container {
    padding: 10px;
    box-sizing: border-box;
    width: 600px;
    height: auto;
    display: flex;
}

.content-box {
    height: 300px;
    width: 300px;
    box-sizing: border-box;
    padding: 10px;
    border: 1px solid #ededed;
    margin: 5px;
    box-shadow: 1px 1px 1px #ededed;
}

.preview-box {
    width: 300px;
    height: 300px;
    background-color: #0461A6;
    padding: 10px;
    box-sizing: border-box;
    margin: 5px;
    box-shadow: 1px 1px 1px #ededed;

    .pre-title {
        color: #fff;
        font-size: 14px;
    }

    .pre-box {
        text-align: center;
    }
}
</style>