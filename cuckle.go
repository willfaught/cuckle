package cuckle

import (
	"fmt"
	"strings"
)

const (
	KeyspaceDurableWrites = "durable writes"
	KeyspaceReplication   = "replication"
)

func QueryAggregateCreate(keyspace, aggregate, stateFunc Identifier, parameters []Type, stateType Type, o ...Option) string {
	var options = combine(o)
	var q []string

	if _, ok := options[optionReplace]; ok {
		q = append(q, "replace")
	} else {
		q = append(q, "create")
	}

	q = append(q, "aggregate")

	if _, ok := options[optionIfNotExists]; ok {
		q = append(q, "if not exists")
	}

	var sig = fmt.Sprintf("%v.%v(%v)", keyspace, aggregate, strings.Join(stringsFromTypes(parameters), ", "))
	var sfunc = fmt.Sprint(stateFunc)
	var stype = fmt.Sprint(stateType)

	q = append(q, sig, "sfunc", sfunc, "stype", stype)

	if f, ok := options[optionFinalFunc]; ok {
		q = append(q, "finalfunc", fmt.Sprint(f))
	}

	if i, ok := options[optionInitCond]; ok {
		q = append(q, "initcond", i.(string))
	}

	return strings.Join(q, " ")
}

func QueryAggregateDrop(keyspace, aggregate Identifier, parameters []Type, o ...Option) string {
	return queryDrop("aggregate", queryFunc(keyspace, aggregate, parameters), o)
}

func QueryFunctionCreate(keyspace, function Identifier, parameters []Type, returns Type, language string, body string, o ...Option) string {
	var options = combine(o)
	var q []string

	if _, ok := options[optionReplace]; ok {
		q = append(q, "replace")
	} else {
		q = append(q, "create")
	}

	q = append(q, "function")

	if _, ok := options[optionIfNotExists]; ok {
		q = append(q, "if not exists")
	}

	q = append(q, fmt.Sprintf("%v.%v(%v)", keyspace, function, strings.Join(stringsFromTypes(parameters), ", ")))

	if _, ok := options[optionCalled]; ok {
		q = append(q, "called")
	} else {
		q = append(q, "returns null")
	}

	q = append(q, "on null input returns", string(returns), "language", language, "as", string(ConstantString(body)))

	return strings.Join(q, " ")
}

func QueryFunctionDrop(keyspace, function Identifier, parameters []Type, o ...Option) string {
	return queryDrop("function", queryFunc(keyspace, function, parameters), o)
}

func QueryIndexCreate(keyspace, table, column Identifier, o ...Option) string {
	var options = combine(o)
	var q = []string{"create index"}

	if _, ok := options[optionIfNotExists]; ok {
		q = append(q, "if not exists")
	}

	if n, ok := options[optionIndexIdentifier]; ok {
		q = append(q, fmt.Sprint(n))
	}

	var id string

	if _, ok := options[optionIndexKeys]; ok {
		id = fmt.Sprintf("keys(%v)", column)
	} else {
		id = fmt.Sprint(column)
	}

	q = append(q, fmt.Sprintf("on %v.%v(%v)", keyspace, table, id))

	if u, ok := options[optionUsing]; ok {
		q = append(q, "using", string(ConstantString(u.(string))))

		if o, ok := options[optionOptions]; ok {
			q = append(q, "with options", string(TermMap(o.(map[Term]Term))))
		}
	}

	return strings.Join(q, " ")
}

func QueryIndexDrop(keyspace, index Identifier, o ...Option) string {
	return queryDrop("index", queryID(keyspace, index), o)
}

func QueryKeyspaceAlter(keyspace Identifier, properties map[Identifier]Term) string {
	var q = []string{fmt.Sprintf("alter keyspace %v with", keyspace)}
	var ss []string

	for i, t := range properties {
		ss = append(ss, fmt.Sprintf("%v = %v", i, t))
	}

	q = append(q, strings.Join(ss, " and "))

	return strings.Join(q, " ")
}

func QueryKeyspaceCreate(keyspace Identifier, properties map[Identifier]Term, o ...Option) string {
	var options = combine(o)
	var q = []string{"create keyspace"}

	if _, ok := options[optionIfNotExists]; ok {
		q = append(q, "if not exists")
	}

	q = append(q, fmt.Sprintf("%v with", keyspace))

	var ss []string

	for i, t := range properties {
		ss = append(ss, fmt.Sprintf("%v = %v", i, t))
	}

	q = append(q, strings.Join(ss, " and "))

	return strings.Join(q, " ")
}

func QueryKeyspaceDrop(keyspace Identifier, o ...Option) string {
	return queryDrop("keyspace", fmt.Sprint(keyspace), o)
}

func QueryMaterializedViewAlter() string {
	return ""
}

func QueryMaterializedViewCreate(keyspace, table, view Identifier, o ...Option) string {
	var options = combine(o)
	var q = []string{"create materialized view"}

	if _, ok := options[optionIfNotExists]; ok {
		q = append(q, "if not exists")
	}

	q = append(q, fmt.Sprintf("%v as select", view))

	// TODO

	return strings.Join(q, " ")
}

func QueryMaterializedViewDrop(keyspace, table Identifier, o ...Option) string {
	return queryDrop("materialized view", queryID(keyspace, table), o)
}

func QueryRowsDelete(keyspace, table Identifier, r []Relation, o ...Option) string {
	var options = combine(o)
	var q = []string{"delete"}

	if sels, ok := options[optionSelectors]; ok {
		var strs []string

		for _, sel := range sels.([]Selector) {
			strs = append(strs, string(sel))
		}

		q = append(q, strings.Join(strs, ", "))
	}

	q = append(q, fmt.Sprintf("from %v.%v", keyspace, table))

	if t, ok := options[optionTimestamp]; ok {
		q = append(q, fmt.Sprintf("using timestamp %v", t))
	}

	var ss []string

	for _, r := range r {
		ss = append(ss, string(r))
	}

	q = append(q, fmt.Sprintf("where %v", strings.Join(ss, " and ")))

	if _, ok := options[optionIfExists]; ok {
		q = append(q, "if exists")
	}

	if cs, ok := options[optionConditions]; ok {
		ss = nil

		for _, c := range cs.([]Relation) {
			ss = append(ss, string(c))
		}

		q = append(q, fmt.Sprintf("if %v", strings.Join(ss, " and ")))
	}

	return strings.Join(q, " ")
}

func QueryRowsGet(keyspace, table Identifier, o ...Option) string {
	var options = combine(o)
	var q = []string{"select"}

	if _, ok := options[optionJSON]; ok {
		q = append(q, "json")
	}

	if _, ok := options[optionDistinct]; ok {
		q = append(q, "distinct")
	}

	if sels, ok := options[optionSelectors]; ok {
		var strs []string

		for _, sel := range sels.([]Selector) {
			strs = append(strs, string(sel))
		}

		q = append(q, strings.Join(strs, ", "))
	} else {
		q = append(q, "count(*)")

		if a, ok := options[optionCountAlias]; ok {
			q = append(q, fmt.Sprintf("as %v", a))
		}
	}

	q = append(q, fmt.Sprintf("from %v.%v", keyspace, table))

	if rs, ok := options[optionRelations]; ok {
		var ss []string

		for _, r := range rs.([]Relation) {
			ss = append(ss, string(r))
		}

		q = append(q, fmt.Sprintf("where %v", strings.Join(ss, " and ")))
	}

	if cs, ok := options[optionOrderByColumns]; ok {
		var is = cs.([]Identifier)
		var ds = options[optionOrderByDirections].([]Order)
		var ss []string

		for i := range is {
			ss = append(ss, fmt.Sprintf("%v %v", is[i], ds[i]))
		}

		q = append(q, fmt.Sprintf("order by %v", strings.Join(ss, ", ")))
	}

	if limit, ok := options[optionLimit]; ok {
		q = append(q, fmt.Sprintf("limit %v", limit))
	}

	if _, ok := options[optionAllowFiltering]; ok {
		q = append(q, "allow filtering")
	}

	return strings.Join(q, " ")
}

func QueryTableAlter(keyspace, table Identifier, properties map[Identifier]Term, o ...Option) string {
	var options = combine(o)
	var q = []string{fmt.Sprintf("alter table %v.%v with", keyspace, table)}
	var ss []string

	if _, ok := options[optionClusteringOrder]; ok {
		ss = append(ss, "clustering order")
	}

	if _, ok := options[optionCompactStorage]; ok {
		ss = append(ss, "compact storage")
	}

	for i, t := range properties {
		ss = append(ss, fmt.Sprintf("%v = %v", i, t))
	}

	q = append(q, strings.Join(ss, " and "))

	return strings.Join(q, " ")
}

func QueryTableColumnAdd(keyspace, table, column Identifier, data Type) string {
	return fmt.Sprintf("alter table %v.%v add %v %v", keyspace, table, column, data)
}

func QueryTableColumnAlter(keyspace, table, column Identifier, data Type) string {
	return fmt.Sprintf("alter table %v.%v alter %v type %v", keyspace, table, column, data)
}

func QueryTableColumnDrop(keyspace, table, column Identifier) string {
	return fmt.Sprintf("alter table %v.%v drop %v", keyspace, table, column)
}

func QueryTableCreate(keyspace, table Identifier, columns map[Identifier]Type, static map[Identifier]struct{}, partition, cluster []Identifier, o ...Option) string {
	var options = combine(o)
	var q = []string{"create table"}

	if _, ok := options[optionIfExists]; ok {
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

	if _, ok := options[optionClusteringOrder]; ok {
		ss = append(ss, "clustering order")
	}

	if _, ok := options[optionCompactStorage]; ok {
		ss = append(ss, "compact storage")
	}

	if vs, ok := options[optionProperties]; ok {
		for k, v := range vs.(map[Identifier]Term) {
			ss = append(ss, fmt.Sprintf("%v = %v", k, v))
		}
	}

	if len(ss) > 0 {
		q = append(q, "with", strings.Join(ss, " and "))
	}

	return strings.Join(q, " ")
}

func QueryTableDrop(keyspace, table Identifier, o ...Option) string {
	return queryDrop("table", queryID(keyspace, table), o)
}

func QueryTableTruncate(keyspace, table Identifier) string {
	return fmt.Sprintf("truncate table %v.%v", keyspace, table)
}

func QueryTriggerCreate(keyspace, table, trigger Identifier, class string, o ...Option) string {
	var options = combine(o)
	var q = []string{"create trigger"}

	if _, ok := options[optionIfNotExists]; ok {
		q = append(q, "if not exists")
	}

	if t, ok := options[optionTriggerIdentifier]; ok {
		q = append(q, fmt.Sprint(t))
	}

	q = append(q, fmt.Sprintf("on %v.%v using %v", keyspace, table, string(ConstantString(class))))

	return strings.Join(q, " ")
}

func QueryTriggerDrop(keyspace, table Identifier, o ...Option) string {
	var options = combine(o)
	var q = []string{"drop trigger"}

	if _, ok := options[optionIfExists]; ok {
		q = append(q, "if exists")
	}

	if t, ok := options[optionTriggerIdentifier]; ok {
		q = append(q, fmt.Sprint(t))
	}

	q = append(q, fmt.Sprintf("on %v.%v", keyspace, table))

	return strings.Join(q, " ")
}

func QueryTypeCreate(keyspace, type_ Identifier, fields map[Identifier]Type, o ...Option) string {
	var options = combine(o)
	var q = []string{"create type"}

	if _, ok := options[optionIfNotExists]; ok {
		q = append(q, "if not exists")
	}

	var ss []string

	for i, t := range fields {
		ss = append(ss, fmt.Sprintf("%v %v", i, t))
	}

	q = append(q, fmt.Sprintf("%v.%v(%v)", keyspace, type_, strings.Join(ss, ", ")))

	return strings.Join(q, " ")
}

func QueryTypeDrop(keyspace, type_ Identifier, o ...Option) string {
	return queryDrop("type", queryID(keyspace, type_), o)
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

func queryDrop(kind string, id string, o []Option) string {
	var options = combine(o)
	var q = []string{"drop", kind}

	if _, ok := options[optionIfExists]; ok {
		q = append(q, string(optionIfExists))
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
	OperatorNotEqual     Operator = "!="
)

type Option map[option]interface{}

var (
	OptionAllowFiltering Option = Option{optionAllowFiltering: nil}
	OptionDistinct       Option = Option{optionDistinct: nil}
	OptionIndexKeys      Option = Option{optionIndexKeys: nil}
)

func OptionAliases(aliases map[Identifier]Identifier) Option {
	return Option{optionAliases: aliases}
}

func OptionCountAlias(alias Identifier) Option {
	return Option{optionCountAlias: alias}
}

func OptionFinalFunc(finalFunc Identifier) Option {
	return Option{optionFinalFunc: finalFunc}
}

func OptionIndexIdentifier(index Identifier) Option {
	return Option{optionIndexIdentifier: index}
}

func OptionInitCond(initCond Term) Option {
	return Option{optionInitCond: initCond}
}

func OptionLimit(limit int) Option {
	return Option{optionLimit: limit}
}

func OptionOrder(i []Identifier, o []Order) Option {
	return Option{optionOrderByColumns: i, optionOrderByDirections: o}
}

func OptionRelations(r ...Relation) Option {
	return Option{optionRelations: r}
}

func OptionSelectors(s ...Selector) Option {
	return Option{optionSelectors: s}
}

func OptionTriggerIdentifier(trigger Identifier) Option {
	return Option{optionTriggerIdentifier: trigger}
}

func combine(os []Option) Option {
	var combined = Option{}

	for _, o := range os {
		for k, v := range o {
			combined[k] = v
		}
	}

	return combined
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

func SelectorIndex(i Identifier, t Term) Selector {
	return Selector(fmt.Sprintf("%v[%v]", i, t))
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

type option int

const (
	optionAliases option = iota
	optionAllowFiltering
	optionCalled
	optionClusteringOrder
	optionCompactStorage
	optionConditions
	optionCountAlias
	optionDistinct
	optionFinalFunc
	optionIfExists
	optionIfNotExists
	optionIndexIdentifier
	optionIndexKeys
	optionInitCond
	optionJSON
	optionLimit
	optionOptions
	optionOrderByColumns
	optionOrderByDirections
	optionProperties
	optionRelations
	optionReplace
	optionSelectors
	optionTimestamp
	optionTrigger
	optionTriggerIdentifier
	optionUsing
)
