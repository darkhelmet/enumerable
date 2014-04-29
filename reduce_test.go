package enumerable_test

import (
	"fmt"
	"testing"

	E "github.com/darkhelmet/enumerable"
	"github.com/stretchr/testify/assert"
)

func ExampleReduce_withInitialValue() {
	ints := []int{1, 2, 3, 4, 5}
	sum := E.Reduce(ints, 10, func(memo, i int) int {
		return memo + i
	}).(int)
	fmt.Println(sum)
	// Output: 25
}

func ExampleReduce_implicitInitialValue() {
	ints := []int{1, 2, 3, 4, 5}
	sum := E.Reduce(ints, nil, func(memo, i int) int {
		return memo + i
	}).(int)
	fmt.Println(sum)
	// Output: 15
}

var (
	reduceInts    = []int{1, 2, 3, 4, 5}
	reduceStrings = []string{"wat", "batman", "superman"}
)

func TestReduceWorks(t *testing.T) {
	sum := E.Reduce(reduceInts, 0, func(i, j int) int {
		return i + j
	}).(int)
	assert.Equal(t, sum, 15)

	sum = E.Reduce(reduceStrings, 0, func(i int, s string) int {
		return i + len(s)
	}).(int)
	assert.Equal(t, sum, 17)

	sum = E.Reduce([]int{}, 10, func(i, j int) int {
		return i + j
	}).(int)
	assert.Equal(t, sum, 10)
}

func TestReduceWithNilInitialValue(t *testing.T) {
	concat := E.Reduce(reduceStrings, nil, func(memo, s string) string {
		return fmt.Sprintf("%s, %s", memo, s)
	}).(string)
	assert.Equal(t, concat, "wat, batman, superman")
}

func TestReduceRequiresSlice(t *testing.T) {
	assert.Panics(t, func() {
		E.Reduce(1, nil, "doesn't matter")
	}, "requires a slice as the first arg")
}

func TestReduceRequiresAFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Reduce([]int{1}, nil, "not a func")
	}, "requires a func as the second arg")
}

func TestReduceRequiresTwoArgFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Reduce([]int{1}, nil, func(i int) {})
	}, "requires a 2 arg func")

	assert.Panics(t, func() {
		E.Reduce([]int{1}, nil, func(i, j, k int) {})
	}, "requires a 2 arg func")
}

func TestReduceRequiresASingleReturnFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Reduce([]int{1}, nil, func(i, j int) {})
	}, "requires a single return arg func")

	assert.Panics(t, func() {
		E.Reduce([]int{1}, nil, func(i, j int) (int, int) {
			return 0, 0
		})
	}, "requires a single return arg func")
}

func TestReduceRequiresFuncToTakeInitialArg(t *testing.T) {
	assert.Panics(t, func() {
		E.Reduce([]int{1}, nil, func(i string, j int) int {
			return 0
		})
	}, "requires func to take initial arg type")

	assert.Panics(t, func() {
		E.Reduce([]int{1}, 0.0, func(i float64, j int) int {
			return 0
		})
	}, "requires func to return initial arg type")
}

func TestReduceRequiresFuncToTakeSliceArg(t *testing.T) {
	assert.Panics(t, func() {
		E.Reduce([]int{1}, nil, func(i int, j string) int {
			return 0
		})
	}, "requires func to take initial arg")
}

var (
	benchInts         = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	benchIntsExpected = 55
)

func BenchmarkGenericReduce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := E.Reduce(benchInts, 0, func(i, j int) int {
			return i + j
		}).(int)
		assert.Equal(b, result, benchIntsExpected)
	}
}

func BenchmarkHandReduce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Generic does allocation, so...
		result := 0
		for _, v := range benchInts {
			result += v
		}
		assert.Equal(b, result, benchIntsExpected)
	}
}
