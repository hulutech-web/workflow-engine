package auth

type Role struct {
	Name        string
	Permissions []Permission
}

type Permission struct {
	Name   string
	Policy string // "allow" or "deny"
}

type RBAC struct {
	roles map[string]Role
}

func NewRBAC() *RBAC {
	return &RBAC{
		roles: make(map[string]Role),
	}
}

func (r *RBAC) AddRole(name string, permissions []Permission) {
	r.roles[name] = Role{
		Name:        name,
		Permissions: permissions,
	}
}

func (r *RBAC) Can(roleName, permission string) bool {
	role, exists := r.roles[roleName]
	if !exists {
		return false
	}

	for _, perm := range role.Permissions {
		if perm.Name == permission {
			return perm.Policy == "allow"
		}
	}

	return false
}
