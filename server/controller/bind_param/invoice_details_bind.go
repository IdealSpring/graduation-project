package bind_param

type InvoiceDetailsParam struct {
	DetailsId string `json:"detailsId"`

	// 分页信息
	Current int `json:"current"`
	Size    int `json:"size"`
}
