<template>
  <div>
    <context-menu v-model:show="show" :options="optionsComponent">
      <context-menu-item @click="addNode">
        <template #default>
          <div class="h-12">
            <HuluIcon name="icon-add-location"></HuluIcon>
            <span>添加步骤</span>
          </div>
        </template>
      </context-menu-item>
      <context-menu-item @click="onMenuClick(2)">
        <template #default>
          <div class="h-12">
            <HuluIcon name="icon-refresh"></HuluIcon>
            <span>刷新</span>
          </div>
        </template>
      </context-menu-item>
    </context-menu>
  </div>
</template>
<script lang="ts">
import { http } from "@/plugins/axios";
export default {
  props: {
    flow_id: {
      type: Number,
      default: 0
    },
    init: {
      type: Function,
      default: () => { }
    },
  },
  data() {
    return {
      show: false,
      //For component
      optionsComponent: {
        theme: 'flat',
        zIndex: 3,
        minWidth: 230,
        x: 500,
        y: 200
      }
    }
  },
  mounted() {
    this.listenRightClick()
  },
  expose: ['setPos'],
  methods: {
    setPos(pos) {
      this.optionsComponent.x = pos.x;
      this.optionsComponent.y = pos.y;
    },
    async addNode(e) {
      //当前节点相对节点flow-chart-container的相对位置
      let x = e.x - document.querySelector("#flow-chart-container")?.getBoundingClientRect().x;
      let y = e.y - document.querySelector("#flow-chart-container")?.getBoundingClientRect().y;
      let node = {
        flow_id: this.flow_id,
        left: `${x}px`,
        top: `${y}px`,
      }
      await http.request({
        url: `process`,
        method: "POST",
        data: node
      });
      if (this.init) {
        await this.init();
      }

    },
    listenRightClick() {
      let pNode = document.querySelector("#flow-chart-container")
      pNode?.addEventListener
        ("contextmenu", (e) => {
          e.preventDefault();
          // 显示组件菜单
          this.show = true;
          this.optionsComponent.x = e.x;
          this.optionsComponent.y = e.y;
        });
    },
    onMenuClick(index: number) {
      // 1保存
      // 2刷新
      if (index == 2) {
        // 使用router刷新当前页面router.go(0)
        this.$router.go(0)
      }
      console.log('onMenuClick', index);
    }
  }
}
</script>