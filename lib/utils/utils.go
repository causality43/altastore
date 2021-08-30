package utils

import (
	"reflect"
	"strconv"
)

func CompareStrings(a string, b string) string {
	if b == "" {
		return a
	} else {
		return b
	}
}

func CompareId(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return b
	}
}

func IsInt(i interface{}) bool {
	var a int = 30
	if reflect.TypeOf(i) == reflect.TypeOf(a) {
		return true
	} else {
		return false
	}
}

func StringIsNotNumber(a string) bool {
	_, err := strconv.Atoi(a)
	if err != nil {
		return false
	} else {
		return true
	}
}
