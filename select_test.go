package enumerable_test

import (
	"fmt"
	"strings"
	"testing"

	E "github.com/darkhelmet/enumerable"
	"github.com/stretchr/testify/assert"
)

func ExampleSelect() {
	words := strings.Fields("To be or not to be, that is the question!")
	values := E.Select(words, func(s string) bool {
		return len(s) > 2
	})
	fmt.Println(values)
	// Output: [not be, that the question!]
}

func TestSelectWorks(t *testing.T) {
	ints := []int{2, 4, 6, 8}
	values := E.Select(ints, func(i int) bool {
		return i > 5
	})
	assert.Equal(t, []int{6, 8}, values)

	values = E.Select(ints, func(i int) bool {
		return i < 0
	})
	assert.Equal(t, []int{}, values)
}

func TestSelectRequiresSlice(t *testing.T) {
	assert.Panics(t, func() {
		E.Select(1, "doesn't matter")
	}, "requires a slice as the first arg")
}

func TestSelectRequiresAFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Select([]int{1}, "not a func")
	}, "requires a func as the second arg")
}

func TestSelectRequiresASingleArgFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Select([]int{1}, func() {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.Select([]int{1}, func(i, j int) {})
	}, "requires a single arg function")
}

func TestSelectRequiresASingleReturnFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Select([]int{1}, func(i int) {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.Select([]int{1}, func(i int) (int, int) { return 1, 0 })
	}, "requires a single arg function")
}

func TestSelectRequiresFuncToTakeSliceArg(t *testing.T) {
	assert.Panics(t, func() {
		E.Select([]int{1}, func(s string) int { return 0 })
	}, "requires the func to take the same type as the slice")
}

func TestSelectRequiresFuncToReturnABool(t *testing.T) {
	assert.Panics(t, func() {
		E.Select([]int{1}, func(i int) int { return i })
	}, "requires the func to take the same type as the slice")
}
