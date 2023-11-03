package simplecalc

import "testing"

// TestSum tests the Sum function.
func TestAdd(t *testing.T) {
	got := Add(1, 2)
	want := 3.0

	if got != want {
		t.Errorf("Sum(1, 2) = %f; want %f", got, want)
	}
}
