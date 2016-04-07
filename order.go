package cuckle

// Order is an ordering direction, either ascending or descending.
type Order string

// The two orderings.
const (
	// OrderAscending is ascending order.
	OrderAscending Order = "asc"

	// OrderDescending is descending order.
	OrderDescending Order = "desc"
)
