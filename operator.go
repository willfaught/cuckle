package cuckle

type Operator string

const (
	OperatorAdd          Operator = "+"
	OperatorContains     Operator = "contains"
	OperatorContainsKey  Operator = "contains key"
	OperatorEqual        Operator = "="
	OperatorGreater      Operator = ">"
	OperatorGreaterEqual Operator = ">="
	OperatorIn           Operator = "in"
	OperatorLess         Operator = "<"
	OperatorLessEqual    Operator = "<="
	OperatorNotEqual     Operator = "!="
	OperatorSubtract     Operator = "-"
)
