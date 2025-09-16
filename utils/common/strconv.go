package common

import "strconv"

func StringToUint(s string) (uint, error) {
	n, err := strconv.ParseUint(s, 10, 64) // 10表示十进制，64是最大位数
	if err != nil {
		return 0, err
	}
	return uint(n), nil
}
