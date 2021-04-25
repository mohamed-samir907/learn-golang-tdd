package integers

import "testing"

func TestAdder(t *testing.T) {
	assertEqual := func(t testing.TB, got, expected int) {
		t.Helper()
		if got != expected {
			t.Errorf("got %d, expected %d", got, expected)
		}

	}

	t.Run("it can add two numbers", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5

		assertEqual(t, sum, expected)
	})

	t.Run("run example add", func(t *testing.T) {
		ExampleAdd()
	})
}
