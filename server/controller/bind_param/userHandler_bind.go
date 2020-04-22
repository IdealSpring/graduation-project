package bind_param

type PostQueryBindParam struct {
	Nick    string `json:"nick"`
	Current int    `json:"current"`
	Size    int    `json:"size"`
}

type PostAddUserBindParam struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Nick     string `json:"nick"`
	RoleId   int    `json:"roleId"`
	Pwd      string `json:"pwd"`
	Pwd2     string `json:"pwd2"`
}

type UserInfoParam struct {
	UserId   int    `json:"userId"`
	RoleId   int    `json:"roleId"`
}
