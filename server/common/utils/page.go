package utils

// 分页相关
type Page struct {
	// 数据
	Records interface{}	`json:"records"`
	// 当前页
	Current int	`json:"current"`
	// 总页数
	Pages int	`json:"pages"`
	// 每页显示条数
	Size int	`json:"size"`
	// 总数
	Total int	`json:"total"`
}
