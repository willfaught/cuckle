package cuckle

type Option map[option]interface{}

var (
	OptionAllowFiltering  Option = Option{optionAllowFiltering: nil}
	OptionCalled          Option = Option{optionCalled: nil}
	OptionClusteringOrder Option = Option{optionClusteringOrder: nil}
	OptionCompactStorage  Option = Option{optionCompactStorage: nil}
	OptionDistinct        Option = Option{optionDistinct: nil}
	OptionIfExists        Option = Option{optionIfExists: nil}
	OptionIfNotExists     Option = Option{optionIfNotExists: nil}
	OptionIndexKeys       Option = Option{optionIndexKeys: nil}
	OptionJSON            Option = Option{optionJSON: nil}
	OptionReplace         Option = Option{optionReplace: nil}
)

func OptionAliases(aliases map[Identifier]Identifier) Option {
	return Option{optionAliases: aliases}
}

func OptionAssignments(r ...Relation) Option {
	return Option{optionAssignments: r}
}

func OptionFinalFunc(finalFunc Identifier) Option {
	return Option{optionFinalFunc: finalFunc}
}

func OptionIf(r ...Relation) Option {
	return Option{optionIf: r}
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

func OptionOptions(options map[Term]Term) Option {
	return Option{optionOptions: options}
}

func OptionOrder(i []Identifier, o []Order) Option {
	return Option{optionOrderByColumns: i, optionOrderByDirections: o}
}

func OptionSelectors(s ...Selector) Option {
	return Option{optionSelectors: s}
}

func OptionTTL(ttl int64) Option {
	return Option{optionTTL: ttl}
}

func OptionTimestamp(timestamp int64) Option {
	return Option{optionTimestamp: timestamp}
}

func OptionTriggerIdentifier(trigger Identifier) Option {
	return Option{optionTriggerIdentifier: trigger}
}

func OptionUsing(class string) Option {
	return Option{optionUsing: class}
}

func OptionWhere(r ...Relation) Option {
	return Option{optionWhere: r}
}

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
	optionAssignments
	optionCalled
	optionClusteringOrder
	optionCompactStorage
	optionDistinct
	optionFinalFunc
	optionIf
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
	optionReplace
	optionSelectors
	optionTTL
	optionTimestamp
	optionTriggerIdentifier
	optionUsing
	optionWhere
	optionWith
)
