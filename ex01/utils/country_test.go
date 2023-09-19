package utils

import "testing"

func TestCheckCountryCode(t *testing.T) {
	input := []string{"Doe", "John", "VN"}
	output, outputStatus := "VN", true

	ct, ok := CheckCountryCode(input)

	if output != ct || outputStatus != ok {
		t.Errorf("Expected %s, %v; Have %s, %v", output, outputStatus, ct, ok)
	}

	wrongInput := []string{"Trung", "Le", "DE"}
	wrongOutput, wrongOutputStatus := "", false

	testCt, status := CheckCountryCode(wrongInput)
	if wrongOutput != testCt || wrongOutputStatus != status {
		t.Errorf("Expected %s, %v; Have %s, %v", wrongOutput, wrongOutputStatus, testCt, status)
	}
}
