package bqdata

import (
	"fmt"
)

type Mode string

type SchemaRowJson struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Mode string `json:"mode"`
}

type Schema struct {
	Name string
	Type ColumnType
	Mode Mode
}

const (
	Nullable Mode = "nullable"
	Required Mode = "required"
)

func (mode Mode) string() string {
	return string(mode)
}

func mode(str string) (Mode, error) {
	switch str {
	case Nullable.string():
		return Nullable, nil
	case Required.string():
		return Required, nil
	default:
		return Nullable, fmt.Errorf("error: %s is not found in Mode", str)
	}
}
