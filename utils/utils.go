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
