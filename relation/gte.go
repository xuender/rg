package relation

import "cmp"

// Gte returns a function that checks if the input value is greater than or equal to the given base value.
// The type parameter T must implement the cmp.Ordered interface.
func Gte[T cmp.Ordered](base T) func(T) bool {
	return func(val T) bool {
		return base >= val
	}
}