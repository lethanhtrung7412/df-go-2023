package utils

import "testing"

func TestReName(t *testing.T) {
	testCases := []struct {
		args         []string
		expectedName string
	}{
		{[]string{"Trung", "Thanh", "Le", "VN"}, "Le Thanh Trung"},
		{[]string{"Trung", "Thanh", "Nguyen", "Tran", "Le", "VN"}, "Le Thanh Nguyen Tran Trung"},
		{[]string{"Trung", "Le", "VN"}, "Le Trung"},
		{[]string{"Trung", "Thanh", "Le", "US"}, "Trung Thanh Le"},
		{[]string{"Trung", "Le", "US"}, "Trung Le"},
		{[]string{"Trung", "Thanh", "Le", "VNN"}, ""},
	}

	for _, testCase := range testCases {
		actualName, _ := ReName(testCase.args)

		if actualName != testCase.expectedName {
			t.Errorf("Expected %s, got %s", testCase.expectedName, actualName)
		}
	}

}
