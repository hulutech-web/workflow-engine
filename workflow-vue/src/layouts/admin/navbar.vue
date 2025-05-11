<script setup lang="ts">
import FullScreen from '@/components/fullscreen/index.vue'
import Breadcrumb from '@/components/breadcrumb/index.vue'
import useMenuStore from '@/store/useMenuStore'
import router from '@/plugins/router'
const menuStore = useMenuStore()
const storage = useStorage()
const { open } = useUtil()
const keyword = ref("")
const onSearch = () => {
  console.log(keyword.value)
}

const logout = () => {
  localStorage.clear();
  router.push({ name: "auth.login" })
}
const collapsed = ref(false)
const emit = defineEmits(["changeCollapse"])
const updateCollapse = () => {
  collapsed.value = !collapsed.value
  emit("changeCollapse", collapsed.value)
}
const user = ref(JSON.parse(storage.get(CacheEnum.USER)))
</script>

<template>
  <div class="flex justify-between items-center w-full">
    <div>
      <menu-unfold-outlined style="font-size: 20px;" v-if="collapsed" class="trigger" @click="updateCollapse" />
      <menu-fold-outlined style="font-size: 20px;" v-else class="trigger" @click="updateCollapse" />
    </div>
    <div class="flex justify-center items-center" style="width: 400rpx;">
      <div class="text-xl font-bold">
        {{ user?.name }}
      </div>
      <a-space :size="24">
        <FullScreen />
      </a-space>
      <a-space :size="24">
        <a-dropdown :trigger="['hover']" placement="bottom" arrow>
          <a class="ant-dropdown-link" @click.prevent>
            <a-avatar src="https://xsgames.co/randomusers/avatar.php?g=pixel&key=1" />
          </a>
          <template #overlay>
            <a-menu>
              <a-menu-item key="0">
                <a href="#">个人中心</a>
              </a-menu-item>
              <a-menu-divider />
              <a-menu-item @click="logout" key="1">退出登录</a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </a-space>
    </div>
  </div>
</template>