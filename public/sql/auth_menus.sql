INSERT INTO `auth_menus` (`id`, `pid`, `name`, `path`, `redirect`, `componentPath`, `isDisabled`, `icon`, `menuType`, `title`, `requiresAuth`,  `keepAlive`, `hide`, `href`, `activeMenu`, `withoutTab`, `pinTab`, `sort`) VALUES
    (1, 0, 'workbench', '/workbench', '', '/dashboard/workbench/index.vue',  0, 'icon-park-outline:alarm', 'page', '工作台', 1,  0, 0, NULL, NULL, 0, 1, 0),
    (2,  0, 'system', '/system', NULL, NULL,  0, 'icon-park-outline:setting', 'dir', '系统设置', 1,  0, 0, NULL, NULL, 0, 0, 99),
    (3,  2, 'system_menu', '/system/menu', '', '/setting/menu/index.vue',  0, 'icon-park-outline:application-menu', 'page', '系统菜单', 1,  0, 0, '', '', 0, 0, 1),
    (4, 0, 'org', '/org', '', '',  0, 'icon-park-outline:broadcast-one', 'dir', '组织管理', 1,  0, 0, '', '', 1, 0, 98),
    (5,  4, 'org_dept', '/org/dept', '', '/demo/map/index.vue',  0, 'carbon:category', 'page', '部门管理', 1,  0, 0, '', '', 1, 0, 0);
