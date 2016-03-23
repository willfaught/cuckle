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

func QueryBatch(counter bool, timestamp int64, queries []string) string {
	var q = []string{"begin"}

	var kind string

	if counter {
		kind = "counter"
	} else {
		kind = "unlogged"
	}

	q = append(q, kind, "batch")

	if timestamp != 0 {
		q = append(q, fmt.Sprintf("using timestamp %v", timestamp))
	}

	q = append(q, strings.Join(queries, "; "), "apply batch")

	return strings.Join(q, " ")
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

func QueryRowsDelete(keyspace, table Identifier, where []Relation, o ...Option) string {
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

	for _, r := range where {
		ss = append(ss, string(r))
	}

	q = append(q, fmt.Sprintf("where %v", strings.Join(ss, " and ")))

	if _, ok := options[optionIfExists]; ok {
		q = append(q, "if exists")
	}

	if is, ok := options[optionIf]; ok {
		ss = nil

		for _, i := range is.([]Relation) {
			ss = append(ss, string(i))
		}

		q = append(q, fmt.Sprintf("if %v", strings.Join(ss, " and ")))
	}

	return strings.Join(q, " ")
}

func QueryRowsInsert(keyspace, table Identifier, columns []Identifier, values []Term, json string, o ...Option) string {
	var options = combine(o)
	var q = []string{fmt.Sprintf("insert into %v.%v", keyspace, table)}

	if json != "" {
		q = append(q, "json", string(ConstantString(json)))
	}

	if len(columns) > 0 {
		var cs, vs []string

		for i := range columns {
			cs = append(cs, fmt.Sprint(columns[i]))
			vs = append(vs, string(values[i]))
		}

		q = append(q, fmt.Sprintf("(%v) values (%v)", strings.Join(cs, ", "), strings.Join(vs, ", ")))
	}

	if _, ok := options[optionIfNotExists]; ok {
		q = append(q, "if not exists")
	}

	var ss []string

	if t, ok := options[optionTimestamp]; ok {
		ss = append(ss, fmt.Sprintf("timestamp %v", t))
	}

	if t, ok := options[optionTTL]; ok {
		ss = append(ss, fmt.Sprintf("ttl %v", t))
	}

	if len(ss) > 0 {
		q = append(q, fmt.Sprintf("using %v", strings.Join(ss, " and ")))
	}

	return strings.Join(q, " ")
}

func QueryRowsSelect(keyspace, table Identifier, s []Selector, o ...Option) string {
	var options = combine(o)
	var q = []string{"select"}

	if _, ok := options[optionJSON]; ok {
		q = append(q, "json")
	}

	if _, ok := options[optionDistinct]; ok {
		q = append(q, "distinct")
	}

	var ss []string

	for _, s := range s {
		ss = append(ss, string(s))
	}

	q = append(q, strings.Join(ss, ", "), fmt.Sprintf("from %v.%v", keyspace, table))

	if rs, ok := options[optionWhere]; ok {
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

func QueryRowsUpdate(keyspace, table Identifier, assign, where []Relation, o ...Option) string {
	var options = combine(o)
	var q = []string{fmt.Sprintf("update %v.%v", keyspace, table)}
	var ss []string

	if t, ok := options[optionTimestamp]; ok {
		ss = append(ss, fmt.Sprintf("timestamp %v", t))
	}

	if t, ok := options[optionTTL]; ok {
		ss = append(ss, fmt.Sprintf("ttl %v", t))
	}

	if len(ss) > 0 {
		q = append(q, fmt.Sprintf("using %v", strings.Join(ss, " and ")))
	}

	ss = nil

	for _, r := range assign {
		ss = append(ss, string(r))
	}

	q = append(q, fmt.Sprintf("set %v", strings.Join(ss, ", ")))

	ss = nil

	for _, r := range where {
		ss = append(ss, string(r))
	}

	q = append(q, fmt.Sprintf("where %v", strings.Join(ss, " and ")))

	if is, ok := options[optionIf]; ok {
		ss = nil

		for _, i := range is.([]Relation) {
			ss = append(ss, string(i))
		}

		q = append(q, fmt.Sprintf("if %v", strings.Join(ss, " and ")))
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

func QueryViewAlter() string {
	return ""
}

func QueryViewCreate(keyspace, table, view Identifier, o ...Option) string {
	var options = combine(o)
	var q = []string{"create materialized view"}

	if _, ok := options[optionIfNotExists]; ok {
		q = append(q, "if not exists")
	}

	q = append(q, fmt.Sprintf("%v as select", view))

	// TODO

	return strings.Join(q, " ")
}

func QueryViewDrop(keyspace, table Identifier, o ...Option) string {
	return queryDrop("materialized view", queryID(keyspace, table), o)
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
