package integers

import "testing"

func TestAdder(t *testing.T) {
	assertEqual := func(t testing.TB, got, expected int) {
		t.Helper()
		if got != expected {
			t.Errorf("got %d, expected %d", got, expected)
		}

	}

	sum := Add(2, 3)
	expected := 5

	assertEqual(t, sum, expected)
}
