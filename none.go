package enumerable

// None checks that no values in a collection satisfy a predicate function
func None(collection interface{}, predicate interface{}) bool {
	return !Any(collection, predicate)
}
