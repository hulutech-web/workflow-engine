<script setup lang="ts">
import { nextTick, provide, ref } from 'vue'

const props = defineProps({
  async: {
    type: Boolean,
    default: false,
  },
  title: {
    type: String,
    default: '斯通纳',
  },
  disabled: {
    type: Boolean,
    default: false,
  },
  width: {
    type: Number,
    default: 502,
  },
  autoFocus: {
    type: Boolean,
    default: false,
  },
  confirmText: {
    type: String,
    default: '确定',
  },
  cancelText: {
    type: String,
    default: '取消',
  },
  showMask: {
    type: Boolean,
    default: true,
  },
  showClose: {
    type: Boolean,
    default: true,
  },
  maskClosable: {
    type: Boolean,
    default: true,
  },
  content: {
    type: String,
    default: '',
  },
})
const emit = defineEmits(['confirm', 'cancel', 'close', 'open'])
const visible = ref(false)
function handleEvent(type: 'confirm' | 'cancel') {
  emit(type)
  if (!props.async || type === 'cancel') {
    close()
  }
}
function close() {
  visible.value = false
  nextTick(() => {
    emit('close')
  })
}

function open() {
  if (props.disabled) {
    return
  }
  nextTick(() => {
    emit('open')
    visible.value = true
  })
}

provide('visible', visible)

defineExpose({
  open,
  close,
  visible,
  handleEvent,
})
</script>

<template>
  <div class="drawer">
    <n-drawer
      v-model:show="visible"
      :width="props.width"
      :auto-focus="props.autoFocus"
      :disabled="props.disabled"
      :before-close="close"
      :show-mask="props.showMask"
      :show-close="props.showClose"
      :mask-closable="props.maskClosable"
    >
      <n-drawer-content :title="props.title" closable>
        <template #default>
          <div class="content">
            <slot>
              <span v-if="props.content">{{ props.content }}</span>
            </slot>
          </div>
        </template>
        <template #footer>
          <n-button type="primary" @click="handleEvent('confirm')">{{ props.confirmText }}</n-button>
          <n-button @click="handleEvent('cancel')">{{ props.cancelText }}</n-button>
        </template>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<style scoped>

</style>
