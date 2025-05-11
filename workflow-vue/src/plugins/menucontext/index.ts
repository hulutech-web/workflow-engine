import '@imengyu/vue3-context-menu/lib/vue3-context-menu.css';
import ContextMenuGlobal, { ContextMenu, ContextMenuGroup, ContextMenuSeparator, ContextMenuItem } from '@imengyu/vue3-context-menu';

const setup = (app: App) => {
    app.component('context-menu', ContextMenu)
    app.component('context-menu-group', ContextMenuGroup)
    app.component('context-menu-separator', ContextMenuSeparator)
    app.component('context-menu-item', ContextMenuItem)
    app.config.globalProperties.$contextmenu = ContextMenuGlobal.showContextMenu;
}

export { setup }
