package utils

import "strings"

func ReName(args []string) (string, bool) {
	ctCode, ok := CheckCountryCode(args)
	if !ok {
		return "", false
	}
	firstName, middleName, lastName := getNameParts(args)

	if middleName == "" {
		return getName(ctCode, firstName, lastName), true
	}

	return getName(ctCode, firstName, middleName, lastName), true
}

func getName(ct string, nameParts ...string) string {
	if ct == "US" {
		return strings.Join(nameParts, " ")
	}

	return strings.Join(ReverseSlice(nameParts), " ")
}

func getNameParts(args []string) (firstName, middleName, lastName string) {
	firstName = args[0]
	lastName = args[len(args)-2]
	middleName = strings.Join(args[1:len(args)-2], " ")
	return
}
