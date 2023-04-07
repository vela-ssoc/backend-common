package opdata

// EdictSubstanceEvent 告知 broker 节点这些 agent 的配置发生了变动。
type EdictSubstanceEvent struct {
	MinionID []int64 `json:"minion_id"` // 波及的 minion(agent) 节点 ID
}
