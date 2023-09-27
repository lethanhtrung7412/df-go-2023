package cmd

import (
	"errors"
	"flag"

	"github.com/lethanhtrung7412/df-go-2023/ex02/utils"
)

func Execute() (string, error) {
	flag.Parse()
	args := flag.Args()
	if err := ValidateArguments(args); err != nil {
		return "", err
	}

	sortedString, err := getSortedString(args)
	if err != nil {
		return "", err
	}
	return sortedString, nil
}

var (
	intType    bool
	stringType bool
	mixType    bool
)

func init() {
	flag.BoolVar(&intType, string(utils.INTEGER), false, "Type of input array is integer")
	flag.BoolVar(&stringType, string(utils.STRING), false, "Type of input array is string")
	flag.BoolVar(&mixType, string(utils.MIX), false, "Type of input array is a mix of primitive types")
}

func ValidateArguments(args []string) error {
	if len(args) < 1 {
		return errors.New("requires at least 1 argument, found 0")
	}

	if err := utils.ValidateNeedOneFlag(); err != nil {
		return err
	}

	return nil
}

func getSortedString(args []string) (string, error) {
	var sortedString string
	var err error
	inputType := utils.GetInputType()
	switch inputType {
	case utils.INTEGER:
		sortedString, err = utils.IntSort(args)
		if err != nil {
			return "", err
		}
	case utils.STRING:
		sortedString, err = utils.StringSort(args)
		if err != nil {
			return "", err
		}
	case utils.MIX:
		sortedString, err = utils.StringMix(args)
		if err != nil {
			return "", err
		}
	}

	return sortedString, nil
}
