package bind_param

type RoleInfoParam struct {
	// 角色信息
	RoleId   string `json:"roleId"`
	RoleName string `json:"roleName"`
	Desc     string `json:"desc"`

	// 权限信息
	PermType int      `json:"permType"`
	Values   []string `json:"vals"`

	// 分页信息
	Current int `json:"current"`
	Size    int `json:"size"`
}
type RoleInfoParam2 struct {
	// 角色信息
	RoleId   int    `json:"roleId"`
	RoleName string `json:"roleName"`
	Desc     string `json:"desc"`

	// 权限信息
	PermType int      `json:"permType"`
	Values   []string `json:"vals"`

	// 分页信息
	Current int `json:"current"`
	Size    int `json:"size"`
}
