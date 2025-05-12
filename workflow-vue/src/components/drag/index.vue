<template>
  <div class="flex">
    <div class="text-bold">选择</div>
    <VueDraggable
      class="flex flex-col gap-2 p-4 w-300px h-300px m-auto bg-gray-200 rounded overflow-auto"
      v-model="list1"
      animation="150"
      ghostClass="ghost"
      group="people"
      @onChoose="onChoose"
      @update="onUpdate"
      @add="onAdd"
      @remove="remove"
    >
      <div
        v-for="item in list1"
        :key="item.id"
        class="cursor-move h-30 bg-gray-900 rounded p-3 text-white"
      >
        {{ item.name }}
      </div>
    </VueDraggable>

    <div class="text-bold">候选区</div>
    <VueDraggable
      class="flex flex-col gap-2 p-4 w-300px h-300px m-auto bg-green-300 rounded overflow-auto"
      v-model="list2"
      animation="150"
      group="people"
      ghostClass="ghost"
      @update="onUpdate"
      @add="onAdd"
      @remove="remove"
    >
      <div
        v-for="item in list2"
        :key="item.id"
        class="cursor-move h-30 bg-green-400 rounded p-3"
      >
        {{ item.name }}
      </div>
    </VueDraggable>
  </div>
  <!-- <div class="flex justify-between">
    <PreviewList :list="list1" />
    <PreviewList :list="list2" />
  </div> -->
</template>

<script setup>
import { ref } from 'vue'
import { VueDraggable } from 'vue-draggable-plus'
const list1 = ref([
  {
    name: 'Joao',
    id: '1'
  },
  {
    name: 'Jean',
    id: '2'
  },
  {
    name: 'Johanna',
    id: '3'
  },
  {
    name: 'Juan',
    id: '4'
  }
])
const list2 = ref(
  list1.value.map(item => ({
    name: `${item.name}-2`,
    id: `${item.id}-2`
  }))
)
function onUpdate() {
  console.log('update')
}
function onAdd() {
  console.log('add')
}
function remove() {
  console.log('remove')
}
const onChoose = (event, item) => {
  console.log('choose', event, item)
}
</script>
