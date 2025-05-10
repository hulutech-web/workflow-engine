package types

type admin struct {
	BackstageManageKey       string
	BackstageRolesKey        string
	BackstageTenantsKey      string
	BackstageTokenKey        string
	BackstageTokenSet        string
	BackstageTokenExpireTime int
	NotLoginUri              []string
	NotAuthUri               []string
	ShowWhitelistUri         []string
	CommonUri                []string
}

var Admin = admin{
	// 管理缓存键
	BackstageManageKey: "backstage:manage",
	// 角色缓存键
	BackstageRolesKey: "backstage:roles",
	// 租户缓存键
	BackstageTenantsKey: "backstage:tenants",
	// 令牌缓存键
	BackstageTokenKey: "backstage:token:",
	// 令牌的集合
	BackstageTokenSet: "backstage:token:set:",
	// 令牌过期时间
	BackstageTokenExpireTime: 86400,
	// 未登录的URI
	NotLoginUri: []string{
		"/login",
		"/logout",
		"/register",
	},
	// 未授权的URI
	NotAuthUri: []string{},
	// 白名单URI
	CommonUri: []string{},
}
