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
	Menus       []uint          `json:"menus,omitempty" structs:"menus"`
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
	ID         uint         `json:"id" structs:"id"`
	Title      string       `json:"title" structs:"title"`
	Pid        uint         `json:"pid" structs:"pid"`
	Name       string       `json:"name" structs:"name"`
	Path       string       `json:"path" structs:"path"`
	Component  string       `json:"component" structs:"component"`
	Icon       string       `json:"icon" structs:"icon"`
	MenuType   string       `json:"menu_type" structs:"menu_type"`
	Cacheable  bool         `json:"cacheable" structs:"cacheable"`
	RenderMenu bool         `json:"render_menu" structs:"render_menu"`
	Permission string       `json:"permission" structs:"permission"`
	Sort       uint16       `json:"sort" structs:"sort"`
	Target     string       `json:"target" structs:"target"`
	Badge      string       `json:"badge" structs:"badge"`
	Button     []MenuButton `json:"button" structs:"button"`
}

type MenuButton struct {
	ID    uint   `json:"id" structs:"id"`
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
