package utils

import "fmt"

func CheckCountryCode(args []string) (string, bool) {
	countryCode := args[len(args)-1]
	if !isValidCountryCode(countryCode) {
		fmt.Printf("Country with this code %s is not valid", countryCode)
		return "", false
	}

	return countryCode, true
}

func isValidCountryCode(countryCode string) bool {
	supportCountries := []string{"VN", "US"}
	for _, ctCode := range supportCountries {
		if ctCode == countryCode {
			return true
		}
	}

	return false
}
