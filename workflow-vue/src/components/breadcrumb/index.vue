<script setup lang="ts">
import router from '@/plugins/router'
const route = router.currentRoute
</script>

<template>
    <a-breadcrumb>
      <a-breadcrumb-item v-for="(r, index) of route?.matched" :key="index" v-show="r.meta.title">
        {{ r.meta.title }}
      </a-breadcrumb-item>
    </a-breadcrumb>
</template>

<style lang="scss" scoped>

</style>
