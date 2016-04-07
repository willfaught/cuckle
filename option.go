package cuckle

// Option stores optional query parameters.
type Option map[option]interface{}

// Options with no parameters.
var (
	// OptionAllowFiltering allows filtering result rows.
	OptionAllowFiltering Option = Option{optionAllowFiltering: nil}

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

// OptionClass returns an Option for a custom index class.
func OptionClass(class string) Option {
	return Option{optionClass: class}
}

// OptionClassOptions returns an Option for custom index class options.
func OptionClassOptions(options map[Term]Term) Option {
	return Option{optionClassOptions: options}
}

// OptionFinalFunc returns an Option for an aggregate final function.
func OptionFinalFunc(finalFunc Identifier) Option {
	return Option{optionFinalFunc: finalFunc}
}

// OptionIf returns an Option for lightweight transaction conditions.
func OptionIf(r ...Relation) Option {
	return Option{optionIf: r}
}

// OptionInitCond returns an Option for an aggregate initial condition.
func OptionInitCond(initCond Term) Option {
	return Option{optionInitCond: initCond}
}

// OptionLimit returns an Option for a result row limit.
func OptionLimit(limit int) Option {
	return Option{optionLimit: limit}
}

// OptionName returns an Option for a name.
func OptionName(name Identifier) Option {
	return Option{optionName: name}
}

// OptionOrder returns an Option for ordering result rows by columns and directions.
func OptionOrder(i []Identifier, o []Order) Option {
	return Option{optionOrderByColumns: i, optionOrderByDirections: o}
}

// OptionProperties is key-value pairs.
func OptionProperties(options map[Identifier]Term) Option {
	return Option{optionProperties: options}
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

// OptionWhere returns an Option for conditions.
func OptionWhere(r ...Relation) Option {
	return Option{optionWhere: r}
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
	optionClass
	optionClusteringOrder
	optionCompactStorage
	optionDistinct
	optionFinalFunc
	optionIf
	optionIfExists
	optionIfNotExists
	optionIndexKeys
	optionInitCond
	optionJSON
	optionLimit
	optionName
	optionClassOptions
	optionOrderByColumns
	optionOrderByDirections
	optionProperties
	optionReplace
	optionSelectors
	optionTTL
	optionTimestamp
	optionWhere
)
