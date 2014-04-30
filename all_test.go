package enumerable_test

import (
	"fmt"
	"strings"
	"testing"

	E "github.com/darkhelmet/enumerable"
	"github.com/stretchr/testify/assert"
)

func ExampleAll() {
	words := strings.Fields("To be or not to be, that is the question!")
	okay := E.All(words, func(s string) bool {
		return len(s) < 10
	})
	fmt.Println(okay)
	// Output: true
}

func TestAllWorks(t *testing.T) {
	evens := []int{2, 4, 6, 8}
	even := E.All(evens, func(i int) bool {
		return i%2 == 0
	})
	assert.True(t, even)

	odd := E.All(evens, func(i int) bool {
		return i%2 == 1
	})
	assert.False(t, odd)
}

func TestAllRequiresSlice(t *testing.T) {
	assert.Panics(t, func() {
		E.All(1, "doesn't matter")
	}, "requires a slice as the first arg")
}

func TestAllRequiresAFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.All([]int{1}, "not a func")
	}, "requires a func as the second arg")
}

func TestAllRequiresASingleArgFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.All([]int{1}, func() {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.All([]int{1}, func(i, j int) {})
	}, "requires a single arg function")
}

func TestAllRequiresASingleReturnFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.All([]int{1}, func(i int) {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.All([]int{1}, func(i int) (int, int) { return 1, 0 })
	}, "requires a single arg function")
}

func TestAllRequiresFuncToTakeSliceArg(t *testing.T) {
	assert.Panics(t, func() {
		E.All([]int{1}, func(s string) int { return 0 })
	}, "requires the func to take the same type as the slice")
}

func TestAllRequiresFuncToReturnABool(t *testing.T) {
	assert.Panics(t, func() {
		E.All([]int{1}, func(i int) int { return i })
	}, "requires the func to take the same type as the slice")
}
