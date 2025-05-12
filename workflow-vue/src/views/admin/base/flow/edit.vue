<template>
    <div class="p-3">
        <p>
            编辑流程
            <Form :id="id"/>
        </p>
    </div>

</template>

<script setup lang='ts'>
const route = useRoute()
let id = route.params.id
console.log(id)
import Form from './form.vue'
</script>

<style></style>