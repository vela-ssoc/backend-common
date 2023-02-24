package pubrr

type ErrorResult struct {
	Code  int    `json:"code"`
	Node  string `json:"node"`
	Cause string `json:"cause"`
}
