package enumerable_test

import (
	"fmt"
	"strings"
	"testing"

	E "github.com/darkhelmet/enumerable"
	"github.com/stretchr/testify/assert"
)

func ExampleReject() {
	words := strings.Fields("To be or not to be, that is the question!")
	values := E.Reject(words, func(s string) bool {
		return len(s) > 2
	})
	fmt.Println(values)
	// Output: [To be or to is]
}

func TestRejectWorks(t *testing.T) {
	ints := []int{2, 4, 6, 8}
	values := E.Reject(ints, func(i int) bool {
		return i > 5
	})
	assert.Equal(t, []int{2, 4}, values)

	values = E.Reject(ints, func(i int) bool {
		return i < 0
	})
	assert.Equal(t, values, values)
}

func TestRejectRequiresSlice(t *testing.T) {
	assert.Panics(t, func() {
		E.Reject(1, "doesn't matter")
	}, "requires a slice as the first arg")
}

func TestRejectRequiresAFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Reject([]int{1}, "not a func")
	}, "requires a func as the second arg")
}

func TestRejectRequiresASingleArgFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Reject([]int{1}, func() {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.Reject([]int{1}, func(i, j int) {})
	}, "requires a single arg function")
}

func TestRejectRequiresASingleReturnFunc(t *testing.T) {
	assert.Panics(t, func() {
		E.Reject([]int{1}, func(i int) {})
	}, "requires a single arg function")

	assert.Panics(t, func() {
		E.Reject([]int{1}, func(i int) (int, int) { return 1, 0 })
	}, "requires a single arg function")
}

func TestRejectRequiresFuncToTakeSliceArg(t *testing.T) {
	assert.Panics(t, func() {
		E.Reject([]int{1}, func(s string) int { return 0 })
	}, "requires the func to take the same type as the slice")
}

func TestRejectRequiresFuncToReturnABool(t *testing.T) {
	assert.Panics(t, func() {
		E.Reject([]int{1}, func(i int) int { return i })
	}, "requires the func to take the same type as the slice")
}
