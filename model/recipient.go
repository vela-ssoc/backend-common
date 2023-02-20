package model

import "time"

// Recipient 事件通知订阅者
type Recipient struct {
	ID         int64     `json:"id,string"   gorm:"column:id;primaryKey"` // ID
	Name       string    `json:"name"        gorm:"column:name"`          // 接收者名字
	SendDong   bool      `json:"send_dong"   gorm:"column:send_dong"`     // 是否发送咚咚通知
	Dong       string    `json:"dong"        gorm:"column:dong"`          // 咚咚号
	SendEmail  bool      `json:"send_email"  gorm:"column:send_email"`    // 是否发送邮件通知
	Email      string    `json:"email"       gorm:"column:email"`         // 邮件地址
	SendWechat bool      `json:"send_wechat" gorm:"column:send_wechat"`   // 是否发送微信通知
	Wechat     string    `json:"wechat"      gorm:"column:wechat"`        // 微信号
	SendSMS    bool      `json:"send_sms"    gorm:"column:send_sms"`      // 是否发送短信通知
	SMS        string    `json:"sms"         gorm:"column:sms"`           // 短信通知号码
	SendPhone  bool      `json:"send_phone"  gorm:"column:send_phone"`    // 是否发送电话通知
	Phone      string    `json:"phone"       gorm:"column:phone"`         // 电话通知号码
	Events     []string  `json:"events"      gorm:"column:events;json"`   // 订阅的事件类型
	Risks      []string  `json:"risks"       gorm:"column:risks;json"`    // 订阅的风险通知类型
	EventCode  []byte    `json:"event_code"  gorm:"column:event_code"`    // 事件过滤规则代码
	RiskCode   []byte    `json:"risk_code"   gorm:"column:risk_code"`     // 风险过滤规则代码
	CreatedAt  time.Time `json:"created_at"  gorm:"column:created_at"`    // 创建时间
	UpdatedAt  time.Time `json:"updated_at"  gorm:"column:updated_at"`    // 修改时间
}

// TableName implemented gorm schema.Tabler
func (Recipient) TableName() string {
	return "recipient"
}

// RecipientMap key: event/risk
type RecipientMap map[string][]*Recipient

type Recipients []*Recipient

// Classify 根据订阅主题对用户分类
func (rts Recipients) Classify() (events, risks RecipientMap) {
	events, risks = make(RecipientMap, 32), make(RecipientMap, 32)
	evtTemp := make(map[string]map[int64]struct{}, 32)
	rskTemp := make(map[string]map[int64]struct{}, 32)

	for _, rec := range rts {
		for _, evt := range rec.Events {
			var exist bool
			uid := rec.ID
			if em, ok := evtTemp[evt]; !ok {
				evtTemp[evt] = map[int64]struct{}{uid: {}}
			} else {
				if _, exist = em[uid]; !exist {
					em[uid] = struct{}{}
				}
			}
			if !exist {
				rs := events[evt]
				events[evt] = append(rs, rec)
			}
		}
		for _, rsk := range rec.Risks {
			var exist bool
			uid := rec.ID
			if rm, ok := rskTemp[rsk]; !ok {
				rskTemp[rsk] = map[int64]struct{}{uid: {}}
			} else {
				if _, exist = rm[uid]; !exist {
					rm[uid] = struct{}{}
				}
			}
			if !exist {
				rs := risks[rsk]
				risks[rsk] = append(rs, rec)
			}
		}
	}

	return
}
