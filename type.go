package cuckle

import (
	"fmt"
	"strings"
)

type Type string

const (
	TypeAscii     Type = "ascii"
	TypeBigint    Type = "bigint"
	TypeBlob      Type = "blob"
	TypeBoolean   Type = "boolean"
	TypeCounter   Type = "counter"
	TypeDate      Type = "date"
	TypeDecimal   Type = "decimal"
	TypeDouble    Type = "double"
	TypeFloat     Type = "float"
	TypeInet      Type = "inet"
	TypeInt       Type = "int"
	TypeSmallint  Type = "smallint"
	TypeText      Type = "text"
	TypeTime      Type = "time"
	TypeTimestamp Type = "timestamp"
	TypeTimeuuid  Type = "timeuuid"
	TypeTinyint   Type = "tinyint"
	TypeUuid      Type = "uuid"
	TypeVarchar   Type = "varchar"
	TypeVarint    Type = "varint"
)

func TypeList(element Type) Type {
	return Type(fmt.Sprintf("list<%v>", element))
}

func TypeMap(key, value Type) Type {
	return Type(fmt.Sprintf("map<%v, %v>", key, value))
}

func TypeSet(element Type) Type {
	return Type(fmt.Sprintf("set<%v>", element))
}

func TypeTuple(elements ...Type) Type {
	var ss []string

	for _, e := range elements {
		ss = append(ss, string(e))
	}

	return Type(fmt.Sprintf("tuple<%v>", strings.Join(ss, ", ")))
}
