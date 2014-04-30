package enumerable

import "reflect"

// Count counts the number of values in a collection that satisfy a predicate
func Count(collection interface{}, predicate interface{}) int {
	cv := ensureSlice(collection)
	pv := ensureFuncReturns(predicate, 1, 1, reflect.Bool)
	ensureCanMap(cv, pv)

	length := cv.Len()
	count := 0
	for i := 0; i < length; i++ {
		input := cv.Index(i)
		if pv.Call([]reflect.Value{input})[0].Bool() {
			count++
		}
	}

	return count
}
