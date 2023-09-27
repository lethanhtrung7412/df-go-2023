package utils

import "flag"

type Input string

const (
	INTEGER Input = "int"
	STRING  Input = "string"
	MIX     Input = "mix"
)

var SupportedTypesInput = []Input{INTEGER, STRING, MIX}

func GetInputType() Input {
	for _, t := range SupportedTypesInput {
		if isInputTypeSet(t) {
			return t
		}
	}

	return MIX
}

func ValidateNeedOneFlag() error {
	count := 0
	for _, inputType := range SupportedTypesInput {
		if isInputTypeSet(inputType) {
			count++
		}
	}

	if count < 0 {
		return InputTypeMissingError
	}

	if count > 1 {
		return MoreThanOneInputTypeError
	}

	return nil
}

func isInputTypeSet(inputType Input) bool {
	if t := flag.Lookup(string(inputType)); t != nil && t.Value.String() == "true" {
		return true
	}
	return false
}
