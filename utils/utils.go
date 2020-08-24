package utils

import (
	"strconv"
	"strings"
)

// 字符 ball 转换为 数值
func BollStrToNum(sss string) (ret []int) {
	for _, s := range strings.Split(sss, " ") {
		j, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ret = append(ret, j)
	}
	return
}

func IntersectionInt(a, b *[]int) *[]int {

	m := make(map[int]uint8)
	for _, k := range *a {
		m[k] |= (1 << 0)
	}
	for _, k := range *b {
		m[k] |= (1 << 1)
	}
	var inAAndB []int
	// var inAAndB, inAButNotB, inBButNotA []int
	for k, v := range m {
		a := v&(1<<0) != 0
		b := v&(1<<1) != 0

		switch {
		case a && b:
			inAAndB = append(inAAndB, k)
			// case a && !b:
			// 	inAButNotB = append(inAButNotB, k)
			// case !a && b:
			// 	inBButNotA = append(inBButNotA, k)
		}
	}
	return &inAAndB
}

// 阶乘
func Factorial(max, min int) int {
	if max >= min && max > 1 {
		return max * Factorial(max-1, min)
	} else {
		return 1
	}
}

// 排列组合
func Combine(m, n int) int {
	if m < n || n < 0 {
		return 0
	}
	return Factorial(m, m-n+1) / Factorial(n, 1)
}
