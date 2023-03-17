package model

// Notifier 告警通知人
type Notifier struct {
	ID        int64    `json:"id,string"  gorm:"column:id;primaryKey"` // 表 ID
	Name      string   `json:"name"       gorm:"column:name"`          // 通知人名字
	Events    []string `json:"events"     gorm:"column:events;json"`   // 订阅的事件
	Risks     []string `json:"risks"      gorm:"column:risks;json"`    // 订阅的风险
	Ways      []string `json:"ways"       gorm:"column:ways;json"`     // 通知方式：dong email wechat sms
	Dong      string   `json:"dong"       gorm:"column:dong"`          // 咚咚号
	Email     string   `json:"email"      gorm:"column:email"`         // 邮箱地址
	Mobile    string   `json:"mobile"     gorm:"column:mobile"`        // 手机号
	EventCode []byte   `json:"event_code" gorm:"column:event_code"`    // 规则引擎代码
	RiskCode  []byte   `json:"risk_code"  gorm:"column:risk_code"`     // 规则引擎代码
}

// TableName implement gorm schema.Tabler
func (Notifier) TableName() string {
	return "notifier"
}

type Notifiers []*Notifier

func (ntfs Notifiers) Subscribers() Subscribers {
	events := make(subscriberMap, 32)
	risks := make(subscriberMap, 32)
	for _, ntf := range ntfs {
		for _, evt := range ntf.Events {
			events.put(evt, ntf, ntf.EventCode)
		}
		for _, rsk := range ntf.Risks {
			risks.put(rsk, ntf, ntf.RiskCode)
		}
	}

	return Subscribers{
		event: events,
		risk:  risks,
	}
}

// Devops 走运维平台
type Devops struct {
	Name          string              `json:"name"`           // 用户名
	Mobile        string              `json:"mobile"`         // 手机号, 短信通知与微信通知只能填一个手机号
	NotifyMethods string              `json:"notify_methods"` // 通知方式：strings.Join(methods, ",")
	methods       map[string]struct{} // 通知方式
}

// Subscriber 订阅者信息
type Subscriber struct {
	Dong   []string           // 咚咚通知
	Email  []string           // 邮件通知
	Wechat []string           // 企业微信通知
	SMS    []string           // 短信通知
	Devops []*Devops          // 企业微信通知和短信通知走运维接口
	Code   []byte             // 规则引擎代码
	devops map[string]*Devops // 运维系统按照 mobile 去重
}

func (sub *Subscriber) putDevops(way string, ntf *Notifier) {
	mobile := ntf.Mobile
	ops, ok := sub.devops[mobile]
	if !ok {
		op := &Devops{
			Name:          ntf.Name,
			Mobile:        ntf.Mobile,
			NotifyMethods: way,
			methods:       map[string]struct{}{way: {}},
		}
		sub.devops[mobile] = op
		sub.Devops = append(sub.Devops, op)
		return
	}

	if _, ok = ops.methods[way]; !ok {
		ops.NotifyMethods += "," + way
		ops.methods[way] = struct{}{}
	}
}

type Subscribers struct {
	event subscriberMap
	risk  subscriberMap
}

// Event 获取 event 事件订阅者，找不到返回 nil
func (sub Subscribers) Event(key string) *Subscriber {
	return sub.event[key]
}

// Risk 获取风险事件订阅者，找不到返回 nil
func (sub Subscribers) Risk(key string) *Subscriber {
	return sub.risk[key]
}

type subscriberMap map[string]*Subscriber

func (sbm subscriberMap) put(key string, ntf *Notifier, code []byte) {
	if len(ntf.Ways) == 0 {
		return
	}

	sub, ok := sbm[key]
	if !ok {
		sub = &Subscriber{Code: code, devops: make(map[string]*Devops, 8)}
		sbm[key] = sub
	}

	const dongType, emailType, wechatType, smsType = "dong", "email", "wechat", "sms"
	dong, email, mobile := ntf.Dong, ntf.Email, ntf.Mobile

	for _, way := range ntf.Ways {
		switch way {
		case dongType:
			if dong != "" {
				sub.Dong = append(sub.Dong, dong)
			}
		case emailType:
			if email != "" {
				sub.Email = append(sub.Email, email)
			}
		case wechatType:
			if mobile != "" {
				sub.Wechat = append(sub.Wechat, mobile)
				sub.putDevops(wechatType, ntf)
			}
		case smsType:
			if mobile != "" {
				sub.SMS = append(sub.SMS, mobile)
				sub.putDevops(smsType, ntf)
			}
		}
	}
}
