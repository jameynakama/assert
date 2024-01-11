package assert

import (
	"fmt"
)

type Test interface {
	Helper()
	Errorf(format string, args ...interface{})
}

func getMsgFormatter(v any) string {
	switch v.(type) {
	case string:
		return "%q"
	default:
		return "%+v"
	}
}

func Equal(t Test, expected, actual any, msg ...string) {
	t.Helper()

	// TODO: Determine if the expected and actual values can be compared using
	// the == operator. If not, then we need to use the fmt.Sprintf() function
	// to format the values for comparison, or use reflect.DeepEqual.

	if actual != expected {
		if len(msg) > 0 {
			t.Errorf(msg[0], expected, actual)
		} else {
			formatterForActual := getMsgFormatter(actual)
			formatterForExpected := getMsgFormatter(expected)
			toBeFormatted := fmt.Sprintf("Expected %s, got %s instead", formatterForExpected, formatterForActual)
			t.Errorf(toBeFormatted, expected, actual)
		}
	}
}
