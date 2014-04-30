package enumerable

import "reflect"

// Any checks if any values in a collection satisfy a predicate function
func Any(collection interface{}, predicate interface{}) bool {
	cv := ensureSlice(collection)
	pv := ensureFuncReturns(predicate, 1, 1, reflect.Bool)
	ensureCanMap(cv, pv)

	length := cv.Len()
	for i := 0; i < length; i++ {
		input := cv.Index(i)
		yes := pv.Call([]reflect.Value{input})[0].Bool()
		if yes {
			return true
		}
	}

	return false
}
