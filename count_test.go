package enumerable_test

import (
	"fmt"
	"strings"
	"testing"

	E "github.com/darkhelmet/enumerable"
	"github.com/stretchr/testify/assert"
)

func ExampleCount() {
	words := strings.Fields("To be or not to be, that is the question!")
	long := E.Count(words, func(s string) bool {
		return len(s) > 2
	})
	fmt.Println(long)
	// Output: 5
}

func TestCountWorks(t *testing.T) {
	ints := []int{2, 4, 6, 8}
	evens := E.Count(ints, func(i int) bool {
		return i%2 == 0
	})
	assert.Equal(t, 4, evens)

	aboveFive := E.Count(ints, func(i int) bool {
		return i > 5
	})
	assert.Equal(t, 2, aboveFive)
}

func TestCountRequiresSlice(t *testing.T) {
	assert.Panics(t, func() {
		E.Count(1, "doesn't matter")
	}, "requires a slice as the first arg")
}

func TestCountRequiresAFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Count([]int{1}, "not a func")
	}, "requires a func as the second arg")
}

func TestCountRequiresASingleArgFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Count([]int{1}, func() {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.Count([]int{1}, func(i, j int) {})
	}, "requires a single arg function")
}

func TestCountRequiresASingleReturnFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Count([]int{1}, func(i int) {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.Count([]int{1}, func(i int) (int, int) { return 1, 0 })
	}, "requires a single arg function")
}

func TestCountRequiresFuncToTakeSliceArg(t *testing.T) {
	assert.Panics(t, func() {
		E.Count([]int{1}, func(s string) int { return 0 })
	}, "requires the func to take the same type as the slice")
}

func TestCountRequiresFuncToReturnABool(t *testing.T) {
	assert.Panics(t, func() {
		E.Count([]int{1}, func(i int) int { return i })
	}, "requires the func to take the same type as the slice")
}
