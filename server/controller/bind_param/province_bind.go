package bind_param

type ProvinceParam struct {
	// 省份信息
	ProvinceName string `json:"provinceName"`

	// 分页信息
	Current int `json:"current"`
	Size    int `json:"size"`
}
