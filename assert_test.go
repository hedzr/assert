package assert_test

import (
	"github.com/hedzr/assert"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestDiff(t *testing.T) {
	expected := []*Person{{"Alec", 20}, {"Bob", 21}, {"Sally", 22}}
	actual := []*Person{{"Alex", 20}, {"Bob", 22}, {"Sally", 22}}
	assert.NotEqual(t, expected, actual)
	t.Log(assert.DiffValues(expected, actual))
}

func TestEqual(t *testing.T) {
	expected := []*Person{{"Alec", 20}, {"Bob", 21}, {"Sally", 22}}
	actual := []*Person{{"Alex", 20}, {"Bob", 22}, {"Sally", 22}}
	assert.NotEqual(t, expected, actual)

	assert.Equal(t, actual, actual)
}

func TestEqualTrue(t *testing.T) {
	assert.EqualTrue(t, true)
	assert.EqualFalse(t, false)
}
