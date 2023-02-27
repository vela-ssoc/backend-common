package pubody

// BizError 业务错误结构体
type BizError struct {
	Code  int    `json:"code"`  // 错误码
	Node  string `json:"node"`  // 报错节点
	Cause string `json:"cause"` // 错误原因
}
