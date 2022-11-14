package helpers

import (
	"math"
	"strconv"
	"strings"

	"github.com/realtemirov/tasks/project5/consts"
)

type SubNum struct {
	R int
	B int
}

func Checker(number *string) error {

	allowed := "1234567890"

	str := *number
	isPositive := true

	if str[0] == '-' {
		isPositive = false
	}

	if !strings.Contains("+-"+allowed, string(str[0])) {
		return consts.ErrorBadInput
	}

	for i := 1; i < len(str); i++ {
		if !strings.Contains(allowed, string(str[i])) {
			return consts.ErrorBadInput
		}
	}

	str = ClearFirstSymbol(str)

	if str != "0" {
		for str[0] == '0' {
			if str == "0" {
				break
			}
			str = str[1:]
		}

		if !isPositive && str != "0" {
			str = "-" + str[:]
		}
	}
	*number = str
	return nil
}
func ClearFirstSymbol(str string) string {
	if str[0] == '-' || str[0] == '+' {
		return str[1:]
	} else {
		return str
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func FillWithZero(a, b string) (string, string) {
	minLength := math.Min(float64(len(a)), float64(len(b)))
	min_s := ""
	max_s := ""
	if len(a) == int(minLength) {
		min_s = a
		max_s = b
	} else {
		min_s = b
		max_s = a
	}
	for len(min_s) != len(max_s) {
		min_s = "0" + min_s
	}
	return max_s, min_s
}
func SubDigit(c1, c2 string, b int) SubNum {
	var res SubNum
	n1, _ := strconv.Atoi(c1)
	n2, _ := strconv.Atoi(c2)
	n := n1 - n2 - b

	if n >= 0 {
		res.R = n
		res.B = 0
	} else {
		res.R = n + 10
		res.B = 1
	}
	return res
}
