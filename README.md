# assert

`assert` provides a set of assertion helpers for unit/bench testing in golang.

`assert` is inspired by these projects:

- <https://github.com/alecthomas/assert>
- <https://github.com/go-playground/assert>
- <https://github.com/stretchr/testify> assert, mock, suite ...

### improvements

- Can be used with both unit test and bench test.
- Most of conventiional assertions: 
  - Equal, NotEqual, EqualTrue, EqualFalse
  - Nil, NotNil
  - Error, NoError
  - PanicMatches: test for the function which might throw a panic
  - Match, NotMatch: compares a value with regexp test 
- Fresh coding in go 1.13~1.15 and later.

### Short guide

```go
package some_test

import (
	"github.com/hedzr/assert"
	"testing"
)

type Person struct {
	Name string
	Age  int
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
``` 

### LICENSE

MIT
