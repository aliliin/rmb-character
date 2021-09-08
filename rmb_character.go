package rmb_character

import (
	"fmt"
	"strconv"
	"strings"
)

var capitalNumbers = [10]string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}

var integerUnits = [4]string{"", "拾", "佰", "仟"}

var placeUnits = [4]string{"", "万", "亿", "兆"}

var decimalUnits = [4]string{"角", "分", "厘", "毫"}

func RmbCapitalCharacters(price int) string {
	if price <= 0 {
		return "输入金额请大于 0 "
	}

	// 翻转一下数字，从个位开始
	num1, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(price)/float64(100)), 64) // 保留2位小数
	strScore := strconv.FormatFloat(num1, 'f', 2, 64)
	s := strings.Split(strScore, ".")

	integerStr := s[0]
	decimal := s[1]

	resString := reverse(integerStr)
	var res []string
	var last int32
	for i, s := range chunk(resString) {
		res = append([]string{placeUnits[i]}, res...)
		for k, i3 := range s {
			if i3 == 48 && (last == 48 || k == 0) {
				last = i3
				continue
			}
			last = i3
			if i3 != 48 {
				res = append([]string{integerUnits[k]}, res...)
			}
			x, _ := strconv.Atoi(string(i3))
			res = append([]string{capitalNumbers[x]}, res...)
		}
	}

	if len(res) == 0 {
		res = append([]string{capitalNumbers[0]}, res...)
	}
	res = append(res, "圆")
	// 没有小数就加整
	if decimal == "00" {
		res = append(res, "整")
	} else {
		// 转换小数位
		for key, number := range decimal {
			x, _ := strconv.Atoi(string(number))
			res = append(res, capitalNumbers[x])
			res = append(res, decimalUnits[key])
		}
	}

	var resStr string
	for _, re := range res {
		resStr = resStr + re
	}
	return resStr
}

func chunk(integer string) []string {
	l := len(integer)
	n := l / 4
	if l%4 != 0 {
		n++
	}
	b := make([]string, n)
	for i := 0; i < n; i++ {
		start := i * 4
		end := (i + 1) * 4
		if end > l {
			end = l
		}
		b[i] = string(integer[start:end])
	}
	return b
}

func reverse(s string) string {
	a := func(s string) *[]rune {
		var b []rune
		for _, k := range []rune(s) {
			defer func(v rune) {
				b = append(b, v)
			}(k)
		}
		return &b
	}(s)
	return string(*a)
}
