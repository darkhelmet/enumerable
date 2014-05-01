package enumerable_test

import (
	"fmt"
	"strings"
	"testing"

	E "github.com/darkhelmet/enumerable"
	"github.com/stretchr/testify/assert"
)

func ExampleDetect() {
	words := strings.Fields("To be or not to be, that is the question!")
	value, ok := E.Detect(words, func(s string) bool {
		return len(s) > 2
	})
	fmt.Println(value, ok)
	// Output: not true
}

func ExampleDetect_notfound() {
	ints := []int{1, 2, 3, 4, 5}
	value, ok := E.Detect(ints, func(i int) bool {
		return i > 10
	})
	fmt.Println(value, ok)
	// Output: 0 false
}

func TestDetectWorks(t *testing.T) {
	ints := []int{2, 4, 6, 8}
	value, ok := E.Detect(ints, func(i int) bool {
		return i > 5
	})
	assert.True(t, ok)
	assert.Equal(t, 6, value)

	value, ok = E.Detect(ints, func(i int) bool {
		return i < 0
	})
	assert.False(t, ok)
	assert.Equal(t, 0, value)
}

func TestDetectRequiresSlice(t *testing.T) {
	assert.Panics(t, func() {
		E.Detect(1, "doesn't matter")
	}, "requires a slice as the first arg")
}

func TestDetectRequiresAFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Detect([]int{1}, "not a func")
	}, "requires a func as the second arg")
}

func TestDetectRequiresASingleArgFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Detect([]int{1}, func() {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.Detect([]int{1}, func(i, j int) {})
	}, "requires a single arg function")
}

func TestDetectRequiresASingleReturnFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Detect([]int{1}, func(i int) {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.Detect([]int{1}, func(i int) (int, int) { return 1, 0 })
	}, "requires a single arg function")
}

func TestDetectRequiresFuncToTakeSliceArg(t *testing.T) {
	assert.Panics(t, func() {
		E.Detect([]int{1}, func(s string) int { return 0 })
	}, "requires the func to take the same type as the slice")
}

func TestDetectRequiresFuncToReturnABool(t *testing.T) {
	assert.Panics(t, func() {
		E.Detect([]int{1}, func(i int) int { return i })
	}, "requires the func to take the same type as the slice")
}
