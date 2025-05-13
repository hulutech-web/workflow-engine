<script lang="ts" setup>
import {nextTick, ref, watch} from 'vue'
import type {VxeTableEvents, VxeTableInstance} from 'vxe-table'

const props = defineProps({
  modelValue: {
    type: Array,
    default: () => [],
  },
})

const emit = defineEmits(['update:modelValue'])

interface RowVO {
  rule_title: string
  rule_name: string
  rule_value: string
}

const tableRef = ref<VxeTableInstance<RowVO>>()
const tableData = ref<RowVO[]>(props.modelValue)

watch(() => props.modelValue, (newVal) => {
  if (newVal == null || newVal.length == 0) {
    tableData.value = []
  } else {
    tableData.value = newVal
    // 从新加载一下数据
    const $table = tableRef.value
    if ($table) {
      // 设置table为可编辑状态
      $table.setEditCell((row) => {
        return row.rule_name == 'required'
      })
      $table.updateData()
    }
  }
}, {
  deep: true,
  immediate: true,
})

const ruleList = ref(
  {
    name: 'VxeSelect',
    options: [
      {title: '必填', label: '必填', value: 'required', locked: true},
      {title: '字符串', label: '字符串', value: 'string'},
      {title: '邮箱', label: '邮箱', value: 'email'},
      {title: '数字', label: '数字', value: 'uint', locked: true},
      {title: '日期', label: '日期', value: 'date', locked: true},
      {title: '最小长度', label: '最小长度', value: 'min_len'},
      {title: '最大长度', label: '最大长度', value: 'max_len'},
      {title: '最大值', label: '最大值', value: 'max'},
      {title: '最小值', label: '最小值', value: 'min'},
      {title: '不等于', label: '不等于', value: 'ne'},
      {title: '文件', label: '文件', value: 'slice', locked: true},
      {title: '图片', label: '图片', value: 'image', locked: true},
      {title: '数字大于0', label: '数字大于0', value: 'number', locked: true},
    ],
    render: (row: RowVO) => {
      // addTitleHandle(row)
    },
  //   定义选择事件
    events: {
      change (cellParams, targetParams) {
        const { row, column } = cellParams
        addTitleHandle(row)
      }
    },
  })

function removeSelectEvent() {
  const $table = tableRef.value
  if ($table) {
    $table.removeCheckboxRow()
  }
}

function insertEvent() {
  const $table = tableRef.value
  if ($table) {
    const record: RowVO = {rule_title: '', rule_name: '', rule_value: ''}
    $table.insert(record)
  }
}



function addTitleHandle(row: RowVO) {

  switch (row.rule_name) {
    case 'required':
      row.rule_title = '必填'
      break
    case 'string':
      row.rule_title = '字符串'
      break
    case 'email':
      row.rule_title = '邮箱'
      break
    case 'uint':
      row.rule_title = '数字'
      break
    case 'date':
      row.rule_title = '日期'
      break
    case 'min_len':
      row.rule_title = '最小长度'
      break
    case 'max_len':
      row.rule_title = '最大长度'
      break
    case 'max':
      row.rule_title = '最大值'
      break
    case 'min':
      row.rule_title = '最小值'
      break
    case 'ne':
      row.rule_title = '不等于'
      break
    case 'file':
      row.rule_title = '文件'
      break
    case 'image':
      row.rule_title = '图片'
      break
    case 'number':
      row.rule_title = '数字大于0'
      break
    default:
      row.rule_title = '未知'
      break
  }
  console.log(row)
}

function updateRoleList() {
  const $table = tableRef.value
  if ($table) {
    const {fullData} = $table.getTableData()
    ruleList.value.options = ruleList.value.options.filter((item) => {
      return !fullData.some(row => row.rule_name === item.value)
    })
  }
}

const editActivatedEvent: VxeTableEvents.EditActivated<RowVO> = () => {
  updateRoleList()
  // 获取表格数据
  const $table = tableRef.value
  const {fullData} = $table.getTableData()
  emit('update:modelValue', fullData)
}

function isRuleValueDisabled(ruleName: string) {
  const disabledRules = ['required', 'uint', 'date', 'file', 'image', 'number', 'email']
  return disabledRules.includes(ruleName)
}

nextTick(() => {
  updateRoleList()
})
</script>

<template>
  <div>
    <vxe-toolbar>
      <template #buttons>
        <n-button type="primary" @click="insertEvent()">
          新增
        </n-button>
        <n-button type="primary" @click="removeSelectEvent()">
          删除
        </n-button>
      </template>
    </vxe-toolbar>

    <vxe-table
      ref="tableRef" border show-overflow max-height="400" :column-config="{ resizable: true }"
      :data="tableData" :edit-config="{ trigger: 'click', mode: 'row' }"
      @edit-activated="editActivatedEvent"
    >
      <vxe-column type="checkbox" width="60"/>
      <vxe-column field="rule_title" width="140" title="规则名称" :edit-render="{ name: 'input' }"/>
      <vxe-column field="rule_name" width="140" title="规则类型" :edit-render="ruleList">

      </vxe-column>

      <vxe-column field="rule_value" width="140" title="规则值" :edit-render="{}">
        <template #edit="{ row }">
          <vxe-input
            v-model="row.rule_value" :disabled="isRuleValueDisabled(row.rule_name)"
            type="text"
          />
        </template>
      </vxe-column>
    </vxe-table>
  </div>
</template>
