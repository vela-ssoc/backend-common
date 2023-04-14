package opdata

import "encoding/json"

type Mtob struct {
	Method    string  `json:"method"`
	Path      string  `json:"path"`
	MinionIDs []int64 `json:"minion_ids"`
	Data      any     `json:"data"`
}

type RouteToAgentReceive struct {
	Method    string
	Path      string
	MinionIDs []int64         `json:"minion_ids"`
	Data      json.RawMessage `json:"data"`
}
