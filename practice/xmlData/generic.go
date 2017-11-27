package main

import (
	"fmt"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func difference(slice1 []int64, slice2 []int64) ([]string) {
	var diffStr []string
	m := map[string]int{}

	for _, s1Val := range slice1 {
		m[fmt.Sprint(s1Val)] = 1
	}
	for _, s2Val := range slice2 {
		m[fmt.Sprint(s2Val)] = m[fmt.Sprint(s2Val)] + 1
	}

	for mKey, mVal := range m {
		if mVal == 1 {
			diffStr = append(diffStr, mKey)
		}
	}

	return diffStr
}

func reduce(value string) (int64, error) {
	ival, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return ival, err
	}
	if ival > 9 {
		val1 := fmt.Sprint(ival)[0:1]
		val2 := fmt.Sprint(ival)[1:]
		ival1, err := strconv.ParseInt(val1, 10, 0)
		if err != nil {
			return ival, err
		}
		ival2, err := strconv.ParseInt(val2, 10, 0)
		if err != nil {
			return ival, err
		}
		return reduce(fmt.Sprint(ival1 + ival2))
	}
	return ival, nil
}
