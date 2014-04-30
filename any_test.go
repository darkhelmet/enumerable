package enumerable_test

import (
	"fmt"
	"strings"
	"testing"

	E "github.com/darkhelmet/enumerable"
	"github.com/stretchr/testify/assert"
)

func ExampleAny() {
	words := strings.Fields("To be or not to be, that is the question!")
	okay := E.Any(words, func(s string) bool {
		return len(s) > 2
	})
	fmt.Println(okay)
	// Output: true
}

func TestAnyWorks(t *testing.T) {
	ints := []int{2, 4, 6, 8}
	aboveFive := E.Any(ints, func(i int) bool {
		return i > 5
	})
	assert.True(t, aboveFive)

	negative := E.Any(ints, func(i int) bool {
		return i < 0
	})
	assert.False(t, negative)
}

func TestAnyRequiresSlice(t *testing.T) {
	assert.Panics(t, func() {
		E.Any(1, "doesn't matter")
	}, "requires a slice as the first arg")
}

func TestAnyRequiresAFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Any([]int{1}, "not a func")
	}, "requires a func as the second arg")
}

func TestAnyRequiresASingleArgFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Any([]int{1}, func() {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.Any([]int{1}, func(i, j int) {})
	}, "requires a single arg function")
}

func TestAnyRequiresASingleReturnFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Any([]int{1}, func(i int) {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.Any([]int{1}, func(i int) (int, int) { return 1, 0 })
	}, "requires a single arg function")
}

func TestAnyRequiresFuncToTakeSliceArg(t *testing.T) {
	assert.Panics(t, func() {
		E.Any([]int{1}, func(s string) int { return 0 })
	}, "requires the func to take the same type as the slice")
}

func TestAnyRequiresFuncToReturnABool(t *testing.T) {
	assert.Panics(t, func() {
		E.Any([]int{1}, func(i int) int { return i })
	}, "requires the func to take the same type as the slice")
}
