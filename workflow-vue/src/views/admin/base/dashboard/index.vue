<template>
   <div>
      <a-row :gutter="[32, 8]">
         <a-col :span="12">
            <a-card title="管理面板">
               <div class="flex">
                  <div @click="toUser" class="mx-3 border border-gray-500 bg-teal-500 p-4 w-64 h-32
               rounded-md hover:scale-105 transition-all duration-300 cursor-pointer">
                     <p class="text-lg text-white font-bold border-b-2 border-gray-900">用户管理</p>
                     <span class="text-gray-700">用户信息，绑定员工</span>
                  </div>
                  <div @click="toEmp" class="mx-3 border border-gray-500 bg-green-400 p-4 w-64 h-32
               rounded-md hover:scale-105 transition-all duration-300 cursor-pointer">
                     <p class="text-lg text-white font-bold border-b-2 border-gray-900">员工管理</p>
                     <span class="text-gray-700">员工信息，部门</span>
                  </div>
                  <div @click="toDept" class="mx-3 border border-gray-500 bg-orange-400 p-4 w-64 h-32
               rounded-md hover:scale-105 transition-all duration-300 cursor-pointer">
                     <p class="text-lg text-white font-bold border-b-2 border-gray-900">部门管理</p>
                     <span class="text-gray-700">部门管理、部门经理、部门主管</span>
                  </div>
                  <div @click="toFlow" class="mx-3 border border-gray-500 bg-red-400 p-4 w-64 h-32
               rounded-md hover:scale-105 transition-all duration-300 cursor-pointer">
                     <p class="text-lg text-white font-bold border-b-2 border-gray-900">流程管理</p>
                     <span class="text-gray-700">流程管理、低代码流程图</span>
                  </div>
                  <div @click="toTemplate" class="mx-3 border border-gray-500 bg-purple-600 p-4 w-64 h-32
               rounded-md hover:scale-105 transition-all duration-300 cursor-pointer">
                     <p class="text-lg text-white font-bold border-b-2 border-gray-900">模板管理</p>
                     <span class="text-gray-700">模板管理、表单控件</span>
                  </div>
               </div>
            </a-card>
         </a-col>
         <a-col :span="12">
            <a-card title="工作流">
               <div class="flex">
                  <div  v-for="(item, index) in state.flows" :key="index" @click="toWorkflow(item)">
                  <div @click="" class="mx-3 border border-gray-500 bg-green-400 p-4 w-64 h-32
               rounded-md hover:scale-105 transition-all duration-300 cursor-pointer">
                     <p class="text-lg text-white font-bold border-b-2 border-gray-900">{{ item.flow_name }}</p>
                     <span class="text-gray-700">编号：{{ item.flow_no }}</span>
                  </div>
               </div>
               </div>
               

            </a-card>
         </a-col>
      </a-row>

      <a-row class="mt-4" :gutter="[32, 8]">

         <a-col :span="12">
            <p class="mx-3 text-lg">审批流程</p>
            <div class="px-3">
               <a-tabs v-model:activeKey="activeKey">
                  <a-tab-pane key="1" tab="我的申请" force-render>
                     <a-card title="我的申请" :headStyle="{ backgroundColor: '#4b7bec', color: 'white' }">
                        <div style="height:400px;overflow-y: auto;" class="bg-blue-100">
                           <vxe-table border show-overflow ref="tableRef" max-height="400"
                              :column-config="{ resizable: true }" :data="state.entries">
                              <vxe-column field="title" width="240" title="标题"
                                 :edit-render="{ name: 'input' }"></vxe-column>
                              <vxe-column field="user" title="发起人" width="180" :edit-render="{ name: 'input' }"
                                 #default="{ row }">
                                 {{ row.Emp.name }}
                              </vxe-column>
                              <vxe-column field="process" title="当前位置" width="80" :edit-render="{ name: 'input' }"
                                 #default="{ row }">
                                 {{ row.Process.process_name }}
                              </vxe-column>
                              <vxe-column field="status" title="当前状态" width="80" :edit-render="{ name: 'input' }"
                                 #default="{ row }">
                                 <a-badge v-if="row.status == 0" status="processing" text="处理中" />
                                 <a-badge v-if="row.status == 9" status="success" text="通过" />
                                 <a-badge v-if="row.status == -1" status="error" text="驳回" />
                                 <a-badge v-if="row.status == -2" status="error" text="撤销" />
                                 <a-badge v-if="row.status == -9" status="warning" text="草稿" />
                              </vxe-column>
                              <vxe-column field="action" title="操作" :edit-render="{ name: 'input' }" #default="{ row }">
                                 <div>
                                    <a-button-group>
                                       <a-button v-if="row.status == -1" type="primary">编辑</a-button>
                                       <a-button v-if="row.status == -1" type="link" @click="resend(row)">重发</a-button>
                                       <a-button v-if="row.status == -1" type="dashed" danger>撤销</a-button>
                                       <a-button type="primary" dashed @click="toEntryData(row)">查看</a-button>
                                       <a-button type="primary" @click="toHistory(row)">进程</a-button>
                                    </a-button-group>
                                 </div>
                              </vxe-column>
                           </vxe-table>
                        </div>
                     </a-card>
                  </a-tab-pane>
                  <a-tab-pane key="2" tab="待办事项">
                     <a-card title="待办事项" :headStyle="{ backgroundColor: '#cc8e35', color: 'white' }">
                        <div style="height:400px;overflow-y: auto;" class="bg-orange-100">
                           <vxe-table border show-overflow ref="tableRef" max-height="400"
                              :column-config="{ resizable: true }" :data="state.procs">
                              <vxe-column field="title" title="标题" :edit-render="{ name: 'input' }" #default="{ row }">
                                 {{ row.Entry.title }}
                              </vxe-column>
                              <vxe-column field="user" title="发起人" :edit-render="{ name: 'input' }" #default="{ row }">
                                 {{ row.Entry.Emp.name }}
                              </vxe-column>
                              <vxe-column field="process" title="流程位置" :edit-render="{ name: 'input' }"
                                 #default="{ row }">
                                 {{ row.process_name }}
                              </vxe-column>
                              <vxe-column field="status" title="流程" :edit-render="{ name: 'input' }" #default="{ row }">
                                 {{ row.Flow.flow_name }}
                              </vxe-column>
                              <vxe-column field="auditor" title="审核人" :edit-render="{ name: 'input' }"
                                 #default="{ row }">
                                 {{ row.emp_name }}
                              </vxe-column>
                              <vxe-column field="action" title="操作" :edit-render="{ name: 'input' }" #default="{ row }"
                                 width="240">
                                 <a-button-group>
                                    <a-button type="primary" v-if="row.status == 0" @click="toProc(row)">批复</a-button>
                                    <a-button type="link" v-if="row.status == 9">已处理(通过)</a-button>
                                    <a-button type="primary" v-if="row.status == -1">已处理(驳回)</a-button>
                                    <a-button type="primary" ghost @click="proc_detail(row)">进程明细</a-button>
                                 </a-button-group>
                              </vxe-column>
                           </vxe-table>

                        </div>
                     </a-card>
                  </a-tab-pane>

                  <a-tab-pane key="3" tab="已处理事项">
                     <a-card title="已处理事项" :headStyle="{ backgroundColor: '#20bf6b', color: 'white' }">
                        <div style="height:400px;overflow-y: auto;" class="bg-green-100">

                           <vxe-table border show-overflow ref="tableRef" max-height="400"
                              :column-config="{ resizable: true }" :data="state.handle_procs">
                              <vxe-column field="id" title="id" width="80"
                                 :edit-render="{ name: 'input' }"></vxe-column>
                              <vxe-column field="title" title="标题" width="80" :edit-render="{ name: 'input' }"
                                 #default="{ row }">
                                 <span>{{ row.Entry.title }}</span>
                              </vxe-column>
                              <vxe-column field="emp" title="发起人" width="80" :edit-render="{ name: 'input' }"
                                 #default="{ row }">
                                 <span>{{ row.Emp.name }}</span>
                              </vxe-column>
                              <vxe-column field="process_name" title="环节" :edit-render="{ name: 'input' }"></vxe-column>
                              <vxe-column field="emp_name" title="当前审批人" :edit-render="{ name: 'input' }"></vxe-column>
                              <vxe-column field="content" title="意见" :edit-render="{ name: 'input' }"></vxe-column>
                              <vxe-column field="status" title="状态" :edit-render="{ name: 'input' }" #default="{ row }">
                                 <a-badge v-if="row.status == 0" status="processing" text="处理中" />
                                 <a-badge v-if="row.status == 9" status="success" text="通过" />
                                 <a-badge v-if="row.status == -1" status="error" text="驳回" />
                                 <a-badge v-if="row.status == -2" status="error" text="撤销" />
                                 <a-badge v-if="row.status == -9" status="warning" text="草稿" />
                              </vxe-column>
                              <vxe-column field="dept_name" title="部门" :edit-render="{ name: 'input' }"></vxe-column>
                              <vxe-column field="detail" title="明细" :edit-render="{ name: 'input' }" #default="{ row }">
                                 <a-button type="primary" ghost @click="proc_detail(row)">进程明细</a-button>
                              </vxe-column>
                           </vxe-table>
                        </div>

                     </a-card>
                  </a-tab-pane>
               </a-tabs>
            </div>
            <a-modal :footer="false" v-model:open="open" width="1000px" title="审批进程" centered
               :bodyStyle="{ height: '1000px' }">
               <vxe-table border show-overflow ref="tableRef" max-height="400" :column-config="{ resizable: true }"
                  :data="modalState.procs">
                  <vxe-column field="id" title="id" width="80" :edit-render="{ name: 'input' }"></vxe-column>
                  <vxe-column field="title" title="标题" width="80" :edit-render="{ name: 'input' }" #default="{ row }">
                     <span>{{ row.Entry.title }}</span>
                  </vxe-column>
                  <vxe-column field="emp" title="发起人" width="80" :edit-render="{ name: 'input' }" #default="{ row }">
                     <span>{{ row.Entry.Emp.name }}</span>
                  </vxe-column>
                  <vxe-column field="process_name" title="环节" :edit-render="{ name: 'input' }"></vxe-column>
                  <vxe-column field="emp_name" title="当前审批人" :edit-render="{ name: 'input' }"></vxe-column>
                  <vxe-column field="content" title="意见" :edit-render="{ name: 'input' }"></vxe-column>
                  <vxe-column field="status" title="状态" :edit-render="{ name: 'input' }" #default="{ row }">
                     <a-badge v-if="row.status == 0" status="processing" text="处理中" />
                     <a-badge v-if="row.status == 9" status="success" text="通过" />
                     <a-badge v-if="row.status == -1" status="error" text="驳回" />
                     <a-badge v-if="row.status == -2" status="error" text="撤销" />
                     <a-badge v-if="row.status == -9" status="warning" text="草稿" />
                  </vxe-column>
                  <vxe-column field="dept_name" title="部门" :edit-render="{ name: 'input' }"></vxe-column>
               </vxe-table>
               <div class="h-96 overflow-y-auto box-border m-3 w-full">
                  <p class="text-xl font-bold mb-3">
                     时间线
                  </p>
                  <div class="w-[90%]">
                     <a-timeline mode="alternate">
                        <a-timeline-item v-for="(proc, index) in modalState.procs" :key="index"
                           :color="showState(proc.status)">
                           {{ proc.emp_name }} {{ proc.created_at }}
                        </a-timeline-item>
                     </a-timeline>
                  </div>

               </div>
               <div class="text-center mt-12">
                  <a-button type="primary" class="w-[50%]" @click="open = false">关闭</a-button>
               </div>

            </a-modal>

         </a-col>
      </a-row>
   </div>
</template>
<script lang="ts" setup>

// 
const router = useRouter();
const { index } = useHome()
const { indexProcs } = useProc();
const { resendEntry } = useEntry()
const showState = (state) => {
   switch (state) {
      case 9:
         return 'green'
         break;
      case -1:
         return 'red'
         break;
      case -2:
         return 'gray'
         break;
      case -9:
         return 'orange'
         break;
      default:
         return 'blue'
         break;
   }
}
const state = ref({
   entries: [],
   flows: [],
   handle_procs: [],
   procs: []
})
const modalState = ref({
   procs: []
})
const activeKey = ref('1')
const formatState = (val) => {
   switch (val) {
      case 0:
         return '处理中'
      case 9:
         return '通过'
      case -1:
         return '驳回'
      case -2:
         return '撤销'
      case -9:
         return '草稿'
      default:
         return '未知'
   }
}
const init = async () => {
   const { data } = await index()
   state.value.entries = data.entries
   state.value.flows = data.flows
   state.value.handle_procs = data.handle_procs
   state.value.procs = data.procs
}
const open = ref(false)
onMounted(() => {
   init()
})
const toEmp = () => {
   router.push({ path: "/admin/base/user/emp" })
}
const toDept = () => {
   router.push({ path: "/admin/base/dept/index" })
}

const toFlow = () => {
   router.push({ path: "/admin/base/flow/index" })
}
const toUser = () => {
   router.push({ path: "/admin/base/user/index" })
}
const toTemplate = () => {
   router.push({ path: "/admin/base/template/index" })
}

const toEntryData = (row) => {
   router.push({ path: `/admin/base/flow/${row.flow_id}/entry/${row.id}` })
}
const toWorkflow = (row) => {
   // 初始化流程
   router.push({ path: `/admin/base/flow/${row.id}/initiation` })

}
const toProc = (row) => {
   router.push({ path: `/admin/base/flow/${row.flow_id}/proc/${row.entry_id}`, query: { process_id: row.process_id, proc_id: row.id } })
}

const toPlugin = ()=>{
   router.push({ path: "/admin/base/plugin/index" })
}

const toHistory = async (row) => {
   open.value = true
   console.log(row.id)
   const { data } = await indexProcs(row.id)
   modalState.value.procs = data
}

const resend = async (row) => {
   await resendEntry(row.id)
   await init();
}
const proc_detail = async (row) => {
   open.value = true
   console.log(row.id)
   const { data } = await indexProcs(row.entry_id)
   modalState.value.procs = data
}
</script>

<style scoped></style>