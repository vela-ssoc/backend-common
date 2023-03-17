package model

import "database/sql/driver"

// 告警分级参考:
// https://docs.oracle.com/cd/E86948_01/docs.465/UAM_UIMs/concepts/c_system_alarm_levels_maintmanual.html
const (
	ELvlCritical EventLevel = 100000
	ELvlMajor    EventLevel = 10000
	ELvlMinor    EventLevel = 1000
	ELvlNote     EventLevel = 100
)

var (
	evtLvlSI = map[string]EventLevel{
		"紧急": ELvlCritical,
		"重要": ELvlMajor,
		"次要": ELvlMinor,
		"普通": ELvlNote,
	}

	evtLvlIS = map[EventLevel]string{
		ELvlCritical: "紧急",
		ELvlMajor:    "重要",
		ELvlMinor:    "次要",
		ELvlNote:     "普通",
	}
)

type EventLevel int

func (el EventLevel) Value() (driver.Value, error) {
	str := evtLvlIS[el]
	return []byte(str), nil
}

func (el *EventLevel) Scan(src any) error {
	if bs, ok := src.([]byte); ok {
		str := string(bs)
		lvl := evtLvlSI[str]
		*el = lvl
	}
	return nil
}

func (el *EventLevel) UnmarshalText(bs []byte) error {
	str := string(bs)
	lvl := evtLvlSI[str]
	*el = lvl
	return nil
}

func (el EventLevel) MarshalText() ([]byte, error) {
	str := evtLvlIS[el]
	return []byte(str), nil
}

func (rl EventLevel) String() string {
	return evtLvlIS[rl]
}
