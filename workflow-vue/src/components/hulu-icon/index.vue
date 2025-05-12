<template>
    <div>
        <HuluIcon :type="iconType" :style="{ fontSize: props.fontSize, color: props.color }">
        </HuluIcon>
    </div>
</template>

<script setup lang="ts">
import { createFromIconfontCN } from '@ant-design/icons-vue';
const props = defineProps({
    name: {
        type: String,
        default: 'icon-stamp',
    },
    fontSize: {
        type: String,
        default: '16px',
    },
    color: {
        type: String,
        default: '#fff',
    },
})
// 直接在 setup 里创建组件，无需再 export default
const HuluIcon = createFromIconfontCN({
    scriptUrl: '//at.alicdn.com/t/c/font_4600562_39p1322t60g.js', // 从 IconFont 获取的实际 URL
});

// 使用 defineProps 和 computed 而不是 props 和 setup 返回函数
const iconType = computed(() => {
    // 如果传入了 name 属性，则使用传入的值，否则使用默认图标
    return props.name || 'icon-stamp';
});
</script>