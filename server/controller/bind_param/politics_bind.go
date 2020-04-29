package bind_param

type PoliticsParam struct {
	PoliticsId    int    `json:"politicsId"`
	PoliticsType  string `json:"politicsType"` // guide、report、notify
	Title         string `json:"title"`
	Abstract      string `json:"abstract"`
	Content       string `json:"content"`
	EnclosurePath string `json:"enclosurePath"`
	ReadingSigns  string `json:"readingSigns"`

	// 分页信息
	Current int `json:"current"`
	Size    int `json:"size"`

	// 用户信息
	UserId int `json:"userId"`
}
