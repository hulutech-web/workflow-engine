//平移组件
import Translation from "@/views/admin/Translation.vue";
export default [
  {
    path: "/admin",
    redirect: "/admin/dashboard",
    component: () => import("@/layouts/admin/index.vue"),
    children: [
      {
        name: "admin.dashboard",
        path: "dashboard",
        component: () => import("@/views/admin/base/dashboard/index.vue"),
        meta: {
          title: "工作台",
          menu: { title: "工作台", icon: "DashboardOutlined", order: 1000 },
        },
        children: [],
      },
      {
        name: "admin.base",
        path: "base",
        component: Translation,
        meta: {
          title: "流程框架",
          menu: { title: "流程框架", icon: "SettingOutlined", order: 2000 },
        },
        children: [
          {
            name: "admin.user",
            path: "user",
            component: Translation,
            meta: {
              auth: true,
              title: "人员管理",
              menu: {
                title: "人员管理",
                icon: "ApartmentOutlined",
                order: 200,
              },
            },
            children: [
              {
                name: "admin.user.index",
                path: "index",
                component: () => import("@/views/admin/base/user/index.vue"),
                meta: {
                  auth: true,
                  title: "用户管理",
                  menu: { title: "用户管理", icon: "HomeOutlined", order: 201 },
                },
              },
              {
                name: "admin.emp.index",
                path: "emp",
                component: () => import("@/views/admin/base/emp/index.vue"),
                meta: {
                  auth: true,
                  title: "员工管理",
                  menu: { title: "员工管理", icon: "HomeOutlined", order: 201 },
                },
              },
            ],
          },
          {
            name: "admin.dept",
            path: "dept",
            component: Translation,
            meta: {
              title: "部门管理",
              menu: { title: "部门管理", icon: "HddOutlined", order: 1000 },
            },
            children: [
              {
                name: "admin.dept.index",
                path: "index",
                component: () => import("@/views/admin/base/dept/index.vue"),
                meta: {
                  title: "部门管理",
                  menu: { title: "部门管理", icon: "HddOutlined", order: 1000 },
                },
              },
            ],
          },
          {
            name: "admin.flow",
            path: "flow",
            component: Translation,
            meta: {
              title: "流程设计",
              menu: { title: "流程设计", icon: "HddOutlined", order: 1000 },
            },
            children: [
              {
                name: "admin.flow.index",
                path: "index",
                component: () => import("@/views/admin/base/flow/index.vue"),
                meta: {
                  title: "流程设计",
                  menu: { title: "流程设计", icon: "HddOutlined", order: 1000 },
                },
              },
              {
                name: "admin.flow.create",
                path: "create",
                component: () => import("@/views/admin/base/flow/create.vue"),
                meta: {
                  title: "创建流程",
                  menu: {
                    title: "创建流程",
                    icon: "HddOutlined",
                    order: 1000,
                    show: false,
                  },
                },
              },
              {
                name: "admin.flow.edit",
                path: ":id/edit",
                component: () => import("@/views/admin/base/flow/edit.vue"),
                meta: {
                  title: "修改流程",
                  menu: {
                    title: "修改流程",
                    icon: "HddOutlined",
                    order: 1000,
                    show: false,
                  },
                },
              },
              {
                name: "admin.flow.initiation",
                path: ":id/initiation",
                component: () =>
                  import("@/views/admin/base/flow/initiation.vue"),
                meta: {
                  title: "发起流程",
                  menu: {
                    title: "发起流程",
                    icon: "HddOutlined",
                    order: 1000,
                    show: false,
                  },
                },
              },
              {
                name: "admin.flow.design",
                path: ":id/design",
                component: () => import("@/views/admin/base/flow/design.vue"),
                meta: {
                  title: "流程设计",
                  menu: {
                    title: "流程设计",
                    icon: "HddOutlined",
                    order: 1000,
                    show: false,
                  },
                },
              },
              {
                name: "admin.flow.entry",
                path: ":flow_id/entry/:entry_id",
                component: () =>
                  import("@/views/admin/base/flow/entrydata.vue"),
                meta: {
                  title: "流程表单",
                  menu: {
                    title: "流程表单",
                    icon: "DashboardOutlined",
                    order: 1000,
                    show: false,
                  },
                },
              },
              {
                name: "admin.flow.proc",
                path: ":flow_id/proc/:entry_id",
                component: () =>
                  import("@/views/admin/base/flow/proc.vue"),
                meta: {
                  title: "审核批复",
                  menu: {
                    title: "审核批复",
                    icon: "DashboardOutlined",
                    order: 1000,
                    show: false,
                  },
                },
              },
            ],
          },
          {
            name: "admin.template",
            path: "template",
            component: Translation,
            meta: {
              title: "模板管理",
              menu: { title: "模板管理", icon: "HddOutlined", order: 1000 },
            },
            children: [
              {
                name: "admin.template.index",
                path: "index",
                component: () =>
                  import("@/views/admin/base/template/index.vue"),
                meta: {
                  title: "模板管理",
                  menu: { title: "模板管理", icon: "HddOutlined", order: 1000 },
                },
              },
              {
                name: "admin.template_form.index",
                path: ":id/template_form",
                component: () =>
                  import("@/views/admin/base/template_form/index.vue"),
                meta: {
                  title: "表单控件",
                  menu: {
                    title: "表单控件",
                    icon: "HddOutlined",
                    order: 1000,
                    show: false,
                  },
                },
              },
            ],
          },
          {
            name: "admin.plugin",
            path: "plugin",
            component: Translation,
            meta: {
              title: "插件管理",
              menu: { title: "插件管理", icon: "HddOutlined", order: 1000,show:false},
              show:false,
            },
            children: [
              {
                name: "admin.plugin.index",
                path: "index",
                component: () =>
                  import("@/views/admin/base/plugin/index.vue"),
                meta: {
                  title: "插件列表",
                  menu: { title: "插件列表", icon: "HddOutlined", order: 1000 },
                },
              },
              {
                name: "admin.plugin.make",
                path: "make",
                component: () =>
                  import("@/views/admin/base/plugin/make.vue"),
                meta: {
                  title: "设计插件",
                  menu: { title: "设计插件", icon: "HddOutlined", order: 1000,show:false },
                },
              },
            ],
          },
        ],
      },
    ],
  },
] as RouteRecordRaw[];
