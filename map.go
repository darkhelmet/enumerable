package enumerable

import "reflect"

// Map is a generic implementation of `map` of `map/reduce` fame
func Map(collection interface{}, mapper interface{}) interface{} {
	cv := ensureSlice(collection)
	mv := ensureFunc(mapper, 1, 1)
	ensureCanMap(cv, mv)

	mt := mv.Type()
	length := cv.Len()
	output := reflect.MakeSlice(reflect.SliceOf(mt.Out(0)), length, cv.Cap())

	for i := 0; i < length; i++ {
		input := cv.Index(i)
		mapped := mv.Call([]reflect.Value{input})
		output.Index(i).Set(mapped[0])
	}

	return output.Interface()
}
