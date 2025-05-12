<template>
    <div>
        <a-card title="细分任务">
            <a-list item-layout="horizontal" :data-source="data">
                <template #renderItem="{ item }">
                    <a-list-item>
                        <template #actions>
                            <a-button type="primary" @click="divide">分解任务</a-button>
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
        <a-modal v-model:open="open" title="分解任务" width="600" :footer="false">
            <a-form style="margin-top: 20px;" ref="formRef" :model="activityForm">
                <a-row>
                    <a-col :span="8">
                        <a-form-item label="单位" name="Groupactivityoptions">
                            <div class="m-2">
                                <a-button type="dashed" @click="addItem" class="w-96">
                                    <PlusOutlined />
                                    添加单位
                                </a-button>
                            </div>
                            <a-col :span="8" v-for="(opt, index) in activityForm.Groupactivityoptions" :key="index">
                                <div class="flex bg-gray-100 p-2 mb-2" style="width:660px;">
                                    <a-form-item name="name" label="单位" style="width:200px">
                                        <a-input-number v-model:value="opt.number" addon-after="部门"></a-input-number>
                                    </a-form-item>
                                    <a-form-item name="percentage" label="数量" style="width:200px;margin-left:10px;">
                                        <a-input-number v-model:value="opt.percentage"
                                            @change="changePercentage(index, opt)" addon-after="单"></a-input-number>
                                    </a-form-item>
                                    <a-form-item name="current_price" label="要求成功率"
                                        style="width:200px;margin-left:10px;">
                                        <a-input-number :value="activityForm.percentage" 
                                            addon-after="%"></a-input-number>
                                    </a-form-item>
                                    <a-form-item style="width:200px;margin-left:10px;">
                                        <MinusCircleOutlined v-if="activityForm.Groupactivityoptions.length > 1"
                                            style="font-size: 18px;color:red;" class="dynamic-delete-button"
                                            @click="removeItem(opt)" />
                                    </a-form-item>
                                </div>

                            </a-col>
                        </a-form-item>

                    </a-col>
                    <a-col :span="8">
                        <a-form-item label="排序" name="sort" :rules="[
                            rulesStore.getRule('sort')
                                ? rulesStore.getRule('sort')
                                : { required: false },
                        ]">
                            <a-input-number v-model:value="activityForm.sort" addon-after="数字越大越靠前"></a-input-number>
                        </a-form-item>

                        <a-form-item label="提前多少天展示" name="before_day" :rules="[
                            rulesStore.getRule('before_day')
                                ? rulesStore.getRule('before_day')
                                : { required: false },
                        ]">
                            <a-input-number v-model:value="activityForm.before_day" addon-after="天"></a-input-number>
                        </a-form-item>

                        <a-form-item label="计划上线时间" name="start_time" :rules="[
                            rulesStore.getRule('start_time')
                                ? rulesStore.getRule('start_time')
                                : { required: false },
                        ]">
                            <a-date-picker valueFormat="YYYY-MM-DD HH:mm:ss" show-time
                                v-model:value="activityForm.start_time" />
                        </a-form-item>

                        <a-form-item label="计划结束时间" name="end_time" :rules="[
                            rulesStore.getRule('end_time')
                                ? rulesStore.getRule('end_time')
                                : { required: false },
                        ]">
                            <a-date-picker valueFormat="YYYY-MM-DD HH:mm:ss" show-time
                                v-model:value="activityForm.end_time" />
                        </a-form-item>

                        <a-form-item label="上线状态" name="status" :rules="[
                            rulesStore.getRule('status')
                                ? rulesStore.getRule('status')
                                : { required: false },
                        ]">
                            <a-radio-group size="large" v-model:value="activityForm.status">
                                <a-radio :value="2">开启</a-radio>
                                <a-radio :value="1">关闭</a-radio>
                                <a-radio :value="3">过期</a-radio>
                            </a-radio-group>
                        </a-form-item>


                    </a-col>
                </a-row>

                <a-form-item>
                    <a-button type="primary" @click="onSubmit">提交</a-button>
                </a-form-item>
            </a-form>
        </a-modal>
    </div>
</template>

<script setup lang='ts'>
import { message } from 'ant-design-vue'
const open = ref(false)
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
const addItem = () => {
    activityForm.value.Groupactivityoptions.push({
        id: null,
        number: 0,
        percentage: 0,
        current_price: 0,
    })
}
const rulesStore = useRulesStore();
const divide = (row) => {
    console.log(row)
    open.value = true
}
let activityForm = ref({
    ID: null,
    groupcourse_id: null,
    groupcourse_name: "",
    title: "",
    percentage:0,
    before_day: 15,
    require_number: 1,//报名人数
    current_number: 0,//当前报名人数
    is_group: 1,//1,未成团，2成团
    start_time: "",//开始报名时间
    end_time: "",//报名截止时间
    price: 0,
    sort: 0,
    introduce: "",
    post: "",
    status: 2,
    Groupactivityoptions: [{
        id: null,
        number: 0,
        percentage: 0,
        current_price: 0,
    }]
});
const onSubmit = async () => {
    // 如果activityForm.value.Groupactivityoptions出现number字段或者percentage为0时，提示错误
    let hasError = false;
    activityForm.value.Groupactivityoptions.forEach((item: any) => {
        if (!item.number || item.percentage === 0) {
            hasError = true;
        }
    });
    if (hasError) {
        message.error("人数或者价格比例不能为空")
        return;
    }
    try {
        //先清空一下验证
        formRef.value.clearValidate();

    } catch (error) {
        console.error(error);
    }
};
</script>

<style></style>