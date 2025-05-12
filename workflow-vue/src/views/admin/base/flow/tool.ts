import ContextMenu from "@imengyu/vue3-context-menu";
import { Modal, message as Message } from "ant-design-vue";

const onContextMenu = (
  e: MouseEvent,
  callback: Function,
  editCbk: Function
) => {
  //prevent the browser's default menu
  e.preventDefault();
  let ele = e.target as HTMLElement;
  // 获取节点的id
  let menu_id = ele.getAttribute("id");
  if (!menu_id) {
    return;
  }
  let id = menu_id.split("-")[1];

  //show your menu
  ContextMenu.showContextMenu({
    x: e.x,
    y: e.y,
    items: [
      {
        label: "属性设置",
        svgIcon: "#icon-operation",
        svgProps: {
          fill: "#096dd9",
        },
        onClick: () => {
            Modal.confirm({
                //居中
                centered: true,
                title: "提示",
                content: `确定编辑节点id:${id}吗？`,
                okText: "确定",
                cancelText: "取消",
                onOk: () => {
                  editCbk(id);
                },
                onCancel() {
                  console.log("Cancel");
                },
              });
        },
      },
      {
        label: "删除节点",
        svgIcon: "#icon-delete",
        svgProps: {
          fill: "#f60",
        },
        onClick: () => {
          Modal.confirm({
            //居中
            centered: true,
            title: "提示",
            content: `确定删除节点id:${id}吗？`,
            okText: "确定",
            cancelText: "取消",
            onOk: () => {
              let global = window.globaljsPlumb;
              const nodeId = `node-${id}`;

              // 检查节点是否存在
              let arrSources = global
                .getAllConnections()
                .map((item) => item.sourceId);
              let arrTargets = global
                .getAllConnections()
                .map((item) => item.targetId);
              let newArr = [...arrSources, ...arrTargets];
              let arr = Array.from(new Set(newArr));
              if (arr.includes(nodeId)) {
                Message.error("该节点已被连线，无法删除！");
                console.log(
                  `Node ${nodeId} is still connected to other nodes.`
                );
                return;
              }
              jsPlumb.remove(nodeId);
              callback(nodeId);
            },
            onCancel() {
              console.log("Cancel");
            },
          });
        },
      },
    ],
  });
};
export { onContextMenu };
