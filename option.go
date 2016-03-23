package cuckle

type Option map[option]interface{}

var (
	OptionAllowFiltering Option = Option{optionAllowFiltering: nil}
	OptionDistinct       Option = Option{optionDistinct: nil}
	OptionIndexKeys      Option = Option{optionIndexKeys: nil}
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

func OptionOrder(i []Identifier, o []Order) Option {
	return Option{optionOrderByColumns: i, optionOrderByDirections: o}
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

type option int

const (
	optionAliases option = iota
	optionAllowFiltering
	optionAssignments
	optionCalled
	optionClusteringOrder
	optionColumns
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
	optionTrigger
	optionTriggerIdentifier
	optionUsing
	optionValues
	optionWhere
)
