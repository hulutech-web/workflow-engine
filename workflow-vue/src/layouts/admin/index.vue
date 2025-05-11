<script setup lang='ts'>
import Left from '@/layouts/admin/Left.vue'
import Navbar from '@/layouts/admin/navbar.vue'
import useLoading from '@/store/useLoadingStore'
import useTipsStore from '@/store/useTipsStore'
const loadingStore = useLoading()
const tipsStore = useTipsStore()
const triggleRef = ref()
const collapse = (collapsed, type) => {
    triggleRef.value.active = !triggleRef.value.active
}
const collapsed = ref(false)
const changeCollapse = (val) => {
    collapsed.value = val
    if(collapsed.value){
        triggleRef.value.active = false
    }else{
        triggleRef.value.active = true
    }
}
</script>

<template>
    <main>
        <a-layout style='min-height: 100vh'>
            <a-layout-sider v-model:collapsed="collapsed" :trigger="null" collapsible>
                <Left ref="triggleRef" />
            </a-layout-sider>
            <a-layout>
                <a-layout-header style='background: #fff;height:80px;padding-inline:20px;'>
                    <Navbar @changeCollapse="changeCollapse" />
                </a-layout-header>
                <a-layout-content style='position:relative;' class='layout-content'>
                    <div id='admin_content' style="padding:15px;box-sizing: border-box;">
                        <section style="background-color: white;min-height: 800px;padding: 5px;box-sizing: border-box;">
                            <router-view />
                        </section>
                        <div id='spinDiv'>
                        </div>
                    </div>
                    <a-back-top />
                </a-layout-content>
                <a-layout-footer style='text-align: center'>
                    <div class='text-black' style="font-size: 12px;;">
                        goravel流程框架 ©2024 Created by HuluTech
                    </div>
                </a-layout-footer>

            </a-layout>
        </a-layout>
        <Teleport to='body'>
            <div v-if='loadingStore.getLoading' class='absolute top-1/2 left-1/2'>
                <a-spin size='large' />
            </div>
        </Teleport>
    </main>
</template>

<style></style>
