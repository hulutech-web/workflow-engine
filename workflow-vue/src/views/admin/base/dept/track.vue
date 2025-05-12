<template>
    <div>
        <a-row :gutter="[24, 1]">
            <a-col :span="4">
                <a-card>
                    <p>选择时间</p>

                </a-card>
            </a-col>
            <a-col :span="20">
                <a-card>
                    <p>设备轨迹</p>
                    <p>总距离：{{ all_distance }}KM</p>
                </a-card>
                <div id="mapContainer"></div>
            </a-col>
        </a-row>
    </div>
</template>

<script setup lang="ts">
import { message } from 'ant-design-vue';
import { loadDataAndRender } from './trackInit';
const route = useRoute();
const id = ref(route.params.id)
console.log(id.value)
const historytracks = ref([])
const startTime = ref("2024-05-15 00:00:00")
const endTime = ref("2024-05-17 23:59:59")
const all_distance=ref(0)
onMounted(async () => {
    const {len,distance} = await loadDataAndRender(id.value, startTime.value, endTime.value);
    if (len == 0) {
        message.error("暂无数据")
    }else{
        all_distance.value=distance
    }
});

</script>

<style scoped></style>