package cuckle

import (
	"fmt"
	"strings"
)

const (
	OptionAllowFiltering  = "allow filtering"
	OptionCalled          = "called"
	OptionClusteringOrder = "clustering order"
	OptionCompactStorage  = "compact storage"
	OptionDistinct        = "distinct"
	OptionIfExists        = "if exists"
	OptionIfNotExists     = "if not exists"
	OptionJSON            = "json"
	OptionLimit           = "limit"
	OptionOptions         = "options"
	OptionOrderBy         = "order by"
	OptionProperties      = "properties"
	OptionReplace         = "replace"
	OptionUsing           = "using"
)

const (
	KeyspaceDurableWrites = "durable writes"
	KeyspaceReplication   = "replication"
)

func QueryAggregateCreate(keyspace, aggregate, stateFunc, finalFunc Identifier, parameters []Type, stateType Type, init Term, options ...Option) string {
	var m = optionMap(options)
	var q []string

	if optionHas(m, OptionReplace) {
		q = append(q, "replace")
	} else {
		q = append(q, "create")
	}

	q = append(q, "aggregate")

	if optionHas(m, OptionIfNotExists) {
		q = append(q, "if not exists")
	}

	var sig = fmt.Sprintf("%v.%v(%v)", keyspace, aggregate, strings.Join(stringsFromTypes(parameters), ", "))
	var sfunc = fmt.Sprint(stateFunc)
	var stype = fmt.Sprint(stateType)

	q = append(q, sig, "sfunc", sfunc, "stype", stype)

	if len(finalFunc) > 0 {
		q = append(q, "finalfunc", fmt.Sprint(finalFunc))
	}

	if len(init) > 0 {
		q = append(q, "initcond", string(init))
	}

	return strings.Join(q, " ")
}

func QueryAggregateDrop(keyspace, aggregate Identifier, parameters []Type, options ...Option) string {
	return queryDrop("aggregate", queryFunc(keyspace, aggregate, parameters), options)
}

func QueryColumnAlter() string {
	return ""
}

func QueryFunctionCreate(keyspace, function Identifier, parameters []Type, returns Type, language string, body Term, options ...Option) string {
	var m = optionMap(options)
	var q []string

	if optionHas(m, OptionReplace) {
		q = append(q, "replace")
	} else {
		q = append(q, "create")
	}

	q = append(q, "function")

	if optionHas(m, OptionIfNotExists) {
		q = append(q, "if not exists")
	}

	q = append(q, fmt.Sprintf("%v.%v(%v)", keyspace, function, strings.Join(stringsFromTypes(parameters), ", ")))

	if optionHas(m, OptionCalled) {
		q = append(q, "called")
	} else {
		q = append(q, "returns null")
	}

	q = append(q, "on null input returns", string(returns), "language", language, "as", string(body))

	return strings.Join(q, " ")
}

func QueryFunctionDrop(keyspace, function Identifier, parameters []Type, options ...Option) string {
	return queryDrop("function", queryFunc(keyspace, function, parameters), options)
}

func QueryIndexCreate(keyspace, table, column, index Identifier, keys bool, options ...Option) string {
	var m = optionMap(options)
	var q = []string{"create index"}

	if optionHas(m, OptionIfNotExists) {
		q = append(q, "if not exists")
	}

	var id string

	if keys {
		id = fmt.Sprintf("keys(%v)", column)
	} else {
		id = fmt.Sprint(column)
	}

	q = append(q, fmt.Sprintf("%v on %v.%v(%v)", index, keyspace, table, id))

	if v, ok := optionGet(m, OptionUsing); ok {
		q = append(q, "using", string(ConstantString(v[0].(string))))

		if v, ok := optionGet(m, OptionOptions); ok {
			q = append(q, "with options", string(TermMap(v[0].(map[Term]Term))))
		}
	}

	return strings.Join(q, " ")
}

func QueryIndexDrop(keyspace, index Identifier, options ...Option) string {
	return queryDrop("index", queryID(keyspace, index), options)
}

func QueryKeyspaceAlter(keyspace Identifier, options ...Option) string {
	var m = optionMap(options)
	var q = []string{fmt.Sprintf("alter keyspace %v", keyspace)}

	q = append(q, fmt.Sprintf("%v with", keyspace))

	var ss []string

	for k, v := range m[OptionProperties][0].(map[Identifier]Term) {
		ss = append(ss, fmt.Sprintf("%v = %v", k, v[0]))
	}

	q = append(q, strings.Join(ss, " and "))

	return strings.Join(q, " ")
}

func QueryKeyspaceCreate(keyspace Identifier, options ...Option) string {
	var m = optionMap(options)
	var q = []string{"create keyspace"}

	if optionHas(m, OptionIfNotExists) {
		q = append(q, "if not exists")
		delete(m, OptionIfNotExists)
	}

	q = append(q, fmt.Sprintf("%v with", keyspace))

	var ss []string

	for k, v := range m[OptionProperties][0].(map[Identifier]Term) {
		ss = append(ss, fmt.Sprintf("%v = %v", k, v[0]))
	}

	q = append(q, strings.Join(ss, " and "))

	return strings.Join(q, " ")
}

func QueryKeyspaceDrop(keyspace Identifier, options ...Option) string {
	return queryDrop("keyspace", fmt.Sprint(keyspace), options)
}

func QueryMaterializedViewAlter() string {
	return ""
}

func QueryMaterializedViewCreate(keyspace, table, view Identifier, options ...Option) string {
	var m = optionMap(options)
	var q = []string{"create materialized view"}

	if optionHas(m, OptionIfNotExists) {
		q = append(q, "if not exists")
	}

	q = append(q, fmt.Sprintf("%v as select", view))

	// TODO

	return strings.Join(q, " ")
}

func QueryMaterializedViewDrop(keyspace, table Identifier, options ...Option) string {
	return queryDrop("materialized view", queryID(keyspace, table), options)
}

func QueryTableAlter(keyspace, table Identifier, options ...Option) string {
	var m = optionMap(options)
	var q = []string{fmt.Sprintf("alter table %v.%v with", keyspace, table)}
	var ss []string

	if optionHas(m, OptionClusteringOrder) {
		ss = append(ss, "clustering order")
	}

	if optionHas(m, OptionCompactStorage) {
		ss = append(ss, "compact storage")
	}

	if vs, ok := m[OptionProperties]; ok {
		for k, v := range vs[0].(map[Identifier]Term) {
			ss = append(ss, fmt.Sprintf("%v = %v", k, v))
		}
	}

	q = append(q, strings.Join(ss, " and "))

	return strings.Join(q, " ")
}

func QueryTableColumnAlter() string {
	return ""
}

func QueryTableCreate(keyspace, table Identifier, columns map[Identifier]Type, static map[Identifier]struct{}, partition, cluster []Identifier, options ...Option) string {
	var m = optionMap(options)
	var q = []string{"create table"}

	if optionHas(m, OptionIfExists) {
		q = append(q, "if not exists")
	}

	var ss []string

	for i, t := range columns {
		var s string

		if _, ok := static[i]; ok {
			s = "%v %v static"
		} else {
			s = "%v %v"
		}

		ss = append(ss, fmt.Sprintf(s, i, t))
	}

	var key = strings.Join(stringsFromIdentifiers(partition), ", ")

	if len(cluster) > 0 {
		key = fmt.Sprintf("(%v), %v", key, strings.Join(stringsFromIdentifiers(cluster), ", "))
	}

	ss = append(ss, fmt.Sprintf("primary key (%v)"), key)
	q = append(q, fmt.Sprintf("%v.%v(%v)", keyspace, table, strings.Join(ss, ", ")))
	ss = nil

	if optionHas(m, OptionClusteringOrder) {
		ss = append(ss, "clustering order")
	}

	if optionHas(m, OptionCompactStorage) {
		ss = append(ss, "compact storage")
	}

	if vs, ok := m[OptionProperties]; ok {
		for k, v := range vs[0].(map[Identifier]Term) {
			ss = append(ss, fmt.Sprintf("%v = %v", k, v))
		}
	}

	if len(ss) > 0 {
		q = append(q, "with", strings.Join(ss, " and "))
	}

	return strings.Join(q, " ")
}

func QueryTableDrop(keyspace, table Identifier, options ...Option) string {
	return queryDrop("table", queryID(keyspace, table), options)
}

func QueryTableTruncate(keyspace, table Identifier) string {
	return fmt.Sprintf("truncate table %v.%v", keyspace, table)
}

func QueryTriggerCreate(keyspace, table, trigger Identifier, class string, options ...Option) string {
	var m = optionMap(options)
	var q = []string{"create trigger"}

	if optionHas(m, OptionIfExists) {
		q = append(q, "if not exists")
	}

	if len(trigger) > 0 {
		q = append(q, fmt.Sprint(trigger))
	}

	q = append(q, fmt.Sprintf("on %v.%v using %v", keyspace, table, string(ConstantString(class))))

	return strings.Join(q, " ")
}

func QueryTriggerDrop(keyspace, table, trigger Identifier, options ...Option) string {
	var m = optionMap(options)
	var q = []string{"drop trigger"}

	if optionHas(m, OptionIfExists) {
		q = append(q, "if not exists")
	}

	if len(trigger) > 0 {
		q = append(q, fmt.Sprint(trigger))
	}

	q = append(q, fmt.Sprintf("on %v.%v", keyspace, table))

	return strings.Join(q, " ")
}

func QueryTypeCreate(keyspace, type_ Identifier, fields map[Identifier]Type, options ...Option) string {
	var m = optionMap(options)
	var q = []string{"create type"}

	if optionHas(m, OptionIfExists) {
		q = append(q, "if not exists")
	}

	var ss []string

	for i, t := range fields {
		ss = append(ss, fmt.Sprintf("%v %v", i, t))
	}

	q = append(q, fmt.Sprintf("%v.%v(%v)", keyspace, type_, strings.Join(ss, ", ")))

	return strings.Join(q, " ")
}

func QueryTypeDrop(keyspace, type_ Identifier, options ...Option) string {
	return queryDrop("type", queryID(keyspace, type_), options)
}

func QueryTypeFieldAlter(keyspace, type_, field Identifier, data Type) string {
	return fmt.Sprintf("alter type %v.%v alter %v type %v", keyspace, type_, field, data)
}

func QueryTypeFieldAdd(keyspace, type_, field Identifier, data Type) string {
	return fmt.Sprintf("alter type %v.%v add %v %v", keyspace, type_, field, data)
}

func QueryTypeFieldRename(keyspace, type_, renames map[Identifier]Identifier) string {
	var ss []string

	for k, v := range renames {
		ss = append(ss, fmt.Sprintf("%v to %v", k, v))
	}

	return fmt.Sprintf("alter type %v.%v rename %v", keyspace, type_, strings.Join(ss, " and "))
}

func optionGet(m map[string][]interface{}, name string) ([]interface{}, bool) {
	var v, ok = m[name]

	return v, ok
}

func optionHas(m map[string][]interface{}, name string) bool {
	var _, ok = m[name]

	return ok
}

func optionMap(options []Option) map[string][]interface{} {
	var m = map[string][]interface{}{}

	for _, o := range options {
		o(m)
	}

	return m
}

func queryDrop(kind string, id string, options []Option) string {
	var m = optionMap(options)
	var q = []string{"drop", kind}

	if optionHas(m, OptionIfExists) {
		q = append(q, string(OptionIfExists))
	}

	q = append(q, id)

	return strings.Join(q, " ")
}

func queryFunc(keyspace, function Identifier, parameters []Type) string {
	return fmt.Sprintf("%v.%v(%v)", keyspace, function, strings.Join(stringsFromTypes(parameters), ", "))
}

func queryID(first, second Identifier) string {
	return fmt.Sprintf("%v.%v", first, second)
}

func stringsFromIdentifiers(is []Identifier) []string {
	var ss []string

	for _, i := range is {
		ss = append(ss, string(i))
	}

	return ss
}

func stringsFromTypes(ts []Type) []string {
	var ss []string

	for _, t := range ts {
		ss = append(ss, string(t))
	}

	return ss
}

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

func ConstantHex(s string) Constant {
	return Constant(fmt.Sprintf("0x%v", s))
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

type Identifier string

const (
	FuncAvg             Identifier = "avg"
	FuncDateOf          Identifier = "dateof"
	FuncFromJSON        Identifier = "fromjson"
	FuncMax             Identifier = "max"
	FuncMaxTimeuuid     Identifier = "maxtimeuuid"
	FuncMin             Identifier = "min"
	FuncMinTimeuuid     Identifier = "mintimeuuid"
	FuncNow             Identifier = "now"
	FuncSum             Identifier = "sum"
	FuncToDate          Identifier = "todate"
	FuncToJSON          Identifier = "tojson"
	FuncToken           Identifier = "token"
	FuncToTimestamp     Identifier = "totimestamp"
	FuncToUnixTimestamp Identifier = "tounixtimestamp"
	FuncUnixTimestampOf Identifier = "unixtimestampof"
	FuncUUID            Identifier = "uuid"
)

const (
	FuncBlobToAscii     Identifier = "blobtoascii"
	FuncBlobToBigint    Identifier = "blobtobigint"
	FuncBlobToBoolean   Identifier = "blobtoboolean"
	FuncBlobToCounter   Identifier = "blobtocounter"
	FuncBlobToDate      Identifier = "blobtodate"
	FuncBlobToDecimal   Identifier = "blobtodecimal"
	FuncBlobToDouble    Identifier = "blobtodouble"
	FuncBlobToFloat     Identifier = "blobtofloat"
	FuncBlobToInet      Identifier = "blobtoinet"
	FuncBlobToInt       Identifier = "blobtoint"
	FuncBlobToSmallint  Identifier = "blobtosmallint"
	FuncBlobToText      Identifier = "blobtotext"
	FuncBlobToTime      Identifier = "blobtotime"
	FuncBlobToTimestamp Identifier = "blobtotimestamp"
	FuncBlobToTimeuuid  Identifier = "blobtotimeuuid"
	FuncBlobToTinyint   Identifier = "blobtotinyint"
	FuncBlobToUuid      Identifier = "blobtouuid"
	FuncBlobToVarchar   Identifier = "blobtovarchar"
	FuncBlobToVarint    Identifier = "blobtovarint"
)

const (
	FuncAsciiToBlob     Identifier = "asciitoblob"
	FuncBigintToBlob    Identifier = "biginttoblob"
	FuncBooleanToBlob   Identifier = "booleantoblob"
	FuncCounterToBlob   Identifier = "countertoblob"
	FuncDateToBlob      Identifier = "datetoblob"
	FuncDecimalToBlob   Identifier = "decimaltoblob"
	FuncDoubleToBlob    Identifier = "doubletoblob"
	FuncFloatToBlob     Identifier = "floattoblob"
	FuncInetToBlob      Identifier = "inettoblob"
	FuncIntToBlob       Identifier = "inttoblob"
	FuncSmallintToBlob  Identifier = "smallinttoblob"
	FuncTextToBlob      Identifier = "texttoblob"
	FuncTimeToBlob      Identifier = "timetoblob"
	FuncTimestampToBlob Identifier = "timestamptoblob"
	FuncTimeuuidToBlob  Identifier = "timeuuidtoblob"
	FuncTinyintToBlob   Identifier = "tinyinttoblob"
	FuncUuidToBlob      Identifier = "uuidtoblob"
	FuncVarcharToBlob   Identifier = "varchartoblob"
	FuncVarintToBlob    Identifier = "varinttoblob"
)

func (i Identifier) String() string {
	return fmt.Sprintf("%q", i)
}

type Operator string

const (
	OperatorContains     Operator = "contains"
	OperatorContainsKey  Operator = "contains key"
	OperatorEqual        Operator = "="
	OperatorGreater      Operator = ">"
	OperatorGreaterEqual Operator = ">="
	OperatorIn           Operator = "in"
	OperatorLess         Operator = "<"
	OperatorLessEqual    Operator = "<="
)

type Option func(map[string][]interface{})

func NewOption(name string, values ...interface{}) Option {
	return func(m map[string][]interface{}) {
		m[name] = append(m[name], values...)
	}
}

type Order string

const (
	OrderAscending  Order = "asc"
	OrderDescending Order = "desc"
)

type Relation string

func NewRelation(left Term, o Operator, right Term) Relation {
	return Relation(fmt.Sprintf("%v %v %v", left, o, right))
}

type Selector string

func SelectorAlias(s Selector, alias Identifier) Selector {
	return Selector(fmt.Sprintf("%v as %v", s, alias))
}

func SelectorFunc(function Identifier, arguments ...Selector) Selector {
	var ss []string

	for _, a := range arguments {
		ss = append(ss, string(a))
	}

	return Selector(fmt.Sprintf("%v(%v)", function, strings.Join(ss, ", ")))
}

func SelectorIdentifier(i Identifier) Selector {
	return Selector(fmt.Sprint(i))
}

func SelectorTTL(i Identifier) Selector {
	return Selector(fmt.Sprintf("ttl(%v)", i))
}

func SelectorWriteTime(i Identifier) Selector {
	return Selector(fmt.Sprintf("writetime(%v)", i))
}

type Term string

func TermConstant(c Constant) Term {
	return Term(c)
}

func TermFunc(function Identifier, arguments ...Term) Term {
	var ss []string

	for _, a := range arguments {
		ss = append(ss, string(a))
	}

	return Term(fmt.Sprintf("%v(%v)", function, strings.Join(ss, ", ")))
}

func TermIdentifier(i Identifier) Term {
	return Term(fmt.Sprint(i))
}

func TermIndex(i Identifier, t Term) Term {
	return Term(fmt.Sprintf("%v[%v]", i, t))
}

func TermList(t ...Term) Term {
	var ss []string

	for _, t := range t {
		ss = append(ss, string(t))
	}

	return Term(fmt.Sprintf("[%v]", strings.Join(ss, ", ")))
}

func TermMap(m map[Term]Term) Term {
	var ss []string

	for k, v := range m {
		ss = append(ss, fmt.Sprintf("%v: %v", k, v))
	}

	return Term(fmt.Sprintf("{%v}", strings.Join(ss, ", ")))
}

func TermSet(t ...Term) Term {
	var ss []string

	for _, t := range t {
		ss = append(ss, string(t))
	}

	return Term(fmt.Sprintf("{%v}", strings.Join(ss, ", ")))
}

func TermTuple(t ...Term) Term {
	var ss []string

	for _, t := range t {
		ss = append(ss, string(t))
	}

	return Term(fmt.Sprintf("(%v)", strings.Join(ss, ", ")))
}

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

type Variable string

const VariableAnonymous Variable = "?"

func NewVariable(name string) Variable {
	return Variable(fmt.Sprintf(":%v", name))
}
