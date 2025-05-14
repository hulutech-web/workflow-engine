<template>
  <div>
    <!-- 添加ref以便调试 -->
    <context-menu
      ref="menuRef"
      v-model:show="show"
      :options="optionsComponent"
    >
      <context-menu-item @click="addNode">
        <div class="h-12 flex items-center px-4">
          <span>添加步骤</span>
        </div>
      </context-menu-item>
      <context-menu-item @click="onMenuClick(2)">
        <div class="h-12 flex items-center px-4">
          <span>刷新</span>
        </div>
      </context-menu-item>
    </context-menu>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { ContextMenu, ContextMenuItem } from '@imengyu/vue3-context-menu';
import '@imengyu/vue3-context-menu/lib/vue3-context-menu.css'; // 确保导入样式
import { request } from "@/service/http";

export default defineComponent({
  components: {
    ContextMenu,
    ContextMenuItem
  },
  props: {
    flow_id: {
      type: Number,
      default: 0
    },
    init: {
      type: Function,
      default: () => {}
    },
  },
  data() {
    return {
      show: false,
      optionsComponent: {
        theme: 'flat',
        zIndex: 3000, // 提高zIndex确保在最上层
        minWidth: 230,
        x: 0, // 初始化为0
        y: 0  // 初始化为0
      }
    }
  },
  mounted() {
    this.listenRightClick();
  },
  beforeUnmount() {
    this.removeRightClickListener();
  },
  methods: {
    setPos(pos) {
      this.optionsComponent.x = pos.x;
      this.optionsComponent.y = pos.y;
    },
    async addNode(e) {
      const container = document.querySelector("#flow-chart-container");
      if (!container) return;

      const rect = container.getBoundingClientRect();
      const x = e.clientX - rect.left;
      const y = e.clientY - rect.top;

      await request.Post(`process`, {
        flow_id: this.flow_id,
        left: `${x}px`,
        top: `${y}px`,
      });

      if (this.init) {
        //刷新当前页面
        window.location.reload();
        // await this.init();
      }
    },
    listenRightClick() {
      const container = document.querySelector("#flow-chart-container");
      if (container) {
        container.addEventListener("contextmenu", this.handleContextMenu);
      }
    },
    removeRightClickListener() {
      const container = document.querySelector("#flow-chart-container");
      if (container) {
        container.removeEventListener("contextmenu", this.handleContextMenu);
      }
    },
    handleContextMenu(e) {
      e.preventDefault();
      console.log('Right click at:', e.clientX, e.clientY);

      this.optionsComponent.x = e.clientX;
      this.optionsComponent.y = e.clientY;
      this.show = true;

      // 强制更新视图
      this.$nextTick(() => {
        console.log('Menu should be visible now');
      });
    },
    onMenuClick(index: number) {
      if (index === 2) {
        this.$router.go(0);
      }
    }
  }
});
</script>

<style scoped>
/* 确保菜单不被遮挡 */
:deep(.mx-context-menu) {
  position: fixed;
  z-index: 9999;
}
</style>
