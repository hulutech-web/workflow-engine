<template>
  <div>
    <el-card>
      <el-space>
        <el-button-group>
          <el-button type="success" @click="handleCreateFolder">创建文件夹</el-button>
        </el-button-group>
        操作：
        <el-button-group>
          <el-button type="success">确认选择</el-button>
          <el-button type="warning" plain>清空</el-button>
        </el-button-group>
      </el-space>

      <el-row>
        <el-col :span="24">
          <FileList :folders="folders"  @refresh="fetchData" />
        </el-col>
      </el-row>
      <FolderCreate v-model="folderCreateVisible" @success="fetchData" />
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import FileList from './FileList.vue';
import FileUpload from './FileUpload.vue';
import FolderCreate from './FolderCreate.vue';
import { storageLocal } from "@pureadmin/utils";
const getFolders = ()=>{
  const url=`http://localhost:3000/api/attachment_category/list`
  fetch(url,{
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + storageLocal().getItem('authorized-token')
    }
  }).then(res => {
    return res.json()
  }).then(response=>{
    // console.log("渲染表单",response.data)
    //添加一个未分组
    console.log(folders.value)
    let emptyFold = {
      enname:"",
      id:0,
      name:"未分组",
      pid:0,
      system_attachments:null,
    }
    folders.value=response.data
    folders.value.push(emptyFold)
  })
}
const folders = ref([]);
const uploadVisible = ref(false);
const folderCreateVisible = ref(false);

// 获取文件和文件夹数据
const fetchData = async () => {
  await getFolders();
};

// 打开上传文件弹窗
const handleUpload = () => {
  uploadVisible.value = true;
};

// 打开创建文件夹弹窗
const handleCreateFolder = () => {
  folderCreateVisible.value = true;
};

// 初始化时加载数据
onMounted(() => {
  fetchData();
});
</script>
