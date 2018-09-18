package bqdata

import (
	"fmt"
)

type ColumnType string

const (
	DATETIME ColumnType = "DATETIME"
	STRING   ColumnType = "STRING"
	INTEGER  ColumnType = "INTEGER"
	FLOAT    ColumnType = "FLOAT"
)

func (columnType ColumnType) string() string {
	return string(columnType)
}

func columnType(str string) (ColumnType, error) {
	switch str {

	case DATETIME.string():
		return DATETIME, nil
	case STRING.string():
		return STRING, nil
	case INTEGER.string():
		return INTEGER, nil
	case FLOAT.string():
		return FLOAT, nil
	default:
		return STRING, fmt.Errorf("error: %s is not found in ColumnType", str)
	}
}
