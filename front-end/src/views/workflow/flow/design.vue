<template>
  <div>
    <div class="h-8 text-right">
      <n-space>
        <n-button type="primary" danger size="small" @click="()=>initAll()">刷新</n-button>
        <n-button type="primary" size="small" @click="saveDesign">保存</n-button>
        <n-button type="primary" size="small" @click="publishDesign">发布流程</n-button>
        <div>
          <n-gradient-text type="error">
            ·空白右键【新建节点】
          </n-gradient-text>
          <n-gradient-text type="error">
            ·节点右键【删除节点】
          </n-gradient-text>
          <n-gradient-text type="error">
            ·保存【位置与连线】
          </n-gradient-text>
          <n-gradient-text type="error">
            ·连线【点击节点从圆点拖拽到另一节点圆点】
          </n-gradient-text>
          <n-gradient-text type="error">
            ·断线【点击连线删除】
          </n-gradient-text>
          <n-gradient-text type="error">
            ·配置【点击🔧进入】
          </n-gradient-text>
        </div>
      </n-space>
    </div>
    <div id="flow-chart-container">
      <div>
        <HuluMenu :flow_id="+id" :init="initAll" ref="menuRef"/>
      </div>
      <!-- 动态生成节点 -->
      <div
        v-for="(node, nodeId) in nodeList"
        :key="node.id"
        :class="'node' + (node.process_to ? ' source-node' : '')"
        :id="'node-' + node.id"
        :style="node.style"
      >
        <div
          class="flex  justify-around items-center  node-element"
          :id="`menu-${node.id}`"
        >
          <div class="flex  justify-between items-center">

            <NovaIcon
              :id="`node-line-${node.id}-pointer`"
              fontSize="28px"
              :icon="node.icon"
              color="#66CDAA"
            />
            <span class="font-bold text-md">{{ node.process_name }}</span>
          </div>
          <n-button quaternary  @click="setProcess(node)" class="setting-btn">
            <template #icon>
              <NovaIcon
                :icon="'tdesign:tools-circle-filled'"
                color="#66CDAA"
              />
            </template>
          </n-button>
        </div>
      </div>
    </div>

    <n-modal
      v-model:show="open"
      style="position: relative"
      transform-origin="center"
    >
      <n-card
        style="width: 1000px;height:800px;"
        title="节点属性"
        :bordered="false"
        size="huge"
        role="dialog"
        aria-modal="true"
      >
      <Attrform :attrs="attrs" @updProcess="updProcess"/>
      </n-card>
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import Attrform from "./component/attrform/index.vue";
import initFlowChart from "./flow";
const route = useRoute();
const router = useRouter();
const {loadFlowDesign, publishFlow} = useFlow();
const {updateFlowlink} = useFlowlink();
const {loadAttributes, updateProcess} = useProcess();
const id = route.params.id;
const jsplumbJSON = ref({});
const nodeList = ref([]);
const flow = ref({});
const menuRef = ref({});
const open = ref(false);

const init = async () => {
  const data = await loadFlowDesign(+id);

  flow.value = data.data;
  // console.log(flow.value)
  if (flow.value.jsplumb) {
    jsplumbJSON.value = JSON.parse(flow.value.jsplumb);
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
  console.log(data.data);
  process_id.value = node.id;
  attrs.value = data.data;
};
const getNewestNodes = async (nodes) => {
  //获取最新的节点信息

  if (jsplumbJSON.value.total == 0) {
    await storeFlow(+id, nodes)
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
  await publishFlow({flow_id: flow.value.id});
};
</script>

<style lang="scss">
#flow-chart-container {
  width: 100%;
  height: 1200px;
  border: 1px solid rgba(255, 255, 255, 0.1); // 调整边框透明度
  position: relative;
  background-color: #1a1a1a; // 主背景色
  background-image: linear-gradient(90deg, rgba(255, 255, 255, 0.05) 1px, transparent 1px),
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

.setting-btn{
//悬浮时旋转270度，在1s钟内完成，平滑，同时高光显示
  transition: transform 1s ease-in-out;
  i{
    border-radius: 50%;
  }
  &:hover{
    //父组件下的i标签高亮,柔光包围效果
    i{
      box-shadow: 0 0 10px rgba(255, 255, 255, 0.5);
      transition: all 0.5s ease-in-out;
      background-color: rgba(255, 255, 255, 0.1);
    }

    transform: rotate(720deg);
  }
}

</style>
