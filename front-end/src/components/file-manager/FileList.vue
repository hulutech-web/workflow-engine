<template>
  <div class="mt-3">

    <a-row :gutter="16">
      <a-col :span="6">
        <a-table :data="folders" style="width: 100%"  border>
          <a-table-column prop="name" label="文件夹">
            <template #default="{ row }">
              <div class="flex justify-between items-center">
                <span class="cursor-pointer" @click="handleOpenFolder(row.id)">{{row.name}}</span>
                <a-space>
                  <a-icon @click="handleUpload(row.id)" class="cursor-pointer a-dropdown-link"><Upload /></a-icon>
                  <a-dropdown trigger="click">
                  <span class="a-dropdown-link">
                     <a-icon><MoreFilled /></a-icon>
                  </span>
                    <template #dropdown>
                      <a-dropdown-menu>
                        <a-dropdown-item @click="handleEditFolder">编辑</a-dropdown-item>
                        <a-dropdown-item @click="handleDeleteFolder(row.id)">删除</a-dropdown-item>
                      </a-dropdown-menu>
                    </template>
                  </a-dropdown>
                </a-space>
              </div>
            </template>
          </a-table-column>
        </a-table>
      </a-col>
      <a-col :span="18">
        <a-space>
          <a-input placeholder="请输入内容" v-model="keyword" style="width:400px;">
          </a-input>
          <a-button type="primary" danger @click="searchBy">搜索</a-button>
          <a-button type="primary" plain @click="keyword=''">清空</a-button>
        </a-space>

        <a-row :gutter="8"  class="min-h-[700px] mt-2">
          <a-col :span="6" v-for="(f,ind) in files" :key="ind" >
            <div class="w-[160px] h-[160px] text-center" style="border: .5px dotted #e5e7eb;">
              <a-image
                  ref="imageRef"
                  style="width: 160px; height: 160px"
                  :src="`http://localhost:3000${f.att_dir}`"
                  fit="cover"
                  :preview-src-list="[`http://localhost:3000${f.att_dir}`]"
              />
              <div class="text-center font-light text-xs text-blue-400">
                {{f.real_name}}
              </div>
              <div>
                <a-space>
                  <a-checkbox></a-checkbox>
                  <a-popconfirm title="确认删除吗?" @confirm="delAtt(f)">
                    <template #reference>
                      <a-icon :size="18" class="cursor-pointer"><DeleteFilled color="#F56C6C" /></a-icon>
                    </template>
                  </a-popconfirm>
                </a-space>
              </div>
            </div>
          </a-col>
        </a-row>

        <a-pagination
            class="mt-3"
            v-model:current-page="pageForm.page"
            v-model:page-size="pageForm.page_size"
            :page-sizes="[12,24,48,96]"
            :background="true"
            layout="total, sizes, prev, pager, next, jumper"
            :total="pageForm.total"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
        />
      </a-col>

      <FileUpload v-model="uploadVisible" @success="fetchData" v-model:cate_id="cate_id" />
    </a-row>
  </div>
</template>

<script setup>
import {ref} from "vue";
import {MoreFilled,DeleteFilled,Upload} from "@element-plus/icons-vue"
import Api from "@/api/backend/api"
import FileUpload from "@/components/ReFileManager/FileUpload.vue";
defineProps({
  files: Array,
  folders: Array,
});
const pageForm = ref({
  page:1,
  total:0,
  page_size:12,
})
const files = ref([]);

const fold_id=ref("")
const keyword = ref("")
const searchBy = ()=>{
  loadFolderData()
}

const delAtt=(row)=>{

  const url=`http://localhost:3000/api/v1/admin/system/attachment/${row.id}`
  fetch(url,{
    method:"DELETE"
  }).then(res => {
    return res.json()
  }).then(response=>{
    loadFolderData()
  })
}
const handleSizeChange = (val) => {
  pageForm.value.page = 1;
  pageForm.value.total = val;
  loadFolderData()
};

const handleCurrentChange = (val) => {
  pageForm.value.page = val;
  loadFolderData()
};
// 上传文件显示
const uploadVisible=ref(false)
//分类id
const cate_id=ref(null)
// 上传文件
const handleUpload=(id)=>{
  uploadVisible.value=true
  cate_id.value=id;
}

const emit = defineEmits(['open-folder', 'delete-folder', 'download-file', 'delete-file']);
const handleEditFolder=()=>{

}
const loadFolderData = async (id) =>{
  const response = await Api.attachmentCategoryController.attachments({id:id})
  console.log("渲染表单",response.total)
  files.value=response.data
  pageForm.value.total=response.total
  pageForm.value.page=response.meta.current_page
}
const handleOpenFolder = (id) => {
  fold_id.value = id
  console.log("打开文件夹",id)
  loadFolderData(id)
};
const fetchData=(val)=>{
  console.log(val)
}
const handleDeleteFolder = (id) => {
  emit('delete-folder', id);
};


</script>
<style>
.a-dropdown-link {
  cursor: pointer;
  color: #409EFF;
}
.a-icon-arrow-down {
  font-size: 12px;
}
</style>
