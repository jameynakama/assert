package assert

import (
	"fmt"
	"reflect"
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

func Equal[T any](t Test, actual, expected T, msg ...string) {
	t.Helper()

	if !reflect.DeepEqual(actual, expected) {
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
