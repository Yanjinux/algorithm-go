package stringalog

import (
	"strconv"
	"unicode"
)

/*
 1 给定二个字符串str1 和 str2，如果str1 和 str2 中出现的字符串种类一样且每种字符出现的次数也一样，则str1和str2互为变形词。
 请实现判断函数
*/

func VariantWord(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var m = make(map[string]int)
	for i, _ := range s1 {
		v := s1[i : i+1]
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}

	for i, _ := range s2 {
		v := s2[i : i+1]
		if _, ok := m[v]; ok {

			m[v] = m[v] - 1
			if m[v] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

/*
	给定一个字符串str，求其中全部数字串所代表的数字之和
*/

func SubNumberSum(s string) int {
	var sum = 0
	var tmp string
	var flag, start = true, false
	for i, _ := range s {
		if !start && s[i] == '-' {
			flag = !flag
			continue
		}
		if unicode.IsDigit(rune(s[i])) {
			start = true
			tmp = tmp + string(s[i])
		} else {
			start = false
			ts, _ := strconv.Atoi(tmp)
			if !flag {
				sum = sum - ts
			} else {
				sum = sum + ts
			}
			flag = true
			tmp = ""
		}
	}

	ts, _ := strconv.Atoi(tmp)
	if !flag {
		sum = sum - ts
	} else {
		sum = sum + ts
	}

	return sum
}
