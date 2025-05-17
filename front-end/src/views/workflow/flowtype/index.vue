<template>
    <div>
        <n-grid :gutter="[24, 1]">
            <n-gi :span="12">
                <n-card>
                    <n-button type="primary" @click="addForm">
                        添加类型
                    </n-button>
                    <vxe-grid ref='xGrid' v-bind="gridOptions" v-on="gridEvent">
                        <template #action="{ row }">
                            <div>
                                <n-button type="primary" @click="deleteForm(row)">删除</n-button>
                                <n-button type="primary" @click="editForm(row)">编辑</n-button>
                            </div>
                        </template>
                    </vxe-grid>
                </n-card>
              <n-modal  v-model:show="open" width="1000px" title="用户"
                        :bodyStyle="{ height: '800px' }">
                <n-card style="width:600px">
                  <n-form>
                    <n-form-item>
                      <n-input v-model:value="flowtypeForm.type_name" placeholder="请输入类型名称"></n-input>
                    </n-form-item>
                    <n-form-item>
                      <n-button type="primary" @click="submitForm">提交</n-button>
                    </n-form-item>
                  </n-form>
                </n-card>
              </n-modal>
            </n-gi>
        </n-grid>
    </div>
</template>

<script setup lang="ts">
import {useMessage} from "naive-ui";
const message  = useMessage()
const {  gridOptions,storeFlowtype,updateFlowtype,deleteFlowtype } = useFlowtype();
const router = useRouter()
const xGrid = ref()
const open = ref(false)
const flowtypeForm = ref({
  id:null,
  type_name:""
})
const addForm=()=>{
  open.value=true
}
const editForm=(row)=>{
  flowtypeForm.value=row
  open.value=true
}
const deleteForm=async(row)=>{
  await deleteFlowtype(row.id)
}
const submitForm=async()=>{
  if(!flowtypeForm.value.type_name){
    message.error("请输入名称")
    return
  }
  if(flowtypeForm.value.id){
    const res=await updateFlowtype(flowtypeForm.value)
    if(res.code==200){
      open.value=false
    }
  }else{
    const res=await storeFlowtype(flowtypeForm.value)
    if(res.code==200){

      open.value=false
    }
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

</script>

<style scoped></style>
