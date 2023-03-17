package model

import (
	"path/filepath"
	"time"
)

// Startup minion 节点启动脚本参数配置
type Startup struct {
	ID        int64          `json:"id,string"  gorm:"column:id;primaryKey"`                                // 取自 Minion.ID
	Node      startupNode    `json:"node"       gorm:"column:node"`                                         // 节点配置项
	Logger    startupLogger  `json:"logger"     gorm:"column:logger"`                                       // 日志配置项
	Console   startupConsole `json:"console"    gorm:"column:console"`                                      // 控制台输出配置项
	Extends   startupExtends `json:"extends"    gorm:"column:extends"    validate:"omitempty,lte=100,dive"` // 其它扩展参数
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`                                   // 创建时间
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`                                   // 更新时间
}

// TableName implement gorm schema.Tabler
func (Startup) TableName() string {
	return "startup"
}

// Parse 检查处理格式化配置
func (s Startup) Parse() (Startup, error) {
	ret := Startup{
		ID: s.Node.ID, Node: s.Node, Logger: s.Logger, Console: s.Console,
		CreatedAt: s.CreatedAt, UpdatedAt: s.UpdatedAt,
	}

	ret.Node.Prefix = filepath.Clean(ret.Node.Prefix)
	ret.Logger.Filename = filepath.Clean(ret.Logger.Filename)
	ret.Console.Script = filepath.Clean(ret.Console.Script)

	// 处理 extend 字段
	extends := make(startupExtends, len(s.Extends))
	for i, extend := range s.Extends {
		ext, err := extend.parse()
		if err != nil {
			return Startup{}, err
		}
		extends[i] = ext
	}
	ret.Extends = extends

	return ret, nil
}
