package cuckle

import (
	"encoding/hex"
	"fmt"
)

type Constant string

func ConstantBoolean(b bool) Constant {
	return Constant(fmt.Sprint(b))
}

func ConstantInteger(i int64) Constant {
	return Constant(fmt.Sprint(i))
}

func ConstantFloat(f float64) Constant {
	return Constant(fmt.Sprint(f))
}

func ConstantHex(b []byte) Constant {
	return Constant(fmt.Sprintf("0x%v", hex.EncodeToString(b)))
}

func ConstantString(s string) Constant {
	return Constant(fmt.Sprintf("'%v'", s))
}

func ConstantStringEscaped(s string) Constant {
	return Constant(fmt.Sprintf("$$%v$$", s))
}

func ConstantUUID(s string) Constant {
	return Constant(fmt.Sprint(s))
}
