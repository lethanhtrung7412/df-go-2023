package utils

import "testing"

func TestReverseSlice(t *testing.T) {
	input := []string{"Doe", "John"}
	output := []string{"John", "Doe"}

	slice := ReverseSlice(input)

	if len(slice) != len(output) {
		t.Errorf("Expected %v, got %v", output, slice)

	}
	for i := range output {
		if slice[i] != output[i] {
			t.Errorf("Expected %v, got %v", output, slice)
		}
	}
}
