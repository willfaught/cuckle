package cuckle

import (
	"encoding/hex"
	"fmt"
)

// Constant is a scalar literal.
type Constant string

// ConstantBoolean returns a Constant for b.
func ConstantBoolean(b bool) Constant {
	return Constant(fmt.Sprint(b))
}

// ConstantInteger returns a Constant for i.
func ConstantInteger(i int64) Constant {
	return Constant(fmt.Sprint(i))
}

// ConstantFloat returns a Constant for f.
func ConstantFloat(f float64) Constant {
	return Constant(fmt.Sprint(f))
}

// ConstantHex returns a Constant for b prefixed by 0x. Two characters represent
// each byte.
func ConstantHex(b []byte) Constant {
	return Constant(fmt.Sprintf("0x%v", hex.EncodeToString(b)))
}

// ConstantString returns a Constant for s.
func ConstantString(s string) Constant {
	return Constant(fmt.Sprintf("'%v'", s))
}

// ConstantStringEscaped returns a Constant for s.
func ConstantStringEscaped(s string) Constant {
	return Constant(fmt.Sprintf("$$%v$$", s))
}

// ConstantUUID returns a Constant for uuid.
func ConstantUUID(s string) Constant {
	return Constant(fmt.Sprint(s))
}
