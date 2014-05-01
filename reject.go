package enumerable

import "reflect"

// Reject filters a collection to exclude those values that satisfy the predicate
func Reject(collection interface{}, predicate interface{}) interface{} {
	cv := ensureSlice(collection)
	pv := ensureFuncReturns(predicate, 1, 1, reflect.Bool)
	ensureCanMap(cv, pv)

	length := cv.Len()
	output := reflect.MakeSlice(reflect.SliceOf(cv.Type().Elem()), 0, cv.Cap())

	for i := 0; i < length; i++ {
		input := cv.Index(i)
		yes := pv.Call([]reflect.Value{input})[0].Bool()
		if !yes {
			output = reflect.Append(output, input)
		}
	}

	return output.Interface()
}
