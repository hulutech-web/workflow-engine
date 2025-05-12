<template>
  <div class="p-3">
    <p>设计插件</p>
    <a-form style="margin-top: 20px" ref="formRef" :model="formState">
      <a-row :gutter="[16, 32]">
        <a-col :span="8">
          <a-form-item label="步骤" name="process_id">
            <a-select
              @select="selectProc"
              v-model:value="formState.process_id"
              placeholder="请选择步骤"
            >
              <a-select-option
                v-for="(opt, index) in process_opts"
                :key="index"
                :value="opt.id"
              >
                {{ opt.process_name }}
              </a-select-option>
            </a-select>
          </a-form-item>
          <a-form-item label="表单字段" name="field_id">
            <a-select
              v-model:value="formState.field_id"
              placeholder="请选择步骤"
            >
              <a-select-option
                v-for="(opt, index) in templateForms"
                :key="index"
                :value="opt.id"
              >
                {{ opt.field_name }}
              </a-select-option>
            </a-select>
          </a-form-item>
          <a-form-item label="分配规则" name="rule">
            <div class="m-2">
              <a-button type="dashed" @click="addItem" class="w-96">
                <PlusOutlined />
                添加规则
              </a-button>
            </div>
            <a-col
              :span="8"
              v-for="(rule, index) in formState.rules"
              :key="index"
            >
              <div
                class="flex bg-gray-100 p-2 justify-between items-center"
                style="width: 660px; margin-left: 10px"
              >
                <a-form-item
                  name="rule_value"
                  label="部门"
                  style="width: 200px; margin-left: 10px"
                >
                  <a-select
                    v-model:value="rule.rule_id"
                    @change="handleDeptChange"
                    @select="handleSelect(index)"
                  >
                    <a-select-option
                      v-for="(dept, index) in depts"
                      :key="index"
                      :value="dept.id"
                    >
                      {{ dept.dept_name }}
                    </a-select-option>
                  </a-select>
                </a-form-item>
                <a-form-item
                  name="rule_label"
                  label="部门名"
                  style="width: 200px; margin-left: 10px"
                >
                  <a-input
                    v-model:value="rule.rule_label"
                    :disabled="true"
                  ></a-input>
                </a-form-item>
                <a-form-item
                  name="value"
                  label="数量"
                  style="width: 200px; margin-left: 10px"
                >
                  <a-input-number
                    v-model:value="rule.rule_value"
                  ></a-input-number>
                </a-form-item>
                <a-form-item style="width: 200px">
                  <MinusCircleOutlined
                    v-if="formState.rules.length > 1"
                    style="font-size: 18px; color: red"
                    class="dynamic-delete-button"
                    @click="removeItem(rule)"
                  />
                </a-form-item>
              </div>
            </a-col>
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <p>流程字段</p>
          <vxe-table
            border
            show-overflow
            ref="tableRef"
            max-height="400"
            :column-config="{ resizable: true }"
            :data="templateForms"
          >
            <vxe-column field="field" title="名称"> </vxe-column>
            <vxe-column field="field_name" title="标题"> </vxe-column>
          </vxe-table>
        </a-col>
        <a-col :span="8">
          <p>所有规则</p>
          <vxe-table
            border
            show-overflow
            max-height="400"
            :column-config="{ resizable: true }"
            :data="allConfig"
          >
            <vxe-column field="flow_id" title="流程" #default="{ row }">
              {{ row.Flow.flow_name }}
            </vxe-column>
            <vxe-column field="process_id" title="步骤" #default="{ row }">
              {{ row.Process.process_name }}
            </vxe-column>
            <vxe-column field="field_id" title="字段" #default="{ row }">
              {{ row.TemplateForm.field_name }}
            </vxe-column>
            
            <vxe-column
              field="config"
              title="具体规则"
              #default="{ row }"
              height="400"
            >
              <span v-for="(item, index) in row.config" :key="index"
                >{{ item.rule_label }}:{{ item.rule_value }}</span
              >
              <!-- <a-collapse >
                <a-collapse-panel v-for="(item,index) in row.config" :key="index" :header="item.rule_label">
                  <p>{{ item.rule_value }}</p>
                </a-collapse-panel>
              </a-collapse> -->
            </vxe-column>
          </vxe-table>
        </a-col>
      </a-row>
      <a-form-item>
        <a-button type="primary" @click="onSubmit">提交</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
const rulesStore = useRulesStore();
const { getFlowProcesses } = useProcess();
const { storePluginConfig, getAllPluginConfig } =
  usePluginConfig();
const { loadFlowTemplateform } = useFlow();
const { loadDeptList } = useDept();
const route = useRoute();
const flow_id = ref(+route.query.flow_id);
const plugin_id = ref(+route.query.plugin_id);

const process_opts = ref([]);
const depts = ref([]);

const templateForms = ref([]);
const currentSelectIndex = ref(0);
const allConfig = ref([]);
const init = async () => {
  const { data } = await getFlowProcesses(flow_id.value);
  process_opts.value = data;
  const { data: deptData } = await loadDeptList();
  depts.value = deptData;
  const { data: formData } = await loadFlowTemplateform(flow_id.value);
  templateForms.value = formData;
  //加载现有所有规则
  const { data: configs } = await getAllPluginConfig(formState.value);
  allConfig.value = configs;
  console.log(allConfig.value);
};
init();
interface RuleItem {
  rule_id: number;
  rule_label: string;
  rule_value: number;
}
const handleDeptChange = (index) => {
  console.log(index);
  if (formState.value.rules[currentSelectIndex.value]) {
    formState.value.rules[currentSelectIndex.value].rule_label =
      depts.value[index - 1]["dept_name"];
  }
};
const handleSelect = (v) => {
  currentSelectIndex.value = v;
};
const formState = ref({
  flow_id: flow_id.value,
  process_id: null,
  plugin_id: plugin_id.value,
  rules: [] as RuleItem[],
});
const onSubmit = async () => {
  console.log(formState.value);
  await storePluginConfig(formState.value);
};
const removeItem = (val) => {
  const index = formState.value.rule.indexOf(val);
  if (index !== -1) {
    formState.value.rule.splice(index, 1);
  }
};
const addItem = () => {
  formState.value.rules.push({
    rule_id: 0,
    rule_label: "",
    rule_value: 0,
  });
};

const selectProc = async (val) => {
  formState.value.process_id = val;
};
</script>

<style></style>
