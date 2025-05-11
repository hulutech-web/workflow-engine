<template>
  <a-result status="403" title="403" sub-title="Sorry, you are not authorized to access this page.">
    <template #extra>
      <a-button type="primary" @click="backTo">Back Home</a-button>
    </template>
  </a-result>
</template>
<script lang="ts" setup>
const router = useRouter();
const backTo = () => {
  router.replace("/admin")
}
</script>