# Simple Assertions

This is a learning project for me. I'm sure there are more robust mods out there
that do the same thing, but better.

## Usage

The signature for most assertions will take the testing object as the first
argument, followed by the "actual" and "expected" values to compare, and
finally, an optional failure message that can be loaded with string formatters
in case the default error messages aren't to your liking.

Generally speaking, if the assertion fails, it will call `t.Errorf` with the
default or custom message.

```go
import (
    "testing"
    "github.com/jameynakama/assert"
)

func TestOneIsOne(t *testing.T) {
    assert.Equal(t, 1, 1)
}

func TestOneIsNotOne(t *testing.T) {
    assert.NotEqual(t, "one", "two")
}

func TestWithCustomMessage(t *testing.T) {
    assert.Equal(t, []byte{'a', 'b'}, []byte{'c', 'd'}, "I wanted %v, but I got %v instead!")
}
```

## License

This project is licensed under the terms of the MIT license.
