<template>
  <n-icon
    :size="size"
    :color="color"
    :depth="depth"
    :component="resolvedIcon"
  />
</template>

<script setup lang="ts">
import * as allIcons from '@vicons/ionicons5'
import { AlertCircleOutline } from '@vicons/ionicons5'
import { computed } from 'vue'
import type { Component } from 'vue'

interface Props {
  /** 图标名称 (例如: "GameController" 或 "CashOutline") */
  name?: string
  /** 图标大小 (可以传数字或字符串，如: 24 或 "24px") */
  size?: string | number
  /** 图标颜色 (十六进制颜色码) */
  color?: string
  /** 图标深度 (1-5) */
  depth?: number
  /** 找不到图标时是否显示警告图标 */
  showFallback?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  name: undefined,
  size: 24,
  color: undefined,
  depth: undefined,
  showFallback: true
})

// 解析图标组件
const resolvedIcon = computed<Component | undefined>(() => {
  if (!props.name) return undefined

  // 检查图标是否存在
  const iconExists = props.name in allIcons

  if (!iconExists) {
    console.warn(`图标 "${props.name}" 不存在于 @vicons/ionicons5 中`)
    return props.showFallback ? AlertCircleOutline : undefined
  }

  return allIcons[props.name as keyof typeof allIcons]
})
</script>
