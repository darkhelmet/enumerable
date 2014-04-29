package enumerable_test

import (
	"fmt"
	"strings"
	"testing"

	E "github.com/darkhelmet/enumerable"
	"github.com/stretchr/testify/assert"
)

func ExampleMap() {
	words := strings.Fields("To be or not to be, that is the question!")
	lengths := E.Map(words, func(s string) int {
		return len(s)
	}).([]int)
	fmt.Println(lengths)
	// Output: [2 2 2 3 2 3 4 2 3 9]
}

func TestMapWorks(t *testing.T) {
	squares := E.Map([]int{1, 2, 3}, func(i int) int {
		return i * i
	}).([]int)
	assert.Equal(t, squares, []int{1, 4, 9})

	upper := E.Map([]string{"wat"}, strings.ToUpper).([]string)
	assert.Equal(t, upper, []string{"WAT"})
}

func TestMapRequiresSlice(t *testing.T) {
	assert.Panics(t, func() {
		E.Map(1, "doesn't matter")
	}, "requires a slice as the first arg")
}

func TestMapRequiresAFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Map([]int{1}, "not a func")
	}, "requires a func as the second arg")
}

func TestMapRequiresASingleArgFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Map([]int{1}, func() {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.Map([]int{1}, func(i, j int) {})
	}, "requires a single arg function")
}

func TestMapRequiresASingleReturnFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Map([]int{1}, func(i int) {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.Map([]int{1}, func(i int) (int, int) { return 1, 0 })
	}, "requires a single arg function")
}

func TestMapRequiresFuncToTakeSliceArg(t *testing.T) {
	assert.Panics(t, func() {
		E.Map([]int{1}, func(s string) int { return 0 })
	}, "requires the func to take the same type as the slice")
}

var (
	benchMapInts         = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	benchMapIntsExpected = []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
)

func BenchmarkGenericMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := E.Map(benchMapInts, func(i int) int {
			return i * i
		}).([]int)
		assert.Equal(b, result, benchMapIntsExpected)
	}
}

func BenchmarkHandMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Generic does allocation, so...
		result := make([]int, 0, len(benchMapInts))
		for _, v := range benchMapInts {
			result = append(result, v*v)
		}
		assert.Equal(b, result, benchMapIntsExpected)
	}
}
