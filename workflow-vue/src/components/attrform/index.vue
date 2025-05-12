<template>
  <div>
    <div>
      <a-tabs v-model:activeKey="activeKey" type="card">
        <!-- 常规设置 -->
        <a-tab-pane key="1" tab="常规" style="height: 240px" v-if="formState.process">
          <a-form-item label="步骤名称">
            <a-input v-model:value="submitState.process_name" :value="formState.process.process_name"></a-input>
          </a-form-item>
          <a-form-item label="步骤类型">
            <a-radio-group v-model:value="submitState.process_position" :value="formState.process.position">
              <a-radio :value="1">正常步骤</a-radio>
              <a-radio v-if="formState.can_child" :value="2">转入子流程</a-radio>
              <!-- <a-radio  :value="2">转入子流程</a-radio> -->
              <a-radio :value="0">第一步</a-radio>
            </a-radio-group>
          </a-form-item>
          <a-divider></a-divider>
          <!--          非转入子流程-->
          <div
            v-if="formState.next_process && (submitState.process_position == 1 || submitState.process_position == 0)">
            <div class="flex">
              <div class="px-3 flex flex-col item-center justify-center">
                <ArrowRightOutlined />
              </div>
              <div class="flex-1  border-blue-600 border border-solid  p-2">
                <p class="text-md">下一步骤</p>
                <a-row>
                  <a-col :span="16" v-for="(p, index) in formState.next_process" :key="index">
                    <a-tag :bordered="false" color="geekblue" v-if="p.NextProcess.id != -1">
                      {{ p.NextProcess.process_name }}
                    </a-tag>
                  </a-col>
                </a-row>
              </div>

              <!-- <div class="flex-1  border-blue-500 border border-dotted  p-2">
                <p class="text-md">其他步骤</p>
                <a-row>
                  <a-col :span="16" v-for="(p, index) in formState.beixuan_process" :key="index">
                    <a-tag :bordered="false" color="geekblue" v-if="p.NextProcess.id != -1">
                      {{ p.Process.process_name }}
                    </a-tag>
                  </a-col>
                </a-row>
              </div> -->
            </div>

          </div>

          <!--          转入子流程-->

          <div v-if="formState.next_process">
            <div id="child_flow" v-if="formState.process.position == 2">
              <div class="control-group">
                <a-form-item label="子流程">
                  <a-select v-model:value="submitState.child_flow_id" :value="formState.process.child_flow_id">
                    <a-select-option value="0">请选择</a-select-option>

                    <a-select-option v-for="(flow, ind) in formState.flows" :key="ind" :value="flow.id"
                      :selected="formState.process.child_flow_id == flow.id">
                      {{ flow.flow_name }}
                    </a-select-option>
                  </a-select>
                </a-form-item>

              </div>

              <div class="control-group">
                <a-form-item label="子流程结束后动作">
                  <a-radio-group v-model:value="submitState.child_after" :value="formState.process.child_after">
                    <a-radio :value="1">
                      同时结束父流程
                    </a-radio>
                    <a-radio :value="2">
                      返回父流程步骤
                    </a-radio>
                  </a-radio-group>
                </a-form-item>

              </div>

              <div v-if="submitState.child_after == 2">
                <a-form-item label="返回步骤">
                  <a-select name="child_back_process" v-model:value="submitState.child_back_process"
                    :value="formState.child_back_process">
                    <a-select-option value="0">
                      无
                    </a-select-option>
                    <a-select-option v-for="(p, index) in formState.processes" :key="index" :value="p.id"
                      :selected="p.child_back_process == p.id">
                      {{ p.process_name }}
                    </a-select-option>
                  </a-select>
                </a-form-item>
                <span class="help-inline">默认为当前步骤下一步</span>
              </div>
            </div>
          </div>
        </a-tab-pane>


        <a-tab-pane key="2" tab="表单" style="height: 240px">
          <a-table bordered :columns="columns" :dataSource="dataSource"></a-table>
        </a-tab-pane>


        <a-tab-pane key="3" tab="权限" style="height: 240px">
          <div>
            <p>
            <div>
              <a-form-item label="自动选人">
                <a-select v-model:value="submitState.auto_person" @change="changeAuto">
                  <a-select-option value="0">
                    不自动选人
                  </a-select-option>
                  <a-select-option value="-1001">
                    发起人部门主管
                  </a-select-option>
                  <a-select-option value="-1002">
                    发起人部门经理
                  </a-select-option>
                </a-select>
              </a-form-item>
            </div>
            </p>
            <a-divider>
              授权范围（适用于：当需要手动选人时，则授权范围生效）
            </a-divider>

            <div>
              <span>授权人员：</span>
              <a-select v-model:value="submitState.range_emp_ids" mode="tags" placeholder="选择人员" style="width:400px;">
                <a-select-option v-for="(i,ind) in submitState.range_emp_ids" :key="ind" :value="i">
                  {{ submitState.range_emp_text[ind] }}
                </a-select-option>
              </a-select>
              <a-button type="primary" @click="selPer" :disabled="disableAuto">选择</a-button>
            </div>

            <div class="mt-3">
              <span>授权部门：</span>

              <a-select v-model:value="submitState.range_dept_ids" mode="tags" placeholder="选择部门" style="width:400px;">
                <a-select-option v-for="(i,ind) in submitState.range_dept_ids" :key="ind" :value="i">
                  {{ submitState.range_dept_text[ind] }}
                </a-select-option>
              </a-select>
              <a-button :disabled="disableAuto" type="primary" @click="selDep">选择</a-button>
            </div>

            <a-modal :footer="false" v-model:open="open" width="1000px" title="人员&部门选择" centered
              :bodyStyle="{ height: '590px' }">
              <a-table :row-selection="{ selectedRowKeys: state.selectedRowKeys, onChange: onSelectChange }" rowKey="id"
                bordered :columns="tabcolumns" :dataSource="depts" :pagination="false" v-if="selectedEmp == false">
                <template #bodyCell="{ column, text, record }">
                  <template v-if="column.dataIndex === 'html'">
                    <span class="text-blue-500 text-xl mr-3">
                      {{ record.html }}
                    </span>
                    {{ record.dept_name }}
                  </template>
                  <template v-if="column.dataIndex === 'Manager'">
                    {{ record.Manager ? record.Manager.name : '' }}
                  </template>
                  <template v-if="column.dataIndex === 'Director'">
                    {{ record.Director ? record.Director.name : '' }}
                  </template>
                </template>
              </a-table>

              <vxe-grid ref='xGrid' v-bind="gridOptions" v-on="gridEvent" v-if="selectedEmp == true">
                <template #checkbox_header="{ checked, indeterminate }">
                  <div>选择</div>
                </template>
                <template #checkbox_cell="{ row, checked, indeterminate }">
                  <span class="custom-checkbox" @click.stop="toggleCheckboxEvent(row)">
                    <a-checkbox v-if="indeterminate" :checked="checked"></a-checkbox>
                    <a-checkbox v-else-if="checked" :checked="checked"></a-checkbox>
                    <a-checkbox v-else></a-checkbox>
                  </span>
                </template>
                <template #dept="{ row }">
                  <div>
                    {{ row.Dept.id == 0 ? "未分配" : row.Dept.dept_name }}
                  </div>
                </template>
              </vxe-grid>
            </a-modal>


          </div>
        </a-tab-pane>


        <a-tab-pane key="4" tab="转出条件">
          <div style="height: 700px;">
            <a-row>
              <a-col :span="4">
                <div class="text-md font-bold">
                  转出步骤
                </div>
              </a-col>
              <a-col :span="6">
                <div class="text-md font-bold">
                  转出条件
                </div>
              </a-col>
              <a-col :span="14">
                <div class="text-md font-bold">
                  更改规则
                  <a-alert message="注意：填写完规则后请完成校验！！！" type="info" />
                </div>
              </a-col>
            </a-row>
            <div style="height:10px;"></div>
            <a-row v-for="(item, index) in formState.next_process" :key="index"
              v-if="formState.next_process.length > 1">
              <a-col :span="4">
                <div class="show-item">
                  {{ item.NextProcess.process_name }}
                </div>
              </a-col>
              <a-col :span="6">
                <div class="show-expr">
                  <div v-if="!item.Expression">
                    <span class="text-sm text-gray-400">暂无条件</span>
                  </div>
                  <div v-else>
                    <div v-for="(e, i) in JSON.parse(item.Expression)" :key="i">
                      {{ e.field }}{{ e.operator }}{{ e.value }}{{ e.extra }}
                    </div>
                  </div>
                </div>
              </a-col>
              <a-col :span="14" style="padding:6px;background-color:#fafafa;border-bottom:1px solid orange">
                <a-row>
                  <a-col :span="4">
                    <div class="text-center">字段</div>
                    <a-select style="width: 100%;" v-model:value="bindExprs[index]['field']">
                      <a-select-option :value="f.field" v-for="(f, index) in fields" :key="index">
                        {{ f.field_name }}
                      </a-select-option>
                    </a-select>
                  </a-col>
                  <a-col :span="4">
                    <div class="text-center">条件</div>
                    <a-select style="width: 100%;" v-model:value="bindExprs[index]['operator']">
                      <a-select-option value=">">
                        大于
                      </a-select-option>
                      <a-select-option value="<">
                        小于
                      </a-select-option>
                      <a-select-option value="=">
                        等于
                      </a-select-option>
                      <a-select-option value="<=">
                        小于等于
                      </a-select-option>
                      <a-select-option value=">=">
                        大于等于
                      </a-select-option>
                    </a-select>
                  </a-col>
                  <a-col :span="4">
                    <div class="text-center">值</div>
                    <a-input v-model:value="bindExprs[index]['value']"></a-input>
                  </a-col>
                  <a-col :span="4">
                    <div class="text-center">其他条件</div>
                    <a-select style="width: 100%;" v-model:value="bindExprs[index]['extra']">
                      <a-select-option value="" default>
                        无
                      </a-select-option>
                      <a-select-option value="AND">
                        并且
                      </a-select-option>
                      <a-select-option value="OR">
                        或者
                      </a-select-option>
                    </a-select>
                  </a-col>

                  <a-col :span="4">
                    <div>
                      <div class="text-center">操作</div>
                    </div>
                    <div class="text-center">
                      <a-button-group>
                        <a-button type="primary" @click="addCondi(index)">新增</a-button>
                      </a-button-group>
                    </div>
                  </a-col>
                  <a-col :span="4">
                    <div class="text-center">确认条件</div>
                    <div class="text-center">
                      <a-button type="primary" @click="validateExpr(index)">确认</a-button>
                    </div>
                  </a-col>
                </a-row>
                <template v-for="(sE, ind) in stateExprs[index]" :key="ind">
                  <div class="expr" v-if="index == sE['index']">
                    <a-row>
                      <a-col :span="4">
                        <div class="text-center"> {{ sE['field'] }}</div>
                      </a-col>
                      <a-col :span="4">
                        <div class="text-center">{{ sE['operator'] }}</div>
                      </a-col>
                      <a-col :span="4">

                        <div class="text-center">{{ sE['value'] }}</div>
                      </a-col>
                      <a-col :span="4">
                        <div class="text-center">{{ sE['extra'] }}</div>
                      </a-col>
                      <a-col :span="4">
                        <a-space>
                          <MinusCircleOutlined class="cursor-pointer ml-4" style="font-size: 16px; color:red"
                            @click="delExpr(index, ind)" />
                        </a-space>
                      </a-col>
                    </a-row>
                  </div>
                </template>

              </a-col>
            </a-row>


          </div>

        </a-tab-pane>


        <a-tab-pane key="5" tab="样式" style="height: 240px">
          <div class="p-3">
            <div class="flex justify-start items-center mt-3">
              <div class="flex-4" style="width:80px;">尺寸</div>
              <div class="flex-8 flex">
                <a-space>
                  <a-input-number v-model:value="submitState.style_width"></a-input-number>
                  X
                  <a-input-number v-model:value="submitState.style_height"></a-input-number>
                </a-space>
              </div>
            </div>
            <div class="flex mt-3 items-center">
              <div class="flex-4" style="width:80px;">字体颜色</div>
              <input type="text" v-model="submitState.style_color"
                class="w-24 h-8 border-none outline-none bg-gray-100 rounded-sm px-3 mx-3">
              <div v-for="(c, ind) in colors" :key="ind" :style="{ background: `${c}` }"
                class="h-8 w-8 cursor-pointer hover:scale-105" @click="setColor(c)"></div>

            </div>
            <div class="flex mt-3 items-center">
              <div style="width:80px;">图标</div>
              <div>
                <HuluIcon :name="submitState.style_icon" :fontSize="'24px'" style="line-height:24px"
                  class="cursor-pointer bg-black text-white  rounded w-6" />
              </div>
              <input type="text" class="h-8 border-none outline-none bg-gray-100 rounded-sm px-3 flex-2"
                v-model="submitState.style_icon">

              <div style="width:600px;background-color: black;line-height:24px;" class="ml-4 flex flex-wrap ">
                <HuluIcon @click="onMyIcon(ic)" fontSize="24px" :name="ic" v-for="(ic, index) in MyIcons" :key="index"
                  class="m-3 cursor-pointer hover:scale-125" />
              </div>
            </div>
          </div>
        </a-tab-pane>

      </a-tabs>

      <div class="absolute bottom-0 left-0 ml-5 mb-5">
        <a-button type="primary" @click="onSubmit">
          <SendOutlined />
          提交
        </a-button>
      </div>
    </div>

  </div>

</template>

<script lang='ts'>

import { icons } from './icon'
import useEmpconfig from './empconfig'
import { message } from "ant-design-vue";
import { ExplainConditionSql } from "@/components/attrform/sql/explain";
const { getCurrCond } = useProcess()
const { gridOptions } = useEmpconfig()
const { loadDepts } = useDept()


export default {
  props: ['attrs'],
  emits: ["updProcess"],
  setup(props, context) {
    // #region 常规
    const MyIcons = ref(icons)
    const columns = [
      {
        title: '字段名称',
        dataIndex: 'field_name',
        key: 'field_name',
      },
      {
        title: '字段标识',
        dataIndex: 'field',
        key: 'field',
      },
      {
        title: '字段类型',
        dataIndex: 'field_type',
        key: 'field_type',
      },
    ];

    const submitState = ref({
      process_name: "",
      process_position: "",
      auto_person: "0",
      process_to: [],
      child_flow_id: "",
      child_after: "",
      range_emp_ids: [],
      range_emp_text: [],
      range_dept_ids: [],
      range_dept_text: [],
      range_role_ids: [],
      range_role_text: [],
      process_mode: "",
      con_sign: "",
      con_sign_vsb: "",
      process_in_set: "",
      process_condition: []<Expression>,
      style_width: "",
      style_height: "",
      style_color: "",
      style_icon: "",
    })
    const dataSource = computed(() => {
      return props.attrs.fields
    })
    const formState = ref(props.attrs)



    watch(() => props.attrs, (newVal, oldVal) => {
      if (newVal.process != oldVal.process) {
        console.log(newVal)
        formState.value = newVal
        fillSubmitState(newVal)
        initBase(newVal)
        initExprs()
        initStyle(newVal)
        initPer(newVal)
      }
    })

    const initBase = (attrs) => {
      submitState.value.child_flow_id = attrs.process.child_flow_id
      submitState.value.child_after = attrs.process.child_after
      submitState.value.child_back_process = attrs.process.child_back_process
    }

    const fillSubmitState = (attrs) => {
      submitState.value.process_name = attrs.process.process_name
      submitState.value.process_position = attrs.process.position
      submitState.value.process_to = attrs.next_process.map(item => item.NextProcess.id)
    }
    const activeKey = ref('1');

    const onSubmit = () => {
      context.emit("updProcess", submitState.value)
    }

    // 下一步子流程

    const tmpNextProcess = ref({})
    const tmpBeixuanProcess = ref({})
    const removePrs = (item) => {
      let tmpIndex = nextProcesses.value.findIndex(b => b.id == item.id)
      if (tmpIndex != -1) {
        tmpNextProcess.value = nextProcesses.value[tmpIndex]
        nextProcesses.value.splice(tmpIndex, 1)
        beixuanProcess.value.push(tmpNextProcess.value)
      }
    }

    const addPrs = (item) => {
      let tmpIndex = beixuanProcess.value.findIndex(b => b.id == item.id)
      if (tmpIndex != -1) {
        tmpBeixuanProcess.value = beixuanProcess.value[tmpIndex]
        beixuanProcess.value.splice(tmpIndex, 1)
        nextProcesses.value.push(tmpBeixuanProcess.value)
        console.log(nextProcesses.value)
      }
    }
    // #endregion 常规


    // #region 权限

    const initPer = (attrs) => {
      submitState.value.auto_person = attrs.sys
      submitState.value.range_emp_ids = attrs.select_emps.map(item => item.id)
      submitState.value.range_emp_text = attrs.select_emps.map(item => item.name)
      submitState.value.range_dept_ids = attrs.select_depts.map(item => item.id)
      submitState.value.range_dept_text = attrs.select_depts.map(item => item.dept_name)
    }

    const open = ref(false)
    const depts = ref([])
    const xGrid = ref()
    const toggleAllCheckboxEvent = () => {
      const $grid = xGrid.value
      if ($grid) {
        $grid.toggleAllCheckboxRow()
      }
    }

    const disableAuto = ref(props.attrs.sys == '0')
    const changeAuto = (val) => {
      if (val == '-1001') {
        disableAuto.value = true
      }
      if (val == '-1002') {
        disableAuto.value = true
      }
      if (val == "0") {
        disableAuto.value = false
      }
    }
    const selectRecords = ref([])
    const toggleCheckboxEvent = (row) => {
      const $grid = xGrid.value
      if ($grid) {
        $grid.toggleCheckboxRow(row)
        // 获取所有已经选择的项目
        const records = $grid.getCheckboxRecords()
        submitState.value.range_emp_ids = records.map(item => item.id)
        submitState.value.range_emp_text = records.map(item => item.name)
      }
    }

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

    const selectedEmp = ref(false)
    const selPer = () => {
      selectedEmp.value = true
      open.value = true
    }
    const selDep = async () => {
      open.value = true
      selectedEmp.value = false
      const { data } = await loadDepts()
      depts.value = data
    }


    const tabcolumns = [
      {
        title: '层级',
        dataIndex: 'html',
        key: 'html',
      },
      {
        title: '经理',
        dataIndex: 'Manager',
        key: 'Manager',
      },
      {
        title: '负责人',
        dataIndex: 'Director',
        key: 'Director',
      },
    ]


    const state = reactive({
      selectedRowKeys: [],
      // Check here to configure the default column
      loading: false,
    });
    const onSelectChange = (selectedRowKeys, selectedRows) => {
      console.log('selectedRowKeys changed: ', selectedRowKeys);
      state.selectedRowKeys = selectedRowKeys;
      submitState.value.range_dept_ids = selectedRowKeys
      // 获取被选择的选项数据
      console.log(selectedRows)
      submitState.value.range_dept_text = selectedRows.map(item=>item.dept_name)
    };
    // #endregion 权限

    //  #region 转出条件
    const fields = ref([])
    const nextProcesses = ref([])
    //Expression类型的二维数组
    interface Expression {
      id: number,
      index: number,
      field: string,
      operator: string,
      value: string,
      extra: string,
    }

    const bindExprs = ref<Expression>([])
    const stateExprs = ref([])

    const transText = (exp) => {
      if (exp && exp.startsWith("$")) {
        let text = exp.slice(1)
        let fieldItem = fields.value.find(item => text.includes(item['field']))
        if (fieldItem) {
          //按照空格拆分，取除开第一项的
          let arr = text.split(" ")
          let suffix = arr.slice(1).join(" ")
          return `${fieldItem['field_name']}${suffix}`
        }
      } else {
        return exp
      }

    }

    const addCondi = (index) => {
      let keys = ['field', 'operator', 'value']
      console.log("bindExprs.value[index]", bindExprs.value[index])
      if (keys.some(i => bindExprs.value[index][i] === '') == true) {
        return message.error("请填写完整")
      }
      if (bindExprs.value[index]['index'] == index) {
        stateExprs.value[index] = stateExprs.value[index] || []
        stateExprs.value[index].push(bindExprs.value[index])
        bindExprs.value[index] = {
          id: bindExprs.value[index]['id'],
          index: index,
          field: "",
          operator: "",
          value: "",
          extra: "",
        }
      }
    }



    const validateExpr = (index) => {
      let targetArr = stateExprs.value[0].filter(item => item.index == index)
      const {
        success,
        msg
      } = ExplainConditionSql(targetArr)
      if (success == false) {
        message.error(msg)
      } else {
        message.success("条件校验成功")
        submitState.value.process_condition = stateExprs.value[0]
      }
    }


    const getCurrentCond = async (process) => {
      let param = {
        flow_id: process.FlowID,
        process_id: process.ProcessID,
        next_process_id: process.NextProcessID
      }
      await getCurrCond(param)
    }
    const delExpr = (index, ind) => {
      stateExprs.value[index].splice(ind, 1)
    }

    const initExprs = () => {
      fields.value = formState.value.fields
      //初始化表达式
      formState.value.next_process.map((item, index) => {
        bindExprs.value.push({
          id: item.id,
          index: index,
          field: "",
          operator: "",
          value: "",
          extra: "",
        })
      })
      stateExprs.value = new Array(formState.value.next_process.length).fill(new Array())


      formState.value.next_process.map((item, index) => {
        getCurrentCond(item)
      })
    }

    // #endregion 转出条件

    // #region 样式

    const onMyIcon = (icon) => {
      submitState.value.style_icon = icon
    }
    // 定义一组颜色
    const colors = ref([
      '#FFA500', '#800080', '#FFC0CB', '#A0522D', '#6B8E23', '#483D8B', '#D2691E', '#9400D3',
      '#228B22', '#7B68EE', '#B22222', '#8B4726', '#556B2F', '#4B0082', '#9932CC', '#6A5ACD'])

    const setColor = (color) => {
      submitState.value.style_color = color
    }
    const initStyle = (attrs) => {
      // console.log("initStyle", attrs.process)
      submitState.value.style_width = attrs.process.style_width
      submitState.value.style_height = attrs.process.style_height
      submitState.value.style_color = attrs.process.style_color
      submitState.value.style_icon = attrs.process.icon
    }
    // #endregion
    return {
      activeKey,
      submitState,
      dataSource,
      columns,
      MyIcons,
      formState,
      onSubmit,
      tmpNextProcess,
      gridOptions,
      tmpBeixuanProcess,
      removePrs,
      addPrs,
      open,
      depts,
      xGrid,
      toggleAllCheckboxEvent,
      selectRecords,
      toggleCheckboxEvent,
      gridEvent,
      selectedEmp,
      selPer,
      selDep,
      tabcolumns,
      state,
      onSelectChange,


      fields,
      nextProcesses,
      bindExprs,
      stateExprs,

      transText,
      addCondi,
      validateExpr,
      getCurrentCond,
      delExpr,

      colors,
      onMyIcon,
      setColor,


      disableAuto,
      changeAuto


    };
  }
};

</script>

<style lang="scss">
.show-item {
  width: 140px;
  height: 240px;
  border-bottom-style: 2px solid #d9d9d9;
  box-sizing: border-box;
  border-radius: 4px;
  padding: 4px 11px;
  margin: 0 8px;
  font-size: 18px;
  color: #1478FF;
  line-height: 22px;
  background-color: #fafafa;
  display: inline-block;
  position: relative;
  margin: 4px;
}

.show-item::before {
  position: absolute;
  content: '';
  width: 6px;
  height: 100%;
  background: #1478FF;
  position: absolute;
  left: 0px;
  top: 50%;
  transform: translateY(-50%);
}

.show-expr {
  width: 200px;
  height: 64px;
  box-sizing: border-box;
  border-radius: 4px;
  padding: 4px 11px;
  margin: 0 8px;
  font-size: 14px;
  color: #20C06B;
  line-height: 22px;
  background-color: #fafafa;
  display: inline-block;
  position: relative;
  margin: 4px;
}

.show-expr::before {
  position: absolute;
  content: '';
  width: 6px;
  height: 100%;
  background: #20C06B;
  position: absolute;
  left: 0px;
  top: 50%;
  transform: translateY(-50%);
}

.expr {
  margin: 2px;
  width: 96%;
  padding: 0 6px;
  height: 44px;
  line-height: 44px;
  border-bottom: 1px solid #f8cb8b;
  border-top: 1px solid #f8cb8b;
  border-right: 1px solid #f8cb8b;
  box-sizing: border-box;
  color: #FC933C;
  background-color: #fafafa;
  display: inline-block;
  position: relative;
}

.expr::before {
  position: absolute;
  content: '';
  width: 6px;
  height: 100%;
  background: #FC933C;
  box-sizing: border-box;
  position: absolute;
  left: 0px;
  top: 50%;

  transform: translateY(-50%);
}
</style>