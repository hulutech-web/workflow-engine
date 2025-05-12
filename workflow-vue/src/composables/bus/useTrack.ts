import { http } from "@/plugins/axios";
import router from "@/plugins/router";
import XEUtils from "xe-utils";
const storage = useStorage();

export default () => {
    //历史轨迹
  const indexTrack = async (id,start_at,end_at) => {
    return await http.request({
      url: `device/${id}/historytrack?start_at=${start_at}&end_at=${end_at}`,
      method: "GET",
    });
  };

  const serveApiUrl = import.meta.env.VITE_API_URL;
  const gridOptions = reactive<VxeGridProps<RowVO>>({
    border: "full",
    // size: "small",
    showHeaderOverflow: true,
    showOverflow: true,
    keepSource: true,
    autoResize: true,
    expandConfig: {
      trigger: "row",
      showIcon: true,
      expandRowKeys: [1, 2],
      iconOpen: "vxe-icon--caret-bottom",
      iconClose: "vxe-icon--caret-top",
    },
    formConfig: {
      titleWidth: 120,
      titleAlign: "right",
      items: [
        {
          field: "code",
          title: "设备编码",
          span: 6,
          titlePrefix: {
            useHTML: true,
            message: "模糊查询",
            icon: "vxe-icon-question-circle-fill",
          },
          itemRender: { name: "$input", props: { placeholder: "请输入编码" } },
        },
        {
          field: "name",
          title: "设备名称",
          span: 6,
          titlePrefix: {
            useHTML: true,
            message: "模糊查询",
            icon: "vxe-icon-question-circle-fill",
          },
          itemRender: { name: "$input", props: { placeholder: "设备名称" } },
        },

        {
          span: 24,
          align: "left",
          collapseNode: true,
          itemRender: {
            name: "$buttons",
            children: [
              {
                props: {
                  type: "submit",
                  content: "搜索",
                  status: "primary",
                },
              },
              { props: { type: "reset", content: "重置" } },
            ],
          },
        },
      ],
    },
    stripe: true,
    id: "full_edit_9999",
    rowConfig: {
      keyField: "#",
      isHover: true,
      useKey: true,
    },
    columnConfig: {
      resizable: true,
    },
    customConfig: {
      storage: true,
      checkMethod({ column }) {
        if (["name", "no"].includes(column.field)) {
          return false;
        }
        return true;
      },
    },
    printConfig: {
      columns: [{ field: "name" }],
    },
    sortConfig: {
      trigger: "cell",
      remote: true,
    },
    filterConfig: {
      remote: true,
    },
    pagerConfig: {
      enabled: true,
      pageSize: 5,
      pageSizes: [5, 10, 15, 20, 50, 100, 200, 500, 1000],
    },

    toolbarConfig: {
      buttons: [
        // {code: "insert_actived", name: "快捷新增", status: "primary"},
        // {code: "delete", name: "直接删除", status: "danger"},
        // {code: "mark_cancel", name: "删除/取消"},
        // {
        //   code: "save",
        //   name: "快捷改价",
        //   status: "success",
        //   click: () => {
        //     console.log("save");
        //   },
        // },
      ],
      refresh: true, // 显示刷新按钮
      import: false, // 显示导入按钮
      export: true, // 显示导出按钮
      print: true, // 显示打印按钮
      zoom: true, // 显示全屏按钮
      custom: true, // 显示自定义列按钮
    },
    proxyConfig: {
      seq: true, // 启用动态序号代理，每一页的序号会根据当前页数变化
      sort: true, // 启用排序代理，当点击排序时会自动触发 query 行为
      filter: true, // 启用筛选代理，当点击筛选时会自动触发 query 行为
      form: true, // 启用表单代理，当点击表单提交按钮时会自动触发 reload 行为
      props: {
        // 对应响应结果 Promise<{ result: [], page: { total: 100 } }>
        result: "data", // 配置响应结果列表字段
        total: "total", // 配置响应结果总页数字段
      },
      // 只接收Promise，具体实现自由发挥
      ajax: {
        // 当点击工具栏查询按钮或者手动提交指令 query或reload 时会被触发
        query: ({ page, sorts, filters, form }) => {
          //如果additionalQueryParams为空对象时不处理，否则合并到form中
          if (additionalQueryParams.value.user_id != 0) {
            form = { ...form, ...additionalQueryParams.value }
          }
          return new Promise((resolve, reject) => {
            const queryParams: any = Object.assign({}, form);
            // 处理排序条件
            const firstSort = sorts[0];
            if (firstSort) {
              queryParams.sort = firstSort.field;
              queryParams.order = firstSort.order;
            }
            // 处理筛选条件
            filters.forEach(({ field, values }) => {
              queryParams[field] = values.join(",");
            });

            const data = http.request({
              url: `device?pageSize=${page.pageSize}&currentPage=${page.currentPage
                }&${XEUtils.serialize(queryParams)}`,
              method: "GET",
            });
            resolve(data);
          });
        },
        save: ({ body }) => {
          return new Promise((resolve, reject) => {
            //删除item中的#字段
            let postData = body.updateRecords.map(item => {
              return {
                id: item.ID,
                price: +item.price
              }
            })
            const data = batchUpdate(postData);
            resolve(data);
          });
        },
      },
    },
    columns: [
      { field: "ID", title: "ID", width: 100 },
      // 配置日期选择器
      {
        field: "user_id",
        title: "运行商ID",
        sortable: true,

      },
      {
        field: "code",
        title: "设备编码",
        sortable: true,
      },
      {
        field: "name",
        title: "设备名称",
        sortable: true,
      },
      {
        field: "board_model",
        title: "主板型号",

      },
      {
        field: "online_at",
        title: "上线时间",
        sortable: true,
      },
      {
        field: "location_at",
        title: "定位时间",
        sortable: true,
      },
      {
        field: "battery_status",
        title: "电池状态",
        sortable: true,
      },
      {
        field: "spec",
        title: "规格",
        sortable: true,
      },
      {
        field: "tower_code",
        title: "基站编码",
        sortable: true,
      },
      {
        field: "serial_number",
        title: "序列号",
        sortable: true,
      },
      {
        field: "action",
        width: 300,
        title: "操作",
        slots: {
          default: "action",
        },
      },
    ],
    importConfig: {},
    exportConfig: {},
    checkboxConfig: {
      labelField: "ID",
      checkStrictly: true,
    },
    editRules: {
      name: [
        { required: true, message: "必填" },
        { min: 3, max: 50, message: "名称长度在 3 到 50 个字符" },
      ],
      no: [{ required: true, message: "必填" }],
    },
    editConfig: {
      trigger: "click", // 设置为 'click' 表示点击单元格时自动进入编辑状态
      mode: "row",
      showStatus: true,
      showUpdateStatus: true,
      showInsertStatus: true,
      autoClear: true,
    },
  });

  return {
    gridOptions,
    indexTrack,
  };
};
