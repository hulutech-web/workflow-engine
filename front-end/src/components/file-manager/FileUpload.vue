<template>
  <a-dialog v-model="innerVisible" title="上传文件" @close="handleClose" width="500" :center="true">
    <a-upload
        class="upload-demo"
        :data="{cate_id:my_id}"
        drag
        action="http://localhost:3000/api/upload"
        :headers="{'Authorization':`Bearer ${token}`}"
        multiple>
      <i class="a-icon-upload"></i>
      <div class="a-upload__text">将文件拖到此处，或<em>点击上传</em></div>
      <div class="a-upload__tip" slot="tip">只能上传jpg/png文件，且不超过500kb</div>
    </a-upload>
  </a-dialog>
</template>

<script setup>
import { ref } from 'vue';
import { ElMessage } from 'element-plus';
import {TokenKey} from "@/utils/auth.js";
import { storageLocal } from "@pureadmin/utils";

const token = ref(storageLocal().getItem(TokenKey))
const props = defineProps({
  modelValue: Boolean,
  cate_id:Number
});

const fileList=ref([]);

const uploadRef=ref();
const emit = defineEmits(['update:modelValue', 'success']);
const innerVisible = ref(props.modelValue);

watch(()=>props.modelValue,(val)=>{
  innerVisible.value=val;
})
const my_id=ref(null)
watch(()=>props.cate_id,(val)=>{
  console.log(val)
  my_id.value=val;
})

const submitUpload=()=> {
  uploadRef.value.upload.submit();
}

const handleClose = () => {
  emit('update:modelValue', false);
};
</script>
