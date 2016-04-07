package cuckle

// Option stores optional query parameters.
type Option map[option]interface{}

// Options with no parameters.
var (
	// OptionAllowFiltering allows filtering result rows.
	OptionAllowFiltering Option = Option{optionAllowFiltering: nil}

	// OptionCalled allows functions to be called with null input.
	OptionCalled Option = Option{optionCalled: nil}

	// OptionClusteringOrder is the ordering of rows on disk.
	OptionClusteringOrder Option = Option{optionClusteringOrder: nil}

	// OptionCompactStorage uses a more compact but less flexible layout for tables.
	OptionCompactStorage Option = Option{optionCompactStorage: nil}

	// OptionDistinct removes rows with duplicate partition keys from result rows.
	OptionDistinct Option = Option{optionDistinct: nil}

	// OptionIfExists does not execute create queries for existing primary keys.
	OptionIfExists Option = Option{optionIfExists: nil}

	// OptionIfNotExists does not execute drop queries for invalid primary keys.
	OptionIfNotExists Option = Option{optionIfNotExists: nil}

	// OptionIndexKeys creates an index on map keys.
	OptionIndexKeys Option = Option{optionIndexKeys: nil}

	// OptionJSON uses JSON for insert and select queries.
	OptionJSON Option = Option{optionJSON: nil}

	// OptionReplace replaces any existing aggregates and functions.
	OptionReplace Option = Option{optionReplace: nil}
)

// OptionAliases returns an Option for column aliases.
func OptionAliases(aliases map[Identifier]Identifier) Option {
	return Option{optionAliases: aliases}
}

// OptionConditions returns an Option for lightweight transaction conditions.
func OptionConditions(r ...Relation) Option {
	return Option{optionConditions: r}
}

// OptionFinalFunc returns an Option for an aggregate final function.
func OptionFinalFunc(finalFunc Identifier) Option {
	return Option{optionFinalFunc: finalFunc}
}

// OptionIndexName returns an Option for an index identifier.
func OptionIndexName(index Identifier) Option {
	return Option{optionIndexName: index}
}

// OptionInitCond returns an Option for an aggregate initial condition.
func OptionInitCond(initCond Term) Option {
	return Option{optionInitCond: initCond}
}

// OptionLimit returns an Option for a result row limit.
func OptionLimit(limit int) Option {
	return Option{optionLimit: limit}
}

// OptionOptions returns an Option for arbitrary Term key-value pairs.
func OptionOptions(options map[Term]Term) Option {
	return Option{optionOptions: options}
}

// OptionOrder returns an Option for ordering result rows by columns and directions.
func OptionOrder(i []Identifier, o []Order) Option {
	return Option{optionOrderByColumns: i, optionOrderByDirections: o}
}

// OptionSelectors returns an Option for selecting column values.
func OptionSelectors(s ...Selector) Option {
	return Option{optionSelectors: s}
}

// OptionTTL returns an Option for insert and update time-to-lives.
func OptionTTL(ttl int64) Option {
	return Option{optionTTL: ttl}
}

// OptionTimestamp returns an Option for insert and update timestamps.
func OptionTimestamp(timestamp int64) Option {
	return Option{optionTimestamp: timestamp}
}

// OptionTriggerName returns an Option for a trigger identifier.
func OptionTriggerName(trigger Identifier) Option {
	return Option{optionTriggerName: trigger}
}

// OptionUsing returns an Option for a custom index class.
func OptionUsing(class string) Option {
	return Option{optionUsing: class}
}

// OptionWhere returns an Option for criteria relations.
func OptionWhere(r ...Relation) Option {
	return Option{optionWhere: r}
}

// OptionWith is key-value pairs.
func OptionWith(options map[Identifier]Term) Option {
	return Option{optionWith: options}
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

type option int

const (
	optionAliases option = iota
	optionAllowFiltering
	optionCalled
	optionClusteringOrder
	optionCompactStorage
	optionConditions
	optionDistinct
	optionFinalFunc
	optionIfExists
	optionIfNotExists
	optionIndexName
	optionIndexKeys
	optionInitCond
	optionJSON
	optionLimit
	optionOptions
	optionOrderByColumns
	optionOrderByDirections
	optionReplace
	optionSelectors
	optionTTL
	optionTimestamp
	optionTriggerName
	optionUsing
	optionWhere
	optionWith
)
