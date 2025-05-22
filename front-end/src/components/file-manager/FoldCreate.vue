<template>
  <a-dialog v-model="innerVisible" title="创建文件夹" @close="handleClose">
    <a-form :model="form" laba-width="100px">
      <a-form-item label="父级">
        <a-select v-model="form.pid">
          <a-option label="" v-for="(c,ind) in cates" :key="ind" :value="c.id">
            {{c.name}}
          </a-option>
        </a-select>
      </a-form-item>
      <a-form-item label="文件夹名称">
        <a-input v-model="form.name"></a-input>
      </a-form-item>
    </a-form>
    <template #footer>
      <a-button @click="handleClose">取消</a-button>
      <a-button type="primary" @click="handleCreate">创建</a-button>
    </template>
  </a-dialog>
</template>

<script setup>
import { ref } from 'vue';
import Api from "@/api/backend/api"
import { ElMessage } from 'element-plus';
const createFolder = async ()=>{
  await Api.attachmentCategoryController.store(form.value)
}

const cates = ref([])
const loadCate = async()=>{
  const {data} = await Api.attachmentCategoryController.list()
  cates.value=data;
}
const props = defineProps({
  modelValue: Boolean
});
const innerVisible = ref(props.modelValue);

watch(()=>props.modelValue,(val)=>{
  console.log("watch",val)
  innerVisible.value=val;
},{immediate:true})
const emit = defineEmits(['update:modelValue', 'success']);

const form = ref({
  pid:null,
  name: '',
});

const handleCreate = async () => {
  await createFolder(form.value);
  ElMessage.success('文件夹创建成功');
  emit('success');
  emit('update:modelValue', false);
};

const handleClose = () => {
  emit('update:visible', false);
};
</script>
