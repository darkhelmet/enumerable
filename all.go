package enumerable

import "reflect"

// All checks if all values in a collection satisfy a predicate function
func All(collection interface{}, predicate interface{}) bool {
	cv := ensureSlice(collection)
	pv := ensureFuncReturns(predicate, 1, 1, reflect.Bool)
	ensureCanMap(cv, pv)

	length := cv.Len()
	for i := 0; i < length; i++ {
		input := cv.Index(i)
		yes := pv.Call([]reflect.Value{input})[0].Bool()
		if !yes {
			return false
		}
	}

	return true
}
