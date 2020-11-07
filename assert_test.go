package assert_test

import (
	"github.com/hedzr/assert"
	"os"
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

	// assert.Equal(t, expected, actual)
	assert.Equal(t, actual, actual)
}

func TestEqualTrue(t *testing.T) {
	assert.EqualTrue(t, true)
	assert.EqualFalse(t, false)
}

func TestErrors(t *testing.T) {
	f, err := os.Open("/tmp/not-exist")
	assert.Error(t, err) // err 应该是一个 notfound 错误
	defer f.Close()

	f, err = os.Open("/etc/passwd")
	assert.NoError(t, err) // err 应该为 nil
	defer f.Close()
}

func TestNilObjects(t *testing.T) {
	src := "source"
	var ch *string
	assert.Nil(t, ch)
	ch = &src
	assert.NotNil(t, ch)
}

func TestStrings(t *testing.T) {
	var s = "365 days"
	assert.Match(t, s, `\d[ ]*days`)
}

func TestPanics(t *testing.T) {
	fn := func() {
		panic("omg omg omg!")
	}

	assert.PanicMatches(t, func() { fn() }, "omg omg omg!")
	assert.PanicMatches(t, func() { panic("omg omg omg!") }, "omg omg omg!")
}
