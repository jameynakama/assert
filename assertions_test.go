package assert_test

import (
	"fmt"
	"testing"

	"jameynakama.com/assert"
)

type mockTestHelper struct {
	ErrorMessages []string
}

func (m *mockTestHelper) Helper() {}

func (m *mockTestHelper) Errorf(format string, args ...interface{}) {
	m.ErrorMessages = append(m.ErrorMessages, fmt.Sprintf(format, args...))
}

func assertFailure(t *testing.T, expErr string, actErrs []string) {
	t.Helper()

	if len(actErrs) != 1 {
		t.Errorf("Expected 1 error message, got %d instead", len(actErrs))
	} else {
		if actErrs[0] != expErr {
			t.Errorf("Expected error message to be %q, got %q instead", expErr, actErrs[0])
		}
	}
}

func TestEqual(t *testing.T) {
	t.Run("DefaultArgs", func(t *testing.T) {
		assert.Equal(t, 1, 1)
	})

	t.Run("DefaultArgsFailure", func(t *testing.T) {
		expectedMsg := "Expected \"one\", got \"two\" instead"
		mt := &mockTestHelper{}
		assert.Equal(mt, "one", "two")
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})

	t.Run("DefaultArgsFloatFailure", func(t *testing.T) {
		expectedMsg := "Expected 1.5, got 2.5 instead"
		mt := &mockTestHelper{}
		assert.Equal(mt, 1.5, 2.5)
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})

	t.Run("CustomMsgArgNoFormatters", func(t *testing.T) {
		expectedMsg := "Custom message%!(EXTRA int=1, int=2)"
		mt := &mockTestHelper{}
		assert.Equal(mt, 1, 2, "Custom message")
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})

	t.Run("CustomMsgArgWithFormatters", func(t *testing.T) {
		expectedMsg := "Wanted 1, got 2 instead"
		mt := &mockTestHelper{}
		assert.Equal(mt, 1, 2, "Wanted %d, got %d instead")
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})
}