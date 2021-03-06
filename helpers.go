package enumerable

import (
	"fmt"
	"reflect"
)

func panicf(format string, a ...interface{}) {
	panic(fmt.Errorf(format, a...))
}

func ensureKind(i interface{}, k reflect.Kind) reflect.Value {
	v := reflect.ValueOf(i)
	if v.Kind() != k {
		panicf("expected %s, got %T", k, i)
	}
	return v
}

func ensureSlice(i interface{}) reflect.Value {
	return ensureKind(i, reflect.Slice)
}

func ensureCanMap(cv reflect.Value, mv reflect.Value) {
	mt := mv.Type()
	it := mt.In(0)
	ct := cv.Type().Elem()
	if it != ct {
		panicf("expected mapper func to take the same type as the collection, %s != %s", it, ct)
	}
}

func ensureCanReduce(cv reflect.Value, mv reflect.Value, iv reflect.Value) {
	ct := cv.Type().Elem()
	mt := mv.Type()
	it := iv.Type()

	ft := mt.In(0)
	if it != ft {
		panicf("expected reduce func first arg to have same type as initial value, %s != %s", it, ft)
	}

	ft = mt.In(1)
	if ct != ft {
		panicf("expected reduce func second arg to have same type as the collection, %s != %s", ct, ft)
	}

	ft = mt.Out(0)
	if it != ft {
		panicf("expected reduce func return to have same type as initial value, %s != %s", it, ft)
	}
}

func ensureFunc(i interface{}, in, out int) reflect.Value {
	v := ensureKind(i, reflect.Func)
	t := v.Type()

	if t.NumIn() != in {
		panicf("expected func to take a single argument, takes %d", t.NumIn())
	}

	if t.NumOut() != out {
		panicf("expected func to return single argument, returns %d", t.NumOut())
	}

	return v
}

func ensureFuncReturns(i interface{}, in, out int, ret reflect.Kind) reflect.Value {
	v := ensureFunc(i, in, out)
	rk := v.Type().Out(0).Kind()
	if rk != ret {
		panicf("expected func to return a bool, returns %s", rk)
	}
	return v
}
