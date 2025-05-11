<!--菜单子组件-->
<script lang='ts'>
import { defineComponent } from 'vue'
import SubMenu from '@/layouts/admin/SubMenu.vue'
import admin_routes from '@/routes/admin'
import useMenuStore from '@/store/useMenuStore'
const storage = useStorage()
const menuStore = useMenuStore()
export default defineComponent({
  components: {
    SubMenu,
  },
  setup(props, { expose }) {
    const user = ref({})
    user.value = JSON.parse(storage.get(CacheEnum.USER))
 
    const new_rts = ref([])
    const currentRoute = ref({})
    // 深度递归，判断meta中的menu中是否存在show,如果存在且值为false,则删除该路由
    const deep = (routes: RouteRecordRaw[]) => {
      //1、标记需要删除的路由
      let shouldDelsRaws: RouteRecordRaw[] = []
      routes.forEach((route) => {
        if (route.meta.menu?.isSuper == true) {
          shouldDelsRaws.push(route)
        }
        // @ts-ignore
        if (route.meta.menu?.show === false) {
          shouldDelsRaws.push(route)
        }

        if (route.children) {
          deep(route.children)
        }
      })
      // console.log(shouldDelsRaws)
      shouldDelsRaws.forEach((shouldDel) => {
        const index = routes.indexOf(shouldDel)
        routes.splice(index, 1)
      })
      return routes
    }


    const initRoutes = () => {
      //每次运行是都复制一份新的admin_routes[0].children
      let cloneRoutes = admin_routes[0].children
      // 获取路由
      const topRoutes = cloneRoutes
        .filter(tr => tr.path.split('/').length == 4)
        .filter(r => r.children.length == 0)
        .filter(r => r.path.includes('admin'))

      let deepedRoutes = deep([...cloneRoutes])
      let mergeRoutes = [...topRoutes, ...deepedRoutes]

      // 1、调用一下menuStore将数据加入进去

      menuStore.menus = mergeRoutes

      new_rts.value = mergeRoutes
      // 2、当页面一刷新时，获取当前的路由

      currentRoute.value = useRoute()

      menuStore.selectedKeys = [currentRoute.value.name]
      initOpenKeys()

    }

    // 辅助方法

    const makeOpenKeys = (pNameArr: []) => {
      // admin.config.equipment
      let resultArr: any = [];
      // 先去头2个数组拼接
      let headStr = pNameArr.slice(0, 2).join(".")
      resultArr.push(headStr)
      for (let j = 0; j < pNameArr.length - 2; j++) {
        headStr += "." + pNameArr[j + 2]
        resultArr.push(headStr)
      }
      return resultArr;
    }


    // 页面一加载初始化路由，并展开对应的菜单

    const initOpenKeys = () => {
      let curRoute = currentRoute.value.name
      if (curRoute.split('.').length == 2) {

        menuStore.openKeys = [curRoute]
      }

      if (curRoute.split('.').length > 2) {

        const pNameArr = curRoute.split(".")

        let resultNames = makeOpenKeys(pNameArr)

        //改变菜单的折叠
        menuStore.openKeys = resultNames
      }
    }
    // 3、设置展开的问题

    const routeTo = (obj: object) => {
      menuStore.routeTo(obj["key"])
    }

    // onOpenChange只控制一级菜单，点击一级菜单时触发，只取最后一个，始终显示折叠的菜单项
    const onOpenChange = (openKeys: string[]) => {
      // console.log(openKeys)
      if (openKeys.length > 0) {
        const lastKey = openKeys[openKeys.length - 1];
        if (lastKey.split('.').length == 2) {
          menuStore.openKeys = [lastKey]
          return
        }

        if (lastKey.split('.').length > 2) {

          const pNameArr = lastKey.split(".")
          let resultNames = makeOpenKeys(pNameArr)
          //改变菜单的折叠
          menuStore.openKeys = resultNames
          return
        }
      }
    }
    //每次页面不同，需要显示不同的路由
    onMounted(() => {
      initRoutes()
    });
    // Expose导出active，供父级调用，控制菜单系统文字的显示
    const active = ref(true)
    expose({ active })

    return { active, routeTo, new_rts, onOpenChange, menuStore,  user }
  },
})
</script>
<!--菜单子组件END-->

<template>
  <div>
    <div class='logo'>
      <img src="@/assets/logo/charging.png" style="width:44px;" alt="">
      <span :class="active == true ? 'block' : 'hidden'" class="sys-title">goravel-workflow</span>
    </div>
    <a-menu mode='inline' theme='dark' @click='routeTo'  v-model:open-keys="menuStore.openKeys"
      v-model:selectedKeys="menuStore.selectedKeys">
      <template v-for='menu in new_rts' :key='menu.name'>
        <template v-if='menu.children && menu.children.length === 0'>
          <a-menu-item :key='menu.name'>
            <template #icon>
              <component
                :is="menu.meta?(menu.meta.menu?(menu.meta.menu['icon'] as string ?menu.meta.menu['icon'] :'AppstoreAddOutlined'):'AppstoreAddOutlined'):'AppstoreAddOutlined'" />
            </template>
            {{ menu.meta.menu['title'] }}
          </a-menu-item>
        </template>
        <template v-else>
          <sub-menu :menu-info='menu' />
        </template>
      </template>
    </a-menu>
  </div>
</template>

<style lang='scss'>
.logo {
  color: rgb(249, 229, 2);
  box-sizing: border-box;
  height: 48px;
  margin: 16px;
  font-weight: bolder;
  font-size: 14px;
  display: flex;
  justify-content: space-evenly;
  align-items: center;

  .sys-title {
    // display:none;时2秒钟
    transition: all 0.2s ease-in-out;
  }
}
</style>
