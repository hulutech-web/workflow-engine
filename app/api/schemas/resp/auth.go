package resp

import "github.com/dromara/carbon/v2"

type TenantResp struct {
	ID          uint            `json:"id" structs:"id"`
	Name        string          `json:"name" structs:"name"`
	Address     string          `json:"address" structs:"address"`
	Phone       string          `json:"phone" structs:"phone"`
	Email       string          `json:"email" structs:"email"`
	Domain      string          `json:"domain" structs:"domain"`
	Logo        string          `json:"logo" structs:"logo"`
	Description string          `json:"description" structs:"description"`
	IsDisable   uint8           `json:"is_disable" structs:"is_disable"`
	ExpiredAt   int64           `json:"expired_at" structs:"expired_at"`
	CreatedAt   carbon.DateTime `json:"created_at" structs:"created_at"`
	UpdatedAt   carbon.DateTime `json:"updated_at" structs:"updated_at"`
}

type UserResp struct {
	ID           uint   `json:"id" structs:"id"`
	Username     string `json:"username" structs:"username"`
	Phone        string `json:"phone" structs:"phone"`
	Email        string `json:"email" structs:"email"`
	Avatar       string `json:"avatar" structs:"avatar"`
	Nickname     string `json:"nickname" structs:"nickname"`
	IsMultipoint uint8  `json:"is_multipoint" structs:"is_multipoint"`
	IsDisable    uint8  `json:"is_disable" structs:"is_disable"`
	Role         struct {
		ID      uint   `json:"id" structs:"id"`
		Name    string `json:"name" structs:"name"`
		IsAdmin uint8  `json:"is_admin" structs:"is_admin"`
	} `json:"role" structs:"role"`
	Tenant struct {
		ID   uint   `json:"id" structs:"id"`
		Name string `json:"name" structs:"name"`
	} `json:"tenant" structs:"tenant"`
	CreatedAt carbon.DateTime `json:"created_at" structs:"created_at"`
	UpdatedAt carbon.DateTime `json:"updated_at" structs:"updated_at"`
}

type UserSelfResp struct {
	Permissions []string `json:"permissions" structs:"permissions"`
	User        UserResp `json:"user" structs:"user"`
}

type MenuResp struct {
	ID            uint            `json:"id" structs:"id"`
	Pid           uint            `json:"pid" structs:"pid"`
	Name          string          `json:"name" structs:"name"`
	Path          string          `json:"path" structs:"path"`
	Redirect      string          `json:"redirect" structs:"redirect"`
	ComponentPath string          `json:"componentPath" structs:"componentPath"`
	Title         string          `json:"title" structs:"title"`               // 页面标题
	Icon          string          `json:"icon" structs:"icon"`                 // 图标
	RequiresAuth  bool            `json:"requiresAuth" structs:"requiresAuth"` // 是否需要登录权限
	KeepAlive     bool            `json:"keepAlive" structs:"keepAlive"`       // 是否开启页面缓存
	Hide          bool            `json:"hide" structs:"hide"`                 // 不显示在菜单中
	Sort          uint            `json:"sort" structs:"sort"`                 // 菜单排序
	Href          string          `json:"href" structs:"href"`                 // 嵌套外链
	ActiveMenu    string          `json:"activeMenu" structs:"activeMenu"`     // 当前路由高亮菜单
	WithoutTab    bool            `json:"withoutTab" structs:"withoutTab"`     // 不添加到Tab
	PinTab        bool            `json:"pinTab" structs:"pinTab"`             // 固定Tab
	MenuType      string          `json:"menuType" structs:"menuType"`         // dir or page
	Button        []MenuButton    `json:"button" structs:"button"`
	CreatedAt     carbon.DateTime `json:"created_at" structs:"created_at"`
	UpdatedAt     carbon.DateTime `json:"updated_at" structs:"updated_at"`
}

type MenuButton struct {
	Title string `json:"title" structs:"title"`
	Name  string `json:"name" structs:"name"`
	Sort  uint   `json:"sort" structs:"sort"`
}

// RoleSimpleResp 系统角色返回简单信息
type RoleSimpleResp struct {
	ID        uint            `json:"id" structs:"id"`     // 主键
	Name      string          `json:"name" structs:"name"` // 角色名称
	CreatedAt carbon.DateTime `json:"created_at" structs:"created_at"`
	UpdatedAt carbon.DateTime `json:"updated_at" structs:"updated_at"`
}

// RoleResp 系统角色返回信息
type RoleResp struct {
	ID        uint            `json:"id" structs:"id"`                 // 主键
	Name      string          `json:"name" structs:"name"`             // 角色名称
	Remark    string          `json:"remark" structs:"remark"`         // 角色备注
	Menus     []uint          `json:"menus" structs:"menus"`           // 关联菜单
	Member    int64           `json:"member" structs:"member"`         // 成员数量
	Sort      uint16          `json:"sort" structs:"sort"`             // 角色排序
	IsDisable uint8           `json:"is_disable" structs:"is_disable"` // 是否禁用: [0=否, 1=是]
	CreatedAt carbon.DateTime `json:"created_at" structs:"created_at"`
	UpdatedAt carbon.DateTime `json:"updated_at" structs:"updated_at"`
}
