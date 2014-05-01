package enumerable_test

import (
	"fmt"
	"strings"
	"testing"

	E "github.com/darkhelmet/enumerable"
	"github.com/stretchr/testify/assert"
)

func ExampleNone() {
	words := strings.Fields("To be or not to be, that is the question!")
	okay := E.None(words, func(s string) bool {
		return len(s) > 8
	})
	fmt.Println(okay)
	// Output: false
}

func TestNoneWorks(t *testing.T) {
	ints := []int{2, 4, 6, 8}
	aboveFive := E.None(ints, func(i int) bool {
		return i > 5
	})
	assert.False(t, aboveFive)

	negative := E.None(ints, func(i int) bool {
		return i < 0
	})
	assert.True(t, negative)
}

func TestNoneRequiresSlice(t *testing.T) {
	assert.Panics(t, func() {
		E.None(1, "doesn't matter")
	}, "requires a slice as the first arg")
}

func TestNoneRequiresAFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.None([]int{1}, "not a func")
	}, "requires a func as the second arg")
}

func TestNoneRequiresASingleArgFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.None([]int{1}, func() {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.None([]int{1}, func(i, j int) {})
	}, "requires a single arg function")
}

func TestNoneRequiresASingleReturnFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.None([]int{1}, func(i int) {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.None([]int{1}, func(i int) (int, int) { return 1, 0 })
	}, "requires a single arg function")
}

func TestNoneRequiresFuncToTakeSliceArg(t *testing.T) {
	assert.Panics(t, func() {
		E.None([]int{1}, func(s string) int { return 0 })
	}, "requires the func to take the same type as the slice")
}

func TestNoneRequiresFuncToReturnABool(t *testing.T) {
	assert.Panics(t, func() {
		E.None([]int{1}, func(i int) int { return i })
	}, "requires the func to take the same type as the slice")
}
