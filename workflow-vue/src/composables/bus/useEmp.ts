import { http } from "@/plugins/axios";
import router from "@/plugins/router";
import XEUtils from "xe-utils";
const storage = useStorage();

export default () => {
  // 方法
  const loadEmps = async (id) => {
    return await http.request({
      url: `emp`,
      method: "GET",
    });
  };

  const storeEmp = async (data) => {
    return await http.request({
      url: `emp`,
      method: "POST",
      data: data,
    });
  };

  const showEmp = async (id) => {
    return await http.request({
      url: `emp/${id}`,
      method: "GET",
    });
  };

  const updateEmp = async (data) => {
    return await http.request({
      url: `emp/${data.ID}`,
      method: "PUT",
      data: data,
    });
  };

  const deleteEmp = async (id) => {
    return await http.request({
      url: `emp/${id}`,
      method: "DELETE",
    });
  };

  const searchEmp = async (name) => {
    return await http.request({
      url: `emp/search`,
      method: "POST",
      data: { name: name },
    });
  };

  const getEmpOpt = async () => {
    return await http.request({
      url: `emp/options`,
      method: "GET",
    });
  };

//   员工绑定用户
  const bindUser = async (data) => {
    return await http.request({
      url: `emp/bind`,
      method: "POST",
      data: data,
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
          field: "name",
          title: "名称",
          span: 6,
          titlePrefix: {
            useHTML: true,
            message: "模糊查询",
            icon: "vxe-icon-question-circle-fill",
          },
          itemRender: { name: "$input", props: { placeholder: "请输入名称" } },
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
      pageSize: 10,
      pageSizes: [5, 10, 15, 20, 50, 100, 200, 500, 1000],
    },

    toolbarConfig: {
      buttons: [
        // {code: "insert_actived", name: "快捷新增", status: "primary"},
        // {code: "delete", name: "直接删除", status: "danger"},
        // {code: "mark_cancel", name: "删除/取消"},
        // {code: "save", name: "保存", status: "success"},
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
              url: `emp?pageSize=${page.pageSize}&currentPage=${
                page.currentPage
              }&${XEUtils.serialize(queryParams)}`,
              method: "GET",
            });
            resolve(data);
          });
        },
      },
    },
    columns: [
      { field: "id", title: "id", width: 100 },
      // 配置日期选择器
      {
        field: "name",
        title: "名称",
        sortable: true,
      },
      {
        field: "workno",
        title: "工号",
        sortable: true,
      },
      {
        field: "email",
        title: "邮箱",
        sortable: true,
      },
      {
        field: "Dept",
        title: "部门",
        sortable: true,
        slots: {
          default: "dept",
        },
      },
      {
        title: "操作",
        sortable: true,
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
      trigger: "click",
      mode: "row",
      showStatus: true,
      showUpdateStatus: true,
      showInsertStatus: true,
      autoClear: true,
    },
  });

  return {
    gridOptions,
    loadEmps,
    updateEmp,
    getEmpOpt,
    deleteEmp,
    searchEmp,
    bindUser,
    storeEmp,
    showEmp,
  };
};
