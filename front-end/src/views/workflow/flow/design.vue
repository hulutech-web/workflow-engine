<template>
  <div>
      <div class="h-8 text-right">
        <n-space>
          <n-button type="primary" danger size="small" @click="()=>initAll()">刷新</n-button>
          <n-button type="primary" size="small" @click="saveDesign">保存位置</n-button>
          <n-button type="primary" size="small" @click="publishDesign">发布流程</n-button>
        </n-space>
      </div>
    <div id="flow-chart-container">
      <HuluMenu :flow_id="+id" :init="initAll" ref="menuRef" />
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
          <span class="font-bold text-md">{{ node.process_name }}</span>
          <n-button
            type="primary"
            style="color: #ffffff; z-index: 20; background-color: #ffa500"
            @click="setProcess(node)"
            shape="circle"
          >
<!--            <FormOutlined class="node-setting" />-->
            <!-- <SettingOutlined  /> -->
          </n-button>
        </div>
      </div>
    </div>

    <n-modal
      v-model:open="open"
      style="position: relative"
      width="1200px"
      title="节点设计"
      centered
      :bodyStyle="{ height: '700px' }"
    >
      <Attrform :attrs="attrs" @updProcess="updProcess" />
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import Attrform from "./component/attrform/index.vue";
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

<style lang="scss">
#flow-chart-container {
  width: 100%;
  height: 1200px;
  border: 1px solid rgba(255, 255, 255, 0.1); // 调整边框透明度
  position: relative;
  background-color: #1a1a1a; // 主背景色
  background-image:
    linear-gradient(90deg, rgba(255, 255, 255, 0.05) 1px, transparent 1px),
    linear-gradient(180deg, rgba(255, 255, 255, 0.05) 1px, transparent 1px); // 网格线调整为白色透明度
  background-size: 20px 20px;
}

.node {
  position: absolute;
  text-align: center;
  // 溢出隐藏
}

/* 1、节点颜色 */
.node-element {
  background-color: #1a1a1a; // 主背景色
  border: 2px solid #e9ecef;
  border-radius: 4px;
  // 文字溢出隐藏，并使用省略号
  overflow: hidden;
  font-size: small;
  text-overflow: ellipsis;
  white-space: nowrap;
  padding: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: grab;

  &:hover {
    transform: translateY(-3px);
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
    border-color: #4a8cff;
  }

  .ant-btn-primary {
    background: linear-gradient(45deg, #4a8cff, #2d5aa3); // 按钮渐变调整
    border: none;
    transition: transform 0.2s ease;

    &:hover {
      transform: scale(1.05);
      box-shadow: 0 2px 8px rgba(74, 140, 255, 0.3);
    }
  }
}
.node-setting {
  cursor: pointer;
}
</style>
