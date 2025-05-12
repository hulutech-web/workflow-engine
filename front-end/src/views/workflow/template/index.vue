<script setup lang="ts">
import Formpart from './formpart.vue'
import useRulesStore from '@/store/useRulesStore.ts'

const { loadTemplates, gridOptions, storeTemplate, deleteTemplate, updateTemplate } = useTemplate()
const { loadTemplateForm, deleteTemplateForm } = useTemplateForm()
const router = useRouter()
const xGrid = ref()
const rulesStore = useRulesStore()
// 模板管理
const templateRef = ref()
const templateState = ref({
  template_name: '',
})
const templateUpdateState = ref({
  id: null,
  template_name: '',
})
const cid = ref()
async function submitTemplate() {
  try {
    // 先清空一下验证
    // templateRef.value.clearValidate()
    await storeTemplate(templateState.value)
    // templateRef.value.resetFormField();
  }
  catch (error) {
    templateRef.value.validate()
  }
  // xGrid.value.commitProxy("query")
}
const editOpen = ref(false)
function editTemplate(row) {
  // console.log(row)
  editOpen.value = true
  templateUpdateState.value = row
}
const templateUptRef = ref()
const formpartRef = ref()
async function submitUpdateTemplate() {
  try {
    // 先清空一下验证
    templateUptRef.value.clearValidate()
    await updateTemplate(templateUpdateState.value)
    templateUptRef.value.resetFormField()
  }
  catch (error) {
    templateUptRef.value.validate()
  }
  xGrid.value.commitProxy('query')
  formpartRef.value.loadTemplateOpts()
  editOpen.value = false
}

// 模板管理

const gridEvent: VxeGridListeners<RowVO> = {
  proxyQuery() {
    console.log('数据代理查询事件')
    const grid = xGrid.value
    // 获取表格中的数据
    const data = grid.getTableData().fullData
    console.log(data)
  },
  proxyDelete() {
    console.log('数据代理删除事件')
  },
  proxySave() {
    console.log('数据代理保存事件')
  },
}
const columns = [
  {
    title: 'id',
    dataIndex: 'id',
    key: 'id',
  },
  {
    title: '控件名称',
    dataIndex: 'field_name',
    key: 'field_name',
  },
  {
    title: '字段名',
    dataIndex: 'field',
    key: 'field',
  },
  {
    title: '字段类型',
    dataIndex: 'field_type',
    key: 'field_type',
  },
  {
    title: '排序',
    dataIndex: 'sort',
    key: 'sort',
  },
  {
    title: '模板id',
    dataIndex: 'template_id',
    key: 'template_id',
  },
  {
    title: '操作',
    key: 'action',
    customRender: ({ record }) => h('div', [
      h(ElButton, {
        size: 'small',
        onClick: () => edit(record),
      }, '编辑'),
      h(ElPopconfirm, {
        title: '确认删除？',
        onConfirm: () => delRecord(record),
      }, {
        default: () => h(ElButton, {
          size: 'small',
          type: 'danger',
          style: 'margin-left: 10px',
        }, '删除'),
      }),
    ]),
  },
]
const fields = ref([])

async function loadTmplForm(row) {
  const data = await loadTemplateForm(row.id)
  fields.value = data
}

const open = ref(false)
function edit(record) {
  cid.value = record.id
  open.value = true
}

const templateOpts = ref([])
async function loadTemplateOpts() {
  const { data } = await loadTemplates()
  templateOpts.value = data
}

async function delRecord(row) {
  await deleteTemplateForm(row.id)
}
</script>

<template>
  <div>
    <n-grid :gutter="[24, 1]">
      <n-gi :span="15">
        <n-card class="mb-3">
          <n-form ref="templateRef" :model="templateState" layout="inline">
            <n-form-item
              label="模板名称" name="template_name"
              :rules="[rulesStore.getRule('template_name') ? rulesStore.getRule('template_name') : { required: false }]"
            >
              <n-input v-model:value="templateState.template_name" />
            </n-form-item>
            <n-form-item>
              <n-button type="primary" @click="submitTemplate">
                添加
              </n-button>
            </n-form-item>
          </n-form>
          <vxe-grid ref="xGrid" v-bind="gridOptions" v-on="gridEvent">
            <template #action="{ row }">
              <div>
                <n-button-group>
                  <n-popconfirm
                    title="将同步删除模板字段，确认删除吗？" ok-text="是" cancel-text="点错了"
                    @confirm="deleteTemplate(row)"
                  >
                    <n-button size="small" danger type="primary">
                      删除
                    </n-button>
                  </n-popconfirm>
                  <n-button type="primary" size="small" @click="editTemplate(row)">
                    编辑
                  </n-button>
                  <n-button type="primary" size="small" @click="loadTmplForm(row)">
                    表单控件
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
        </n-card>
        <n-modal v-model:show="editOpen" width="500px" title="修改模板" centered>
          <n-form ref="templateUptRef" :model="templateUpdateState">
            <n-form-item
              label="模板名称" name="template_name"
              :rules="[rulesStore.getRule('template_name') ? rulesStore.getRule('template_name') : { required: false }]"
            >
              <n-input v-model:value="templateUpdateState.template_name" />
            </n-form-item>
            <n-form-item>
              <n-button type="primary" @click="submitUpdateTemplate">
                保存
              </n-button>
            </n-form-item>
          </n-form>
        </n-modal>

        <n-card>
          <n-data-table bordered :columns="columns" :data-source="fields">
            <template #bodyCell="{ column, record }">
              <template v-if="column.dataIndex === 'action'">
                <div>
                  <n-button-group>
                    <n-button size="small" type="primary" @click="edit(record)">
                      编辑
                    </n-button>
                    <n-popconfirm title="Title" @confirm="delRecord(record)">
                      <n-button size="small" type="primary" danger>
                        删除
                      </n-button>
                    </n-popconfirm>
                  </n-button-group>
                </div>
              </template>
            </template>
          </n-data-table>
        </n-card>
      </n-gi>

      <n-gi :span="9">
        <n-card>
          <p>设计字段</p>
          <Formpart ref="formpartRef" />
        </n-card>
      </n-gi>
    </n-grid>

    <n-modal v-model:show="open" title="控件配置" centered width="800px" :footer="false">
      <Formpart :id="cid" ref="formpartRef" />
    </n-modal>
  </div>
</template>

<style scoped></style>
