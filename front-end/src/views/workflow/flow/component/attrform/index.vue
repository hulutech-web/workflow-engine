<template>
  <div>
    <div>
      <n-tabs v-model:activeKey="activeKey" type="card">
        <!-- 常规设置 -->
        <n-tab-pane name="1" tab="常规" style="height: 240px" v-if="formState.process">
          <n-form-item label="步骤名称">
            <n-input v-model:value="submitState.process_name" :value="formState.process.process_name"></n-input>
          </n-form-item>
          <n-form-item label="步骤类型">
            <n-radio-group v-model:value="submitState.process_position" :value="formState.process.position">
              <n-radio :value="1">正常步骤</n-radio>
              <n-radio v-if="formState.can_child" :value="2">转入子流程</n-radio>
              <!-- <n-radio  :value="2">转入子流程</n-radio> -->
              <n-radio :value="0">第一步</n-radio>
            </n-radio-group>
          </n-form-item>
          <n-divider></n-divider>
          <!--          非转入子流程-->
          <div
            v-if="formState.next_process && (submitState.process_position == 1 || submitState.process_position == 0)">
            <div class="flex">
              <div class="px-3 flex flex-col item-center justify-center">
                -->
              </div>
              <div class="flex-1  border-blue-600 border border-solid  p-2">
                <p class="text-md">下一步骤</p>
                <n-row>
                  <n-col :span="16" v-for="(p, index) in formState.next_process" :key="index">
                    <n-tag :bordered="false" :color="{color:'geekblue'}" v-if="p.NextProcess&&p.NextProcess.id != -1">
                      {{ p.NextProcess.process_name }}
                    </n-tag>
                  </n-col>
                </n-row>
              </div>

              <!-- <div class="flex-1  border-blue-500 border border-dotted  p-2">
                <p class="text-md">其他步骤</p>
                <n-row>
                  <n-col :span="16" v-for="(p, index) in formState.beixuan_process" :key="index">
                    <n-tag :bordered="false" color="geekblue" v-if="p.NextProcess.id != -1">
                      {{ p.Process.process_name }}
                    </n-tag>
                  </n-col>
                </n-row>
              </div> -->
            </div>

          </div>

          <!--          转入子流程-->

          <div v-if="formState.next_process">
            <div id="child_flow" v-if="formState.process.position == 2">
              <div class="control-group">
                <n-form-item label="子流程">
                  <n-select v-model:value="submitState.child_flow_id" :default-value="formState.process.child_flow_id">
                    <!--                    <n-select-option value="0">请选择</n-select-option>-->

                    <!--                    <n-select-option v-for="(flow, ind) in formState.flows" :key="ind" :value="flow.id"-->
                    <!--                      :selected="formState.process.child_flow_id == flow.id">-->
                    <!--                      {{ flow.flow_name }}-->
                    <!--                    </n-select-option>-->
                  </n-select>
                </n-form-item>

              </div>

              <div class="control-group">
                <n-form-item label="子流程结束后动作">
                  <n-radio-group v-model:value="submitState.child_after" :default-value="formState.process.child_after">
                    <n-radio :value="1">
                      同时结束父流程
                    </n-radio>
                    <n-radio :value="2">
                      返回父流程步骤
                    </n-radio>
                  </n-radio-group>
                </n-form-item>

              </div>

              <div v-if="submitState.child_after == 2">
                <n-form-item label="返回步骤">
                  <n-select name="child_back_process" v-model:value="submitState.child_back_process"
                            :default-value="formState.child_back_process">
                    <!--                    <n-select-option value="0">-->
                    <!--                      无-->
                    <!--                    </n-select-option>-->
                    <!--                    <n-select-option v-for="(p, index) in formState.processes" :key="index" :value="p.id"-->
                    <!--                      :selected="p.child_back_process == p.id">-->
                    <!--                      {{ p.process_name }}-->
                    <!--                    </n-select-option>-->
                  </n-select>
                </n-form-item>
                <span class="help-inline">默认为当前步骤下一步</span>
              </div>
            </div>
          </div>
        </n-tab-pane>


        <n-tab-pane name="2" tab="表单" style="height: 240px">
          <n-data-table bordered :columns="columns" :data="dataSource"></n-data-table>
        </n-tab-pane>


        <n-tab-pane name="3" tab="权限" style="height: 240px">
          <div>
            <div>
              <n-form-item label="自动选人">
                <n-select v-model:value="submitState.auto_person" @update:value="changeAuto"
                          :options="options.auto_person">
                </n-select>
              </n-form-item>
            </div>
            <n-divider>
              授权范围（适用于：当需要手动选人时，则授权范围生效）
            </n-divider>

            <div>
              <span>授权人员：</span>
              <div class="flex">
                <n-select v-model:value="submitState.range_emp_ids" :disabled="disableAuto" multiple
                          placeholder="选择人员" style="width:400px;">
                </n-select>
                <n-button class="ml-3" type="primary" @click="selPer" :disabled="disableAuto">选择</n-button>
              </div>
            </div>

            <div class="mt-3">
              <span>授权部门：</span>
              <div class="flex">

                <n-select v-model:value="submitState.range_dept_ids" :disabled="disableAuto" multiple
                          placeholder="选择部门" style="width:400px;">
                  <!--                <n-select-option v-for="(i,ind) in submitState.range_dept_ids" :key="ind" :value="i">-->
                  <!--                  {{ submitState.range_dept_text[ind] }}-->
                  <!--                </n-select-option>-->
                </n-select>
                <n-button class="ml-3" :disabled="disableAuto" type="primary" @click="selDep">选择</n-button>
              </div>

            </div>

            <n-modal v-model:show="open">
              <n-card title="人员&部门选择" style="width: 700px" size="small">
                <span v-if="selectedEmp == false" class="text-sm text-orange-500">部门管理员为审批人</span>
                <vxe-grid ref='xDeptGrid' v-bind="gridDeptOptions" v-on="gridDeptEvent" v-if="selectedEmp == false">
                  <template #action="{ row }">
                    <div>
                      选择
                    </div>
                  </template>
                </vxe-grid>

                <vxe-grid ref='xGrid' v-bind="gridOptions" v-on="gridEvent" v-if="selectedEmp == true">
                  <template #checkbox_header="{ checked, indeterminate }">
                    <div>选择</div>
                  </template>
                  <template #checkbox_cell="{ row, checked, indeterminate }">
                  <span class="custom-checkbox" @click.stop="toggleCheckboxEvent(row)">
                    <n-checkbox v-if="indeterminate" :checked="checked"></n-checkbox>
                    <n-checkbox v-else-if="checked" :checked="checked"></n-checkbox>
                    <n-checkbox v-else></n-checkbox>
                  </span>
                  </template>
                  <template #dept="{ row }">
                    <div>
                      {{ row.Dept.id == 0 ? "未分配" : row.Dept.dept_name }}
                    </div>
                  </template>
                </vxe-grid>

              </n-card>
            </n-modal>


          </div>
        </n-tab-pane>


        <n-tab-pane name="4" tab="转出条件">
          <div class="condi-container">
            <n-row>
              <n-col :span="4">
                <div class="text-md font-bold">
                  转出步骤
                </div>
              </n-col>
              <n-col :span="6">
                <div class="text-md font-bold">
                  转出条件
                </div>
              </n-col>
              <n-col :span="14">
                <div class="text-md font-bold">
                  更改规则
                  <span class="text-orange-500">注意：填写完规则后请完成校验！！！</span>
                </div>
              </n-col>
            </n-row>
            <div style="height:10px;"></div>
            <n-row v-for="(item, index) in formState.next_process" :key="index"
                   v-if="formState.next_process.length > 1">
              <n-col :span="4">
                <div class="show-item">
                  {{ item.NextProcess.process_name }}
                </div>
              </n-col>
              <n-col :span="6">
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
              </n-col>
              <n-col :span="14" style="padding:6px;background-color:#fafafa;border-bottom:1px solid orange">
                <n-row>
                  <n-col :span="4">
                    <div class="text-center">字段</div>
                    <n-select style="width: 100%;" v-model:value="bindExprs[index]['field']" size="small"
                              :options="fields.map(item=>{
                                return {
                                  label: item.field_name,
                                  value: item.field
                                }
                              })">
                    </n-select>
                  </n-col>
                  <n-col :span="4">
                    <div class="text-center">条件</div>
                    <n-select style="width: 100%;" v-model:value="bindExprs[index]['operator']" size="small"
                              :options="selectOpt.condi_opts">

                    </n-select>
                  </n-col>
                  <n-col :span="4">
                    <div class="text-center">值</div>
                    <n-input v-model:value="bindExprs[index]['value']" size="small"></n-input>
                  </n-col>
                  <n-col :span="4">
                    <div class="text-center">其他条件</div>
                    <n-select style="width: 100%;" v-model:value="bindExprs[index]['extra']" size="small"
                              :options="selectOpt.extra_opts">

                    </n-select>
                  </n-col>

                  <n-col :span="4">
                    <div>
                      <div class="text-center">操作</div>
                    </div>
                    <div class="text-center">
                      <n-button-group>
                        <n-button type="primary" @click="addCondi(index)" size="small">新增</n-button>
                      </n-button-group>
                    </div>
                  </n-col>
                  <n-col :span="4">
                    <div class="text-center">确认条件</div>
                    <div class="text-center">
                      <n-button type="primary" @click="validateExpr(index)" size="small">确认</n-button>
                    </div>
                  </n-col>
                </n-row>
                <template v-for="(sE, ind) in stateExprs[index]" :key="ind">
                  <div class="expr" v-if="index == sE['index']">
                    <n-row>
                      <n-col :span="4">
                        <div class="text-center"> {{ sE['field'] }}</div>
                      </n-col>
                      <n-col :span="4">
                        <div class="text-center">{{ sE['operator'] }}</div>
                      </n-col>
                      <n-col :span="4">

                        <div class="text-center">{{ sE['value'] }}</div>
                      </n-col>
                      <n-col :span="4">
                        <div class="text-center">{{ sE['extra'] }}</div>
                      </n-col>
                      <n-col :span="4">
                        <n-space>
                          <n-button type="error" class="cursor-pointer ml-4" size="small" @click="delExpr(index, ind)">
                            X
                          </n-button>
                        </n-space>
                      </n-col>
                    </n-row>
                  </div>
                </template>

              </n-col>
            </n-row>


          </div>

        </n-tab-pane>


        <n-tab-pane name="5" tab="样式" style="height: 240px">
          <div class="p-3">
            <div class="flex justify-start items-center mt-3">
              <div class="flex-4" style="width:80px;">尺寸</div>
              <div class="flex-8 flex">
                <n-space>
                  <n-input-number v-model:value="submitState.style_width"></n-input-number>
                  X
                  <n-input-number v-model:value="submitState.style_height"></n-input-number>
                </n-space>
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
                          class="cursor-pointer bg-black text-white  rounded w-6"/>
              </div>
              <input type="text" class="h-8 border-none outline-none bg-gray-100 rounded-sm px-3 flex-2"
                     v-model="submitState.style_icon">

              <div style="width:600px;background-color: black;line-height:24px;" class="ml-4 flex flex-wrap ">
                <HuluIcon @click="onMyIcon(ic)" fontSize="24px" :name="ic" v-for="(ic, index) in MyIcons" :key="index"
                          class="m-3 cursor-pointer hover:scale-125"/>
              </div>
            </div>
          </div>
        </n-tab-pane>

      </n-tabs>

      <div class="absolute bottom-0 left-0 ml-5 mb-5">
        <n-button type="primary" @click="onSubmit">
          提交
        </n-button>
      </div>
    </div>

  </div>

</template>

<script lang='ts'>

import {icons} from './icon'
import useEmpconfig from './empconfig'
import {useMessage} from "naive-ui";
import {ExplainConditionSql} from "./sql/explain";
import useDeptConfig from "./deptconfig"

const {getCurrCond} = useProcess()
const {gridOptions} = useEmpconfig()
const {gridDeptOptions} = useDeptConfig();
const {loadDepts} = useDept()


export default {
  props: ['attrs'],
  emits: ["updProcess"],
  setup(props, context) {
    // #region 常规
    const message = useMessage();
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
      selectOpt.value.flows = attrs.process.child_flow_id
      selectOpt.value.processes = attrs.process.child_after
      selectOpt.value.back_processes = attrs.process.child_back_process
    }

    const fillSubmitState = (attrs) => {
      submitState.value.process_name = attrs.process.process_name
      submitState.value.process_position = attrs.process.position
      submitState.value.process_to = attrs.next_process.map(item => item.NextProcess ? item.NextProcess.id : -1)
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
      // console.log(submitState.value)
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
    const changeAuto = (value: string, option: SelectOption) => {
      if (value == '-1001') {
        disableAuto.value = true
      }
      if (value == '-1002') {
        disableAuto.value = true
      }
      if (value == "0") {
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


    const gridDeptEvent: VxeGridListeners<RowVO> = {
      proxyQuery() {
        console.log('数据代理查询事件')
        const grid = xDeptGrid.value
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
    }

    const options = ref({
      auto_person: [
        {
          label: '发起人部门经理',
          value: '-1001'
        },
        {
          label: '发起人部门主管',
          value: '-1002'
        },
        {
          label: '手动选择',
          value: '0'
        }


      ]
    })


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
      submitState.value.range_dept_text = selectedRows.map(item => item.dept_name)
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
        return message.warning("请填写完整")
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
        message.warning(msg)
      } else {
        message.success("条件校验成功")
        submitState.value.process_condition = stateExprs.value[0]
      }
    }
    const xDeptGrid = ref()

    const getCurrentCond = async (process) => {
      let param = {
        flow_id: process.flow_id,
        process_id: process.process_id,
        next_process_id: process.next_process_id
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

    //根据不同的输入框，给定不同的选项类型
    const changeSelectDisable = (field_type, typename) => {
      switch (field_type) {
        case "textarea":
          // 只保留包含于不包含，其他disable
          break;
        case "file":
          // 直接禁用输入框
          break;
        case "number":
          // 只保留大于，大于等于，小于小于等于，等于，不等于，其他disable
          break;
        case "select":
          // 渲染select  选项框，并将field_value作为选项
          break;
        case "date":
          // 直接禁用输入框
          break;
        case "radio":
          // 渲染select  选项框，并将field_value作为选项
          break;
        case "checkbox":
          // 直接禁用输入框
          break;
        default:
          // 渲染包含和不包含，其他禁用
          break;
      }
    }

    const selectOpt = ref({
      flows: [],
      processes: [],
      range_emp_ids: [],
      range_dept_ids: [],
      condi_opts: [
        {
          label: '等于',
          value: '='
        },
        {
          label: '不等于',
          value: '!='
        },
        {
          label: '大于',
          value: '>'
        },
        {
          label: '小于',
          value: '<'
        },
        {
          label: '大于等于',
          value: '>='
        },
        {
          label: '小于等于',
          value: '<='
        },
        {
          label: '包含',
          value: 'in'
        },
        {
          label: '不包含',
          value: 'not in'
        },
      ],
      extra_opts: [
        {
          label: '并且',
          value: 'AND'
        },
        {
          label: '或者',
          value: 'OR',
        }
      ]
    })
    // #endregion
    return {
      activeKey,
      submitState,
      selectOpt,
      dataSource,
      columns,
      MyIcons,
      formState,
      onSubmit,
      tmpNextProcess,
      gridDeptOptions,
      gridOptions,
      tmpBeixuanProcess,
      removePrs,
      addPrs,
      open,
      depts,
      xGrid,
      xDeptGrid,
      toggleAllCheckboxEvent,
      selectRecords,
      toggleCheckboxEvent,
      gridEvent,
      gridDeptEvent,
      selectedEmp,
      selPer,
      selDep,
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
      changeAuto,

      options,
    };
  }
};

</script>

<style lang="scss">
.show-item {
  width: 140px;
  height: 140px;
  border-bottom-style: 2px solid #d9d9d9;
  box-sizing: border-box;
  border-radius: 4px;
  padding: 0px 11px;
  //font-size: 16px;
  font-weight: bolder;
  color: #1478FF;
  line-height: 22px;
  background-color: #fafafa;
  display: inline-block;
  position: relative;
  margin: 0 4px;
}

.show-item::before {
  position: absolute;
  content: '';
  width: 2px;
  height: 100%;
  background: #1478FF;
  left: 0px;
  top: 50%;
  transform: translateY(-50%);
}

.show-item::after {
  position: absolute;
  content: '';
  //  设置底部边框
  border-bottom: 1px solid #ddd;
  left: 0;
  right: 0;
  bottom: 0;
  margin: 0 auto;
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
}

.show-expr::before {
  position: absolute;
  content: '';
  width: 2px;
  height: 100%;
  background: #20C06B;
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

.condi-container {
  max-height: 700px;
  overflow-y: scroll;
  scrollbar-width: none; /* Firefox */
  -ms-overflow-style: none; /* IE/Edge */
}
</style>
