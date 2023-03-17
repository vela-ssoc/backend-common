package model

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type Int64s []int64

func (is Int64s) MarshalJSON() (text []byte, err error) {
	size := len(is)
	str := make([]string, 0, size)
	for _, i := range is {
		s := strconv.FormatInt(i, 10)
		str = append(str, s)
	}

	return json.Marshal(str)
}

func (is *Int64s) UnmarshalJSON(raw []byte) error {
	var str []string
	if err := json.Unmarshal(raw, &str); err != nil {
		return err
	}

	dats := make([]int64, 0, len(str))
	for _, st := range str {
		num, err := strconv.ParseInt(st, 10, 64)
		if err != nil {
			return err
		}
		dats = append(dats, num)
	}
	*is = dats

	return nil
}

func (is Int64s) Len() int           { return len(is) }
func (is Int64s) Less(i, j int) bool { return is[i] < is[j] }
func (is Int64s) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }

// startupNode 节点配置项
type startupNode struct {
	ID     int64  `json:"id,string" validate:"required"`
	DNS    string `json:"dns"       validate:"ip"`
	Prefix string `json:"prefix"`
}

// Scan implement std sql sql.Scanner
func (s *startupNode) Scan(v any) error {
	return json.Unmarshal(v.([]byte), s)
}

// Value implement std sql driver.Valuer
func (s startupNode) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type startupLogger struct {
	Level    string `json:"level"    validate:"oneof=debug info error"`
	Filename string `json:"filename" validate:"required"`
	Console  bool   `json:"console"`
	Format   string `json:"format"   validate:"oneof=text json"`
	Caller   bool   `json:"caller"`
	Skip     int    `json:"skip"     validate:"gt=-10,lt=10"`
}

// Scan implement std sql sql.Scanner
func (s *startupLogger) Scan(v any) error {
	return json.Unmarshal(v.([]byte), s)
}

// Value implement std sql driver.Valuer
func (s startupLogger) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type startupConsole struct {
	Enable  bool   `json:"enable"`
	Network string `json:"network" validate:"required_if=Enable true,omitempty,oneof=tcp udp unix"`
	Address string `json:"address" validate:"required_if=Enable true,omitempty,hostname_port"`
	Script  string `json:"script"  validate:"required_if=Enable true"`
}

// Scan implement std sql sql.Scanner
func (s *startupConsole) Scan(v any) error {
	return json.Unmarshal(v.([]byte), s)
}

// Value implement std sql driver.Valuer
func (s startupConsole) Value() (driver.Value, error) {
	return json.Marshal(s)
}

const (
	extNumberType         extFieldType = "number"          // 数字类型
	extBoolType           extFieldType = "bool"            // 布尔类型
	extStringType         extFieldType = "string"          // 字符串类型
	extReferType          extFieldType = "ref"             // 引用类型
	extReadonlyStringType extFieldType = "readonly_string" // 只读字符串类型
)

// extFieldType startup 配置项扩展字段类型
type extFieldType string

// startupExtend 扩展字段
type startupExtend struct {
	Name  string `json:"name"  validate:"required"`                                     // 名字
	Type  string `json:"type"  validate:"oneof=number bool string ref string_readonly"` // 类型
	Value string `json:"value" validate:"required"`                                     // 值
}

// parseField 校验单个一项数据
func (s startupExtend) parse() (*startupExtend, error) {
	ret := &startupExtend{Name: s.Name, Type: s.Type, Value: s.Value}
	switch extFieldType(s.Type) {
	case extNumberType:
		if _, err := strconv.ParseFloat(s.Value, 64); err != nil {
			return nil, err
		}
	case extBoolType:
		tf, err := strconv.ParseBool(ret.Value)
		if err != nil {
			return nil, err
		}
		ret.Value = strconv.FormatBool(tf)
	case extStringType:
		ret.Value = s.escape()
	case extReferType:
	case extReadonlyStringType:
		ret.Value = "r" + s.escape()
	default:
		return nil, errors.New("不支持的数据类型: " + s.Type)
	}

	return ret, nil
}

// escape 对字符串中的字符转译：
// 例如：
//
//	a"bc"d --> a\"bc\"d
//	\r\n   --> \\r\\n
func (s startupExtend) escape() string {
	var buf bytes.Buffer
	buf.WriteByte('"')
	for _, char := range s.Value {
		switch char {
		case '"', '\\':
			buf.WriteByte('\\')
		}
		buf.WriteRune(char)
	}
	buf.WriteByte('"')
	return buf.String()
}

// startupExtends 配置扩展项
type startupExtends []*startupExtend

// Scan implement std sql sql.Scanner
func (s *startupExtends) Scan(v any) error {
	return json.Unmarshal(v.([]byte), s)
}

// Value implement std sql driver.Valuer
func (s startupExtends) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// TaskRunner 内部运行状态
type TaskRunner struct {
	Name   string `json:"name"`   // 内部服务名字
	Type   string `json:"type"`   // 类型
	Status string `json:"status"` // 状态
}

// Semver https://semver.org/lang/zh-CN/
type Semver string

// Int64 计算版本号
func (sv Semver) Int64() int64 {
	sn := strings.SplitN(string(sv), ".", 3)
	if len(sn) != 3 {
		return 0
	}

	sp := strings.SplitN(sn[2], "-", 2)
	sn[2] = sp[0]

	var ret int64
	for _, s := range sn {
		num, _ := strconv.ParseInt(s, 10, 64)
		ret *= 1000000
		ret += num
	}

	return ret
}
