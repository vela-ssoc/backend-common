package model

import (
	"strconv"
	"strings"
)

// Semver1 语义化版本 https://semver.org/lang/zh-CN/
// major.minor.patch
type Semver1 string

// Int64 将版本转换为数值用于方便比较大小
// major, minor, patch 三项都要小于 step 才有意义，即：最大版本号是：99999.99999.99999
func (ver Semver1) Int64() int64 {
	const step = 100000
	str := string(ver)

	elems := strings.Split(str, ".")
	end := len(elems) - 1
	nums := make([]int64, 3)
	for i, elem := range elems {
		if i > 2 { // 只取前 3 位作为有效的计算
			break
		}
		if i == end {
			elem, _, _ = strings.Cut(elem, "-")
		}
		n, _ := strconv.ParseInt(elem, 10, 64)
		nums[i] = n
	}

	var ret int64
	for _, num := range nums {
		ret *= step
		ret += num
	}

	return ret
}

// String fmt.Stringer
func (ver Semver1) String() string {
	return string(ver)
}
