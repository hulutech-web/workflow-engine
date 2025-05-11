<template>
  <div>
    <a-card>
      <div>
        <a-space>
          <a-button type="primary" @click="saveDesign">保存位置</a-button>
          <a-button type="primary" @click="publishDesign">发布流程</a-button>
        </a-space>
      </div>
    </a-card>
    <div id="flow-chart-container">
      <hulu-menu :flow_id="+id" :init="initAll" ref="menuRef" />
      <!-- 动态生成节点 -->

      <div
        v-for="(node, nodeId) in nodeList"
        :key="node.id"
        :class="'node' + (node.process_to ? ' source-node' : '')"
        :id="'node-' + node.id"
        :style="node.style"
      >
        <div
          class="flex justify-center align-items-center node-element"
          :id="`menu-${node.id}`"
        >
          <HuluIcon
            :id="`node-line-${node.id}-pointer`"
            fontSize="28px"
            :name="node.icon"
            color="#66CDAA"
          />
          <span class="font-bold text-lg">{{ node.process_name }}</span>
          <a-button
            type="primary"
            style="color: #ffffff; z-index: 20; background-color: #ffa500"
            @click="setProcess(node)"
            shape="circle"
          >
            <FormOutlined class="node-setting" />
            <!-- <SettingOutlined  /> -->
          </a-button>
        </div>
      </div>
    </div>

    <a-modal
      v-model:open="open"
      style="position: relative"
      width="1200px"
      :footer="false"
      title="节点设计"
      centered
      :bodyStyle="{ height: '700px' }"
    >
      <attrform :attrs="attrs" @updProcess="updProcess" />
    </a-modal>
  </div>
</template>

<script setup lang="ts">
import initFlowChart from "./flow";
const route = useRoute();
const router = useRouter();
const { loadFlowDesign, publishFlow } = useFlow();
const { updateFlowlink } = useFlowlink();
const { loadAttributes, updateProcess } = useProcess();
const id = route.params.id;
const jsplumbJSON = ref({});
const nodeList = ref([]);
const flow = ref({});
const menuRef = ref({});
const open = ref(false);

const init = async () => {
  const data = await loadFlowDesign(+id);

  flow.value = data;
  // console.log(flow.value)
  if (data.jsplumb) {
    jsplumbJSON.value = JSON.parse(data.jsplumb);
    nodeList.value = jsplumbJSON.value.list;
    Object.entries(nodeList.value).map(([key, value]) => {
      value.flow_id = +id;
    });
  }
};

onMounted(async () => {
  await initAll();
});

const updProcess = async (val) => {
  await updateProcess(process_id.value, val);
  // setTimeout(() => {
  //     router.go(0)
  // }, 500)
};

const initAll = async () => {
  await init();
  await initFlowChart(jsplumbJSON.value, getNewestNodes);
};
const saveDesign = async () => {
  // 保存设计逻辑
  console.log(JSON.parse(flow.value.jsplumb));
  await updateFlowlink(flow.value);
};

const attrs = ref({});
const process_id = ref(0);
const setProcess = async (node) => {
  //阻止点击事件向下穿透
  open.value = true;
  const data = await loadAttributes(node.id);
  process_id.value = node.id;
  attrs.value = data;
};
const getNewestNodes = async (nodes) => {
  //获取最新的节点信息
  console.log("getNewestNodes", nodes);
  if (jsplumbJSON.value.total == 0) {
    await storeFlow(+id,nodes)
  } else {
    let newJsplumb = {
      total: nodes.length,
      list: "",
    };
    let list = Object.create(null);
    for (let i = 0; i < newJsplumb.total; i++) {
      let node = nodes[i];
      list[node.id + ""] = node;
    }
    newJsplumb.list = list;
    flow.value.jsplumb = JSON.stringify(newJsplumb);
  }
};

const publishDesign = async () => {
  // 发布设计逻辑
  await publishFlow({ flow_id: flow.value.id });
};
</script>

<style>
#flow-chart-container {
  width: 100%;
  height: 750px;
  border: 1px solid #ccc;
  position: relative;
  background-image: linear-gradient(
      90deg,
      rgba(200, 200, 200, 0.2) 1px,
      transparent 1px
    ),
    linear-gradient(180deg, rgba(200, 200, 200, 0.2) 1px, transparent 1px);
  background-size: 20px 20px;
  /* 这控制了网格线之间的间距 */
}

.node {
  position: absolute;
  text-align: center;
}

/* 1、节点颜色 */
.node-element {
  /* 设置节点的基础样式 */
  background-color: #fffacd;
  border: 1px solid #fffacd;
  border-radius: 5px;
  padding: 10px;
}

/* 如果需要针对特定状态（如鼠标悬停）设置样式 */
.node-element:hover {
  background-color: #08140e;
  cursor: move;
}

.node-setting {
  cursor: pointer;
}
</style>
