package assert_test

import (
	"fmt"
	"testing"

	"github.com/jameynakama/assert"
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

	t.Run("DefaultArgsByteSlice", func(t *testing.T) {
		assert.Equal(t, []byte{'a', 'b'}, []byte{'a', 'b'})
	})

	t.Run("DefaultArgsFailure", func(t *testing.T) {
		expectedMsg := "Expected \"two\", got \"one\" instead"
		mt := &mockTestHelper{}
		assert.Equal(mt, "one", "two")
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})

	t.Run("DefaultArgsFloatFailure", func(t *testing.T) {
		expectedMsg := "Expected 2.5, got 1.5 instead"
		mt := &mockTestHelper{}
		assert.Equal(mt, 1.5, 2.5)
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})

	t.Run("CustomMsgArgNoFormatters", func(t *testing.T) {
		expectedMsg := "Custom message%!(EXTRA int=2, int=1)"
		mt := &mockTestHelper{}
		assert.Equal(mt, 1, 2, "Custom message")
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})

	t.Run("CustomMsgArgWithFormatters", func(t *testing.T) {
		expectedMsg := "Wanted 2, got 1 instead"
		mt := &mockTestHelper{}
		assert.Equal(mt, 1, 2, "Wanted %d, got %d instead")
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})
}

func TestNotEqual(t *testing.T) {
	t.Run("DefaultArgs", func(t *testing.T) {
		assert.NotEqual(t, 1, 2)
	})

	t.Run("DefaultArgsByteSlice", func(t *testing.T) {
		assert.NotEqual(t, []byte{'a', 'b'}, []byte{'c', 'd'})
	})

	t.Run("DefaultArgsFailure", func(t *testing.T) {
		expectedMsg := "Expected different values, but got \"two\" for both"
		mt := &mockTestHelper{}
		assert.NotEqual(mt, "two", "two")
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})

	t.Run("DefaultArgsFloatFailure", func(t *testing.T) {
		expectedMsg := "Expected different values, but got 1.5 for both"
		mt := &mockTestHelper{}
		assert.NotEqual(mt, 1.5, 1.5)
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})

	t.Run("CustomMsgArgNoFormatters", func(t *testing.T) {
		expectedMsg := "Custom message%!(EXTRA int=2)"
		mt := &mockTestHelper{}
		assert.NotEqual(mt, 2, 2, "Custom message")
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})

	t.Run("CustomMsgArgWithFormatters", func(t *testing.T) {
		expectedMsg := "Wanted different, got same instead: 2"
		mt := &mockTestHelper{}
		assert.NotEqual(mt, 2, 2, "Wanted different, got same instead: %d")
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})
}

func TestInSlice(t *testing.T) {
	t.Run("DefaultArgs", func(t *testing.T) {
		assert.InSlice(t, []string{"one", "two"}, "two")
	})

	t.Run("DefaultArgsByteSlice", func(t *testing.T) {
		assert.InSlice(t, [][]byte{{'a', 'b'}, {'c', 'd'}}, []byte{'c', 'd'})
	})

	t.Run("DefaultArgsFailure", func(t *testing.T) {
		expectedMsg := "Expected \"three\" to be in slice [one two]"
		mt := &mockTestHelper{}
		assert.InSlice(mt, []string{"one", "two"}, "three")
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})

	t.Run("DefaultArgsFloatFailure", func(t *testing.T) {
		expectedMsg := "Expected 3.5 to be in slice [1.5 2.5]"
		mt := &mockTestHelper{}
		assert.InSlice(mt, []float64{1.5, 2.5}, 3.5)
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})

	t.Run("CustomMsgArgNoFormatters", func(t *testing.T) {
		expectedMsg := "Custom message%!(EXTRA int=3)"
		mt := &mockTestHelper{}
		assert.InSlice(mt, []int{1, 2}, 3, "Custom message")
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})

	t.Run("CustomMsgArgWithFormatters", func(t *testing.T) {
		expectedMsg := "Wanted 3 to be in slice%!(EXTRA int=3)"
		mt := &mockTestHelper{}
		assert.InSlice(mt, []int{1, 2}, 3, "Wanted 3 to be in slice")
		assertFailure(t, expectedMsg, mt.ErrorMessages)
	})
}
