package main

import "errors"

func testJoin(strs ...string) string {
	ret := ""
	for _, s := range strs {
		ret += s
	}
	return ret
}

func testDivision(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}
