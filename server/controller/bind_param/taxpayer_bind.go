package bind_param

type TaxpayerParam struct {
	TaxpayerId string `json:"taxpayerId"`

	// 分页信息
	Current int `json:"current"`
	Size    int `json:"size"`
}
