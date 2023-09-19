package utils

func ReverseSlice(args []string) []string {
	for i, j := 0, len(args)-1; i < j; i, j = i+1, j-1 {
		args[i], args[j] = args[j], args[i]
	}

	return args
}
