package opencond

import (
	"encoding/json"
	"net/url"
)

type environment struct {
	Key      string `json:"key"      validate:"lte=30"`
	Operator string `json:"operator" validate:"oneof=eq ne gt lt gte lte in notin like notlike"`
	Value    string `json:"value"    validate:"lte=200"`
}

type Environments []*environment

func (env *environment) UnmarshalBind(raw string) error {
	data, err := url.PathUnescape(raw)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), env)
}

type Environment struct {
	Group string `json:"group" validate:"lte=20"`
}

type field struct {
	Key      string `json:"key"      validate:"required,lte=30"`
	Operator string `json:"operator" validate:"oneof=eq ne gt lt gte lte in notin like notlike"`
	Value    string `json:"value"    validate:"required,lte=200"`
}

type Fields []*field
