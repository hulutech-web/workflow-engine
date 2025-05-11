<template>
    <div>
        <a-card title="领取任务">
            <a-list item-layout="horizontal" :data-source="data">
                <template #renderItem="{ item }">
                    <a-list-item>
                        <template #actions>
                            <a-button type="primary" @click="message.success('领用成功')">领用</a-button>
                            <span v-for="{ icon, text } in actions" :key="icon">
                                <component :is="icon" style="margin-right: 8px" />
                                {{ text }}
                            </span>
                        </template>
                        <a-list-item-meta :description="item.desc">
                            <template #title>
                                <a href="https://www.antdv.com/">{{ item.title }}</a>
                            </template>
                            <template #avatar>
                                <a-avatar src="https://joeschmoe.io/api/v1/random" />
                            </template>

                        </a-list-item-meta>
                    </a-list-item>
                </template>
            </a-list>
        </a-card>
    </div>
</template>

<script setup lang='ts'>
import { message } from 'ant-design-vue'
import {StarOutlined,
LikeOutlined,
MessageOutlined} from '@ant-design/icons-vue'
const data = [
    {
        title: '家庭宽带提速用户',
        desc: "家庭宽带50M以下用户，可升级100M"
    },
    {
        title: '流量溢出用户',
        desc: "套内流量溢出超过200m"
    },
    {
        title: '网速投诉用户',
        desc: "家庭宽带卡顿延时，需更新宽带终端设备"
    },
    {
        title: 'WIFI故障用户',
        desc: "家庭宽带WIFI覆盖低，需更新WIFI设备"
    },
];
const actions = [
  { icon: StarOutlined, text: '总数12000' },
  { icon: LikeOutlined, text: '已使用10000' },
  { icon: MessageOutlined, text: '未使用2000' },
];
</script>

<style></style>