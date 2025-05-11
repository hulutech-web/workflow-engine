<template>
    <div>
        <a-row :gutter="[24, 1]">
            <a-col :span="4">
                <a-card>
                    <p>选择运营商</p>
                </a-card>
            </a-col>
            <a-col :span="20">
                <a-card>
                    <p>设备位置</p>
                </a-card>
                <div id="mapContainer"></div>
            </a-col>
        </a-row>
    </div>
</template>

<script setup lang="ts">
import { loadDataAndRender } from './mapInit';
import dayjs from 'dayjs'
const { indexDevice } = useDevice()
// 3天前
const startTime = ref(dayjs().subtract(3, 'day').format('YYYY-MM-DD HH:mm:ss'))
// 当前时间
const endTime = ref(dayjs().format('YYYY-MM-DD HH:mm:ss'))


onMounted(async () => {
    await loadDataAndRender(startTime.value, endTime.value);
});
</script>

<style scoped></style>