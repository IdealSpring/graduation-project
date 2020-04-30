package bind_param

type InvoiceParam struct {
	InvoiceId string `json:"invoiceId"`

	// 分页信息
	Current int `json:"current"`
	Size    int `json:"size"`
}
