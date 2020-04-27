package bind_param

type LogParam struct {
	IsLoginOrLogout bool   `json:"isLoginOrLogout"`
	ModuleName      string `json:"moduleName"`
	Username        string `json:"username"`
	Nick            string `json:"nick"`

	// 分页信息
	Current int `json:"current"`
	Size    int `json:"size"`
}
