package opdata

import "encoding/json"

type BrokerReceive struct {
	MinionIDs []int64         `json:"minion_ids"` // 相关节点
	Original  json.RawMessage `json:"original"`   // 原始信息
}

type ManagerSend struct {
	MinionIDs []int64 `json:"minion_ids"`
	Original  any     `json:"original"`
}
