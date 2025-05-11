<template>
  <a-select v-model:value="value" show-search placeholder="模糊搜索" style="width: 400px"
    :default-active-first-option="false" :show-arrow="false" :filter-option="false" :not-found-content="null"
    :options="data" @search="handleSearch" @change="handleChange"></a-select>
</template>
<script lang="ts" setup>
const { searchUser } = useUser();
let timeout: any;
let currentValue = '';
const props = defineProps({
  modelValue: {
    type: Number,
    default: 0,
  },
})

function fetch(value: string, callback: any) {
  if (timeout) {
    clearTimeout(timeout);
    timeout = null;
  }
  currentValue = value;

  const fake = async () => {
    const data = await searchUser(value)
    const resultData = data.map(item=>{
      return {
        label:item.name,
        value:item.ID
      }
    })
    callback(resultData);
  }

  timeout = setTimeout(fake, 300);
}

const data = ref<any[]>([]);
const value = ref(props.modelValue);

const emit = defineEmits(['update:modelValue'])
watch(value, (val) => {
  emit('update:modelValue', val)
})

const handleSearch = (val: string) => {
  console.log("val",val)
  fetch(val, (d: any[]) => (data.value = d));
};

const handleChange = (val: string) => {
  console.log("handleChange",val)
  emit('update:modelValue', val)
};
</script>