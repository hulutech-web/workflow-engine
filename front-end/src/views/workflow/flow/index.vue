<template>
  <div>
    <n-card>
      <div>
        <n-button type="primary" @click="toAdd">新建流程</n-button>
        <vxe-grid ref='xGrid' v-bind="gridOptions" v-on="gridEvent">
          <template #publish="{ row }">
            <div>
              {{ row.is_publish == true ? "已发布" : "未发布" }}
            </div>
          </template>
          <template #show="{ row }">
            <div>
              <div>
                {{ row.show == true ? "显示" : "隐藏" }}
              </div>
            </div>
          </template>
          <template #design="{ row }">
            <div>
              <n-button type="primary" @click="toDesign(row.id)">
                管理流程图
              </n-button>
            </div>
          </template>
          <template #action="{ row }">
            <div>
              <n-button-group>
                <n-button type="primary" @click="editFlow(row)">编辑</n-button>
                <n-button type="primary" danger>删除</n-button>
                <n-button type="primary" @click="startFlow(row)" :disabled="row.is_publish == false">
                  发起流程
                </n-button>
                <n-button type="primary" @click="startPlugin(row)">
                  插件功能
                </n-button>
              </n-button-group>
            </div>
          </template>
          <template #dept="{ row }">
            <div>
              {{ row.Dept.id == 0 ? "未分配" : row.Dept.dept_name }}
            </div>
          </template>
        </vxe-grid>
      </div>

    </n-card>
  </div>

</template>

<script setup lang="ts">
import {useMessage} from 'naive-ui'

const message = useMessage();
const {gridOptions} = useFlow()
const router = useRouter();
// TABLE
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

const toDesign = (id) => {
  console.log(id)
  router.push({path: `/workflow/flow/${id}/design`})
}

const toAdd = () => {
  router.push({path: `/workflow/flow/create`})
}

const startFlow = (row) => {
  if (row.is_publish == false) {
    message.error("流程尚未发布，无法发起流程")
    return
  }
  router.push({path: `/workflow/flow/${row.id}/initiation`})
}


const editFlow = (row) => {
  router.push({path: `/workflow/flow/${row.id}/edit`})
}
const startPlugin = (row) => {
  router.push({path: `/admin/base/plugin/index`, query: {id: row.id}})
}
</script>

<style scoped></style>
