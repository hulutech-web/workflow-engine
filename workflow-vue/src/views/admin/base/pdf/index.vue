<template>
    <div>PDF管理</div>
    <a-button type="primary" @click="loadPdf">预览</a-button>
</template>

<script setup lang='ts'>
import { http } from "@/plugins/axios";
const loadPdf = async () => {
    return await http.request({
        url: `excel/pdf`,
        method: "GET",
    });
};
</script>

<style></style>