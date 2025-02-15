package arrays

import "testing"

func TestTotal(t *testing.T) {
	numbers := [5]int{1,2,3,4,5}
	result := Total(numbers)
	expected := 15

	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}
