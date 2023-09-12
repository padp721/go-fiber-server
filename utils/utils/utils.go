package utils

import "strconv"

func StringToInteger(str string) int {
	strInt, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return strInt
}
