<template>
    <div class="login-container">
        <div class="background-image">
            <section class="login-title">
                <TypeIt :values="content" className="type-it" :cursor="false" />
            </section>

            <a-form :model="formState" name="normal_login" class="login-form" @finish="onFinish"
                :label-col="{ span: 6 }" :wrapper-col="{ span: 24 }" @finishFailed="onFinishFailed">
                <p class="text-xl text-center mb-10">用户登录</p>
                <a-form-item label="工号" name="username" :rules="[{ required: true, message: '请输入手机号！' }]">
                    <a-input v-model:value="formState.username">
                        <template #prefix>
                            <UserOutlined class="site-form-item-icon" />
                        </template>
                    </a-input>
                </a-form-item>

                <a-form-item label="密码" name="password" :rules="[{ required: true, message: '请输入密码！' }]">
                    <a-input-password v-model:value="formState.password">
                        <template #prefix>
                            <LockOutlined class="site-form-item-icon" />
                        </template>
                    </a-input-password>
                </a-form-item>
                <!-- 
                <a-form-item style="width:100%;">
                    <a-button type="dashed" @click="open = true" style="width:100%;">
                        <span class="text-red-500">点击验证</span>
                    </a-button>
                    <a-modal v-model:open="open" title="验证码" :footer="false" centered>
                        <gocaptcha-rotate style="width:100%;" :config="captcha_config" :data="captcha_data"
                            :events="captcha_events" />
                    </a-modal>
                </a-form-item>
                -->

                <a-form-item :wrapper-col="{ offset: 4, span: 24 }">
                    <a-button size="large" style="width:100%;" type="primary" html-type="submit">
                        登录
                    </a-button>
                </a-form-item>
                <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
                    <a class="login-form-forgot" href="">忘记密码</a>
                    <a-divider type="vertical" />
                    <a class="login-form-forgot" href="">立即注册</a>
                </a-form-item>
            </a-form>
            <div class="login-footer">
                Copyright © 2024 goravel流程框架 All Rights Reserved.
            </div>
        </div>

    </div>
</template>

<script setup lang="ts">
/**
 * 打字机部分
 */

import TypeIt from '@/components/typeit'
import { http } from '@/plugins/axios'
import router from '@/plugins/router';
import { message } from 'ant-design-vue';
const storage = useStorage()

let content: string[] = [
    '1. golang 高效的审批流程框架',
    '2. 面向中小企业、组织、单位提供完善的审批工作流框架，同时提供的hooks，满足个性化需求',
    '3. 内置多个插件支持',
]
const { login } = useUser()
let Spin = inject('Spin')
interface FormState {
    username: string;
    password: string;
}
const checkCaptOk = ref(false);


// 市场部-销售部-员工1
const formState = reactive<FormState>({
    tenantId: 1,
    username: 'admin',
    password: 'admin888',
});
const onFinish = async (values: any) => {
    try {
        Spin(true)
        // if (!checkCaptOk.value) {
        //     message.error("请先完成验证登录")
        //     return;
        // }
        const data = await http.request({
            url: '/account/login',
            method: 'post',
            data: formState
        })
        storage.set(CacheEnum.TOKEN_NAME, data.token)
        storage.set(CacheEnum.USER, JSON.stringify(data.userInfo))
        router.push("/admin/dashboard")
        Spin(false)
    } catch (error) {
        console.log('Failed:', error);
    } finally {
        Spin(false)
    }
};

const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
};
const disabled = computed(() => {
    return !(formState.mobile && formState.password);
});

// config = {}
interface Config {
    width?: number;
    height?: number;
    thumbWidth?: number;
    thumbHeight?: number;
    verticalPadding?: number;
    horizontalPadding?: number;
    showTheme?: boolean;
}

// data = {}
interface Data {
    image: string;
    thumb: string;
}

// events = {}
interface Events {
    click?: (x: number, y: number) => void;
    refresh?: () => void;
    close?: () => void;
    confirm?: (dots: Array<CaptchaDot>) => boolean;
}
const open = ref(false)
const captcha_config: Config = ref<Config>({})
const captcha_data: Data = ref<Data>({})
const captcha_events: Events = ref<Events>({
    click: (x, y) => {
        console.log(x, y)
    },
    refresh: async () => {
        console.log('refresh')
        await initCaptcha()
    },
    close: () => {
        console.log('close')
    },
    confirm: async (dots) => {
        console.log(dots)
        const { data } = await http.request({
            url: `captcha/validate`,
            method: 'post',
            data: { captcha_key: captcha_key.value, angle: dots }
        })
        if (data.code == 0) {
            message.success("校验成功")
            checkCaptOk.value = true
        }else{
            message.error("校验失败")
            checkCaptOk.value = false
        }
        return true
    }
})
const captcha_key = ref("")

const initCaptcha = async () => {
    const { data } = await http.request({
        url: `captcha/get`
    })
    captcha_data.value.image = data.image_base64
    captcha_data.value.thumb = data.thumb_base64
    captcha_key.value = data.captcha_key
}
// initCaptcha();
</script>

<style lang="scss" scoped>
.login-container {
    width: 100vw;
    height: 100vh;
    box-sizing: border-box;
    position: relative;

    .background-image {
        box-sizing: border-box;
        width: 100vw;
        height: 100vh;
        position: relative;
        background: url('@/assets/images/bg/login.svg') no-repeat center center;
        background-size: cover;

        .login-title {
            // 字体
            font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
            color: #1577ff;
            position: absolute;
            top: 100px;
            left: 50%;
        }

        .type-it {
            font-size: 44px;
        }

        .text-container {
            position: absolute;
            border: 1px solid red;
            width: 500px;
            min-height: 200px;
            height: auto;
            top: 30px;
            left: 50%;


        }

        .login-form {
            border-radius: 7% 7% 8% 15% / 5% 6% 10% 53%;
            transition: all 0.5s;
            // 背景磨砂效果
            backdrop-filter: blur(10px);
            background-color: rgba(255, 255, 255, 0.8);
            box-sizing: border-box;

            // 当hover的时候
            &:hover {
                box-shadow:
                    2.8px 2.8px 2.2px rgba(0, 0, 0, 0.058),
                    6.7px 6.7px 5.3px rgba(0, 0, 0, 0.083),
                    12.5px 12.5px 10px rgba(0, 0, 0, 0.099),
                    22.3px 22.3px 17.9px rgba(0, 0, 0, 0.112),
                    41.8px 41.8px 33.4px rgba(0, 0, 0, 0.125),
                    100px 100px 80px rgba(0, 0, 0, 0.15);
            }

            box-shadow: 2.8px 2.8px 2.2px rgba(0, 0, 0, 0.02),
            6.7px 6.7px 5.3px rgba(0, 0, 0, 0.028),
            12.5px 12.5px 10px rgba(0, 0, 0, 0.035),
            22.3px 22.3px 17.9px rgba(0, 0, 0, 0.042),
            41.8px 41.8px 33.4px rgba(0, 0, 0, 0.05),
            100px 100px 80px rgba(0, 0, 0, 0.07);
            padding: 30px;
            box-sizing: border-box;
            position: absolute;
            right: 400px;
            top: 30%;
            width: 400px;
            height: 400px;
        }

        .login-footer {
            display: block;
            width: 100%;
            height: 24px;
            background-color: black;
            color: white;
            position: absolute;
            line-height: 24px;
            text-align: center;
            font-size: 12px;
            bottom: 0px;
            box-sizing: border-box;
        }

        /* 替换为你的SVG文件路径 */
    }

}
</style>