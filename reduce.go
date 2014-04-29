package enumerable

import "reflect"

// Reduce is a generic implementation of `reduce` of `map/reduce` fame
func Reduce(collection interface{}, initial interface{}, reducer interface{}) interface{} {
	cv := ensureSlice(collection)
	mv := ensureFunc(reducer, 2, 1)
	length := cv.Len()

	if length == 0 {
		return initial
	}

	start := 0
	var output reflect.Value
	if initial == nil {
		output = cv.Index(0)
		start = 1
	} else {
		output = reflect.ValueOf(initial)
	}

	ensureCanReduce(cv, mv, output)

	for i := start; i < length; i++ {
		right := cv.Index(i)
		output = mv.Call([]reflect.Value{output, right})[0]
	}

	return output.Interface()
}
