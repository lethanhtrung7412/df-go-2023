package utils

import "errors"

var InputTypeMissingError = errors.New("Need at least 1 flag")
var MoreThanOneInputTypeError = errors.New("Your input is more than 1 flag")
