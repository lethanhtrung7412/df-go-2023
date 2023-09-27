package utils

import (
	"fmt"
	"sort"
	"strconv"
)

func IntSort(args []string) (string, error) {
	var err error

	intArr := make([]int, len(args))
	for i := 0; i < len(args); i++ {
		if intArr[i], err = strconv.Atoi(args[i]); err != nil {
			return "", fmt.Errorf("invalid argument: %s", args[i])
		}
	}

	return getSortIntSliceToString(intArr)
}

func getSortIntSliceToString(args []int) (string, error) {
	sort.Ints(args)

	rs := ""
	for _, v := range args {
		rs += fmt.Sprintf("%d ", v)
	}

	return rs, nil
}

func StringSort(args []string) (string, error) {
	var err error
	stringArr := make([]float64, len(args))
	for i := 0; i < len(args); i++ {
		if stringArr[i], err = strconv.ParseFloat(args[i], 64); err == nil {
			return "", fmt.Errorf("invalid argument: %s", args[i])
		}
	}

	return getStringSort(args)
}

func StringMix(args []string) (string, error) {
	return getStringSort(args)
}

func getStringSort(args []string) (string, error) {
	sort.Strings(args)

	rs := ""
	for _, v := range args {
		rs += fmt.Sprintf("%s ", v)
	}

	return rs, nil
}
