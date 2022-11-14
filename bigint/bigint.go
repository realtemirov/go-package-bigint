package bigint

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/realtemirov/tasks/project5/consts"
	"github.com/realtemirov/tasks/project5/helpers"
)

type bigInt struct {
	value string
}

func NewInt(num string) (bigInt, error) {

	if err := helpers.Checker(&num); err != nil {
		return bigInt{value: num}, errors.New("bad input")
	}

	return bigInt{value: num}, nil
}

func Add(a, b bigInt) bigInt {

	var (
		t            = 0
		res          = ""
		a_isPositive = true
		b_isPositive = true
	)

	if a.value[0] == '-' {
		a_isPositive = false
	}

	if b.value[0] == '-' {
		b_isPositive = false
	}

	a = a.Abs()
	b = b.Abs()

	if a_isPositive != b_isPositive {
		res_sub := Sub(a, b)
		if Max(a, b).value[0] != '-' {
			res_sub.value = res_sub.value[1:]
		}
		return res_sub
	}

	l := int(math.Max(float64(len(a.value)), float64(len(b.value))))
	var a_s = strings.Split(helpers.Reverse(a.value), "")
	var b_s = strings.Split(helpers.Reverse(b.value), "")

	for i := 0; i < l; i++ {
		var pa, pb int

		if len(a_s) > i {
			pa, _ = strconv.Atoi(a_s[i])
		} else {
			pa = 0
		}

		if len(b_s) > i {
			pb, _ = strconv.Atoi(b_s[i])
		} else {
			pb = 0
		}

		k := (pa + pb + t)
		t = k / 10
		res = strconv.Itoa(k%10) + res
	}

	if t != 0 {
		res = strconv.Itoa(t) + res
	}
	if !a_isPositive {
		res = "-" + res
	}
	return bigInt{value: res}
}

func Sub(a, b bigInt) bigInt {

	a_s := Max(a, b).value
	b_s := Min(a, b).value
	a_s, b_s = helpers.FillWithZero(a_s, b_s)
	borrow := 0
	var res bigInt
	fmt.Println(a_s)
	fmt.Println(b_s)
	for i := len(a_s) - 1; i >= 0; i-- {

		subRes := helpers.SubDigit(string(a_s[i]), string(b_s[i]), borrow)
		borrow = subRes.B
		res.value = strconv.Itoa(subRes.R) + res.value
	}
	if a.value != Max(a, b).value {
		res.value = "-" + res.value
	}
	return res
}
func Max(a, b bigInt) bigInt {
	res := a
	if len(a.value) > len(b.value) {
		res = a
	} else if len(a.value) < len(b.value) {
		res = b
	} else {
		for i := 0; i < len(a.value); i++ {
			a1, _ := strconv.Atoi(string(a.value[i]))
			b1, _ := strconv.Atoi(string(b.value[i]))
			if a1 > b1 {
				res = a
			} else if a1 < b1 {
				res = b
			}
		}
	}
	return res
}
func Min(a, b bigInt) bigInt {
	res := a
	if len(a.value) > len(b.value) {
		res = b
	} else if len(a.value) < len(b.value) {
		res = a
	} else {
		for i := 0; i < len(a.value); i++ {
			a1, _ := strconv.Atoi(string(a.value[i]))
			b1, _ := strconv.Atoi(string(b.value[i]))
			if a1 > b1 {
				res = b
			} else if a1 < b1 {
				res = a
			}
		}
	}
	return res
}

func (z *bigInt) Set(num string) error {
	if err := helpers.Checker(&num); err != nil {
		return consts.ErrorBadInput
	}

	z.value = num
	return nil
}
func (z *bigInt) Value() string {
	return z.value
}
func (z *bigInt) Abs() bigInt {

	val := z.value

	if val[0] == '-' {
		return bigInt{
			value: val[1:],
		}
	}
	return bigInt{
		value: val,
	}
}
