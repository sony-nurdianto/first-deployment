package serde

import "encoding/json"

type Serde interface {
	Marshal(v any) ([]byte, error)
}

type serde struct{}

func NewSerde() *serde {
	return &serde{}
}

func (s *serde) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}
