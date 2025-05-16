<template>
    <div class="p-3">
        <p>
            插件
        </p>
        <a-list :grid="{ gutter: 16, xs: 1, sm: 2, md: 4, lg: 4, xl: 6, xxl: 3 }" :data-source="plugins">
            <template #renderItem="{ item }">
                <a-list-item>
                    <a-card :title="item.name">
                        <template #extra>
                            <div>
                                <a-button type="link" danger>卸载</a-button>
                            </div>
                        </template>
                        <p class="bg-orange-100 p-3 min-h-[200px]">
                        <div>
                            {{ item.version }}
                        </div>
                        <a-tooltip>
                            <template #title>
                                {{ item.description }}
                            </template>
                            <a-tag color="blue">{{ item.author }}</a-tag>
                        </a-tooltip>
                        <div>
                            {{ item.description }}
                        </div>
                        </p>
                        <template #actions>
                            <div>
                                <a-button type="link" ghost @click="designPlugin(item)">设计</a-button>
                                <a-button type="link" @click="install(item)">安装</a-button>
                                <a-button type="link" @click="uninstall(item)">卸载</a-button>
                            </div>
                        </template>
                    </a-card>
                </a-list-item>
            </template>
        </a-list>
        <a-modal :footer="false" v-model:open="open" width="1000px" title="流程选择" centered
            :bodyStyle="{ height: '590px' }">
            <vxe-grid ref='xGrid' v-bind="gridOptions" v-on="gridEvent">
                <template #action="{ row }">
                    <div>
                        <a-button type="primary" @click="bindFlow(row)">应用</a-button>
                    </div>
                </template>
            </vxe-grid>
        </a-modal>
    </div>
</template>

<script setup lang='ts'>
const { loadPlugins, installPlugin, uninstallPlugin } = usePlugin()
const { gridOptions } = useFlow()
const router = useRouter()
const route = useRoute()
const flow_id = ref(route.query.id)
const plugins = ref([])
const init = async () => {
    const { data } = await loadPlugins()
    plugins.value = data
}
const open = ref(false)
const xGrid = ref()
const gridEvent: VxeGridListeners<RowVO> = {
    proxyQuery() {
        console.log('数据代理查询事件')
        const grid = xGrid.value
        // 获取表格中的数据
        const data = grid.getTableData().fullData
    },
    proxyDelete() {
        console.log('数据代理删除事件')
    },
    proxySave() {
        console.log('数据代理保存事件')
    }
}
const submitState = ref({
    flow_id: +flow_id.value,
    plugin_id: null
})
const install = async (item: any) => {
    submitState.value.plugin_id = item.id
    const { data } = await installPlugin(submitState.value)
}

const uninstall = async (item: any) => {
    submitState.value.plugin_id = item.id
    const { data } = await uninstallPlugin(submitState.value)
}
const bindFlow = (row: any) => {
    submitState.value.flow_id = row.id
    open.value = false
    console.log(submitState.value)
}


const designPlugin = (item: any) => {
    router.push({
        path: '/admin/base/plugin/make',
        query: {
            plugin_id: item.id,
            flow_id: flow_id.value
        }
    })
}
init();
</script>

<style></style>