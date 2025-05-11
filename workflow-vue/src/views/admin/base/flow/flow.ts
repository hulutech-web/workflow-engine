import jsPlumb from "jsplumb";
import { Modal } from "ant-design-vue";
import { onContextMenu } from "./tool";
const { deleteProcess } = useProcess();
async function initFlowChart(initData, callback) {
  await nextTick(); // 确保DOM更新完成
  let data = toRaw(initData);
  if (!data.list || data.list.length == 0) {
    return;
  }

  // 动态维护一个节点列表，用于后续的拖拽操作，并将他返回
  let displayNodeList = [];

  const nodeList = [];
  let jsPlumbInstance = null;
  // 初始化节点列表

  Object.keys(data.list).forEach((key) => {
    const node = data.list[key];
    nodeList.push(node);
  });
  // console.log("nodeList", nodeList);
  // 初始化节点
  // 初始化jsPlumb实例
  jsPlumbInstance = jsPlumb.jsPlumb.getInstance({
    container: document.getElementById("flow-chart-container"),
    DragOptions: { cursor: "pointer", zIndex: 2000 },
    // 设置成折线
    Connector: [
      "Flowchart",
      { stub: [40, 60], gap: 10, cornerRadius: 5, alwaysRespectStubs: true },
    ],
    
    PaintStyle: { stroke: "#1d39c4", strokeWidth: 2 }, // 设置线条颜色和宽度
    EndpointStyle: { radius: 5, fill: "#1d39c4" }, // 设置端点样式
    HoverPaintStyle: { stroke: "#c92a2a", strokeWidth: 3 }, // 鼠标悬停时的线条样式
    ArrowType: "arrow", // 添加箭头，可选值有 "arrow", "diamond", 等
    ArrowSize: 10, // 箭头大小
  });

  jsPlumbInstance.ready(function () {
    jsPlumbInstance.importDefaults({
      ConnectionsDetachable: false,
    });
    async function canDelCbk(nodeId) {
      console.log(nodeId);

      let index = nodeList.findIndex((node) => `node-${node.id}` === nodeId);
      if (index != -1) {
        // 获取当前被删除的节点
        const delNode = nodeList[index];
        await deleteProcess({ id: delNode.id, flow_id: delNode.flow_id });
        let displayNodeList = nodeList;
        displayNodeList.splice(index, 1);
        callback(displayNodeList);
      }
    }
    const editCbk = (id)=>{
      console.log(`被编辑的节点id:${id}`)
    }
    nodeList.forEach((node) => {
      //设置可拖动
      jsPlumbInstance.draggable(`node-${node.id}`, {
        start: (event) => {
          // 拖动开始时的处理逻辑
          // console.log(`Dragging started for node ${node.id}`,event);
          // 禁止默认创建连线
        },
        drag: (event) => {
          // 阻止默认事件
        },
        stop: (event) => {
          displayNodeList = nodeList;
          // //需要触发更新逻辑，更新位置信息
          let index = nodeList.findIndex((item) => item.id === node.id);
          let node2 = nodeList[index];
          let newStyle = node2.style;
          // //replace
          newStyle = newStyle.replace(
            /left:(\d+)px;/,
            `left:${event.pos[0]}px;`
          );
          newStyle = newStyle.replace(/top:(\d+)px;/, `top:${event.pos[1]}px;`);
          displayNodeList[index]["style"] = newStyle;
          callback(displayNodeList);
        },
      });

      // 为每一个节点添加右键
      // 监听节点的右键事件
      let nodeEle = document.querySelector(`#node-${node.id}`);
      if (nodeEle) {
        nodeEle.addEventListener("contextmenu", (event) => {
          //阻止向上传递，因为父级也注册了contextmenu事件
          event.stopPropagation();
          event.preventDefault();
          onContextMenu(event, canDelCbk,editCbk);
        });
      }
    });

    // 连线检查
    // 监听节点连接前的验证事件
    jsPlumbInstance.bind("beforeDrop", function (info) {
      const sourceId = info.sourceId;
      const targetId = info.targetId;

      // 检查规则：自己不能连接自己
      if (sourceId === targetId) {
        console.log("自己不能连接自己");
        return false; // 返回 false 表示不允许连接
      }

      // 检查规则：不允许二次连接
      const connections = jsPlumbInstance.getConnections({ source: sourceId });
      for (let i = 0; i < connections.length; i++) {
        const connection = connections[i];
        if (connection.targetId === targetId) {
          console.log("已经连接，不允许二次连接");
          return false; // 返回 false 表示不允许连接
        }
      }
      // 检查规则：已经被连接的节点不允许反向连接
      const reverseConnections = jsPlumbInstance.getConnections({
        source: targetId,
      });
      for (let i = 0; i < reverseConnections.length; i++) {
        const reverseConnection = reverseConnections[i];
        if (reverseConnection.targetId === sourceId) {
          console.log("已经被连接，不允许反向连接");
          return false; // 返回 false 表示不允许连接
        }
      }
      // 其他情况允许连接
      return true;
    });
    jsPlumbInstance.bind("click", (conn, originalEvent) => {
      Modal.confirm({
        //居中
        centered: true,
        title: "提示",
        content: "确定删除所点击的链接吗？",
        okText: "确定",
        cancelText: "取消",
        onOk: () => {
          jsPlumbInstance.deleteConnection(conn);
        },
        onCancel() {
          console.log("Cancel");
        },
      });
    });

    // 根据process_to字段创建连接线
    nodeList.forEach((node) => {
      if (node.process_to) {
        node.process_to.split(",").forEach((targetId) => {
          jsPlumbInstance.connect({
            source: `node-${node.id}`,
            target: `node-${targetId}`,
            paintStyle: { stroke: "#4169E1", strokeWidth: 3 }, // 线条样式
            endpointStyle: { fill: "#4169E1", radius: 4 }, // 端点样式
            connector: [
              "Straight",
              {
                stub: [40, 60],
                gap: 24,
                cornerRadius: 24,
                alwaysRespectStubs: true,
              },
            ],
            options: {
              stub: 25,
            },
            anchor: "Continuous",
            overlays: [
              [
                "Arrow",
                {
                  location: 1, // 箭头位置，1为连接线末端
                  length: 20, // 箭头长度
                  width: 20, // 箭头宽度
                  direction: 2, // 箭头方向，1为默认方向
                  paintStyle: { fill: "#4169E1", stroke: "transparent" }, // 箭头填充色和边框色
                },
              ],
            ],
          });
        });
      }
      // 设置节点作为连接节点
      var common = {
        isSource: true,
        isTarget: true,
        Connector: ["Flowchart"],
        paintStyle: { stroke: "#4169E1", strokeWidth: 90 }, // 线条样式
        // 线条宽度
        endpointStyle: { fill: "#4169E1", radius: 4 }, // 端点样式
        options: {
          stub: 25,
        },
        anchors: [jsPlumb.DynamicAnchor], // 自动选择最佳连接点
        maxConnections: -1,
      };
      jsPlumbInstance.addEndpoint(`node-${node.id}`, {
        uuid: `node-${node.id}`,
        anchor: "Continuous",
        ...common,
      });
    });

    // 可选：添加连接后处理程序
    jsPlumbInstance.bind("connection", function (info) {
      console.log("连接已建立:", info.sourceId, info.targetId);
      let sourceNode = nodeList.find(
        (node) => node.id == info.sourceId.replace("node-", "")
      );
      if (sourceNode) {
        //修改process_to的值
        if (sourceNode.process_to == "") {
          //注意：这个时候可能会出现本身为空节点的情况
          sourceNode.process_to = info.targetId.replace("node-", "");
        } else {
          sourceNode.process_to += `,${info.targetId.replace("node-", "")}`;
        }
      }
      // 更新nodeList的子元素
      let index = nodeList.findIndex((node) => node.id === sourceNode.id);

      nodeList[index] = sourceNode;
      displayNodeList = nodeList;
      console.log("displayNodeList", displayNodeList);
      callback(displayNodeList);
      // 添加逻辑，当节点被点击时，添加连接点
      let targetId = info.targetId.replace("node-", "");
      let targetNode = nodeList.find((node) => node.id == targetId);
      if(targetId){
        var common={
          isSource:true,
          isTarget:true,
          Connector:["Flowchart"],
          paintStyle:{stroke:"#000000",strokeWidth:90},//线条样式
          endpointStyle:{fill:"#000000",radius:4},//端点样式
          options:{
            stub:25,
          },
          anchors:[jsPlumb.DynamicAnchor],//自动选择最佳连接点
          maxConnections:-1,
        };
        jsPlumbInstance.addEndpoint(`node-${targetNode.id}`, {
          uuid: `node-${targetNode.id}`,
          anchor: "Continuous",
          ...common,
        })
      }
    });

    jsPlumbInstance.bind("connectionDetached", function (info) {
      let isBreak = false;
      console.log("连接已断开:", info.sourceId, info.targetId);
      requestAnimationFrame(() => {
        let sourceNode = nodeList.find(
          (node) => node.id == info.sourceId.replace("node-", "")
        );
        if (sourceNode) {
          //修改process_to中的process_to值删除，通过字符串拆分后，找到节点并删除，再重新转字符串
          sourceNode.process_to = sourceNode.process_to
            .split(",")
            .filter((item) => item != info.targetId.replace("node-", ""))
            .join(",");
        }
        // 更新nodeList的子元素
        let index = nodeList.findIndex((node) => node.id === sourceNode.id);

        nodeList[index] = sourceNode;
        displayNodeList = nodeList;
        console.log("displayNodeList", displayNodeList);
        callback(displayNodeList);
      });
    });

    return nodeList;
  });

  //注册全局对象
  window.globaljsPlumb = jsPlumbInstance;
}

const deleteNodeCbk = () => {};
export default initFlowChart;
