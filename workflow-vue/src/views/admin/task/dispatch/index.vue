<template>
    <div>
        <a-card title="任务派发">
            <a-list item-layout="horizontal" :data-source="data">
                <template #renderItem="{ item }">
                    <a-list-item>
                        <template #actions>
                            <a-button type="primary" @click="message.success('派发成功')">派发</a-button>
                        </template>
                        <a-list-item-meta :description="item.desc">
                            <template #title>
                                <a href="https://www.antdv.com/">{{ item.title }}</a>
                            </template>
                            <template #avatar>
                                <a-avatar src="https://joeschmoe.io/api/v1/random" />
                            </template>

                        </a-list-item-meta>
                        {{ item.content }}
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
        desc: "家庭宽带50M以下用户，可升级100M",
        content:"宣汉1000，万源2000，开江1500，通川1000，渠县3000"
    },
    {
        title: '流量溢出用户',
        desc: "套内流量溢出超过200m",
        content:"宣汉1000，万源2000，开江1500，通川1000，渠县3000"

    },
    {
        title: '网速投诉用户',
        desc: "家庭宽带卡顿延时，需更新宽带终端设备",
        content:"宣汉1000，万源2000，开江1500，通川1000，渠县3000"

    },
    {
        title: 'WIFI故障用户',
        desc: "家庭宽带WIFI覆盖低，需更新WIFI设备",
       content:"宣汉1000，万源2000，开江1500，通川1000，渠县3000"

    },
];

</script>

<style></style>