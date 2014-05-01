package enumerable

import "reflect"

// Detect returns the first value in the collection that matches a predicate and true, or the zero value and false
func Detect(collection interface{}, predicate interface{}) (interface{}, bool) {
	cv := ensureSlice(collection)
	pv := ensureFuncReturns(predicate, 1, 1, reflect.Bool)
	ensureCanMap(cv, pv)

	length := cv.Len()
	for i := 0; i < length; i++ {
		input := cv.Index(i)
		yes := pv.Call([]reflect.Value{input})[0].Bool()
		if yes {
			return input.Interface(), true
		}
	}

	return reflect.Zero(cv.Type().Elem()).Interface(), false
}
