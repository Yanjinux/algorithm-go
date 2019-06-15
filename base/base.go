package base

import "fmt"

func bubbleSorting(s []int) {
	l := len(s)
	for ; l > 0; l-- {
		tl := len(s)
		for i, _ := range s[:tl-1] {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
		s = s[:tl-1]
	}
}

/* 插入排序 */
func Insertsort(s []int) {
	for i := 1; i < len(s); i++ {
		sk := i
		for j := i; j > 0; j-- {

			if s[sk] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
				sk = j - 1
			}
		}
	}
}

/* 选择排序 */
func Selectsort(s []int) {
	for i := len(s); i > 0; i-- {
		sk := 0
		for j := 0; j < i; j++ {
			if s[sk] < s[j] {
				sk = j
			}
		}
		s[sk], s[i-1] = s[i-1], s[sk]
	}
}

/******* nlogn ***************/

func Quicksort(s []int) {
	if len(s) < 1 {
		return
	}

	leftSk := 0
	sk := len(s) - 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] < s[sk] {
			s[i], s[leftSk] = s[leftSk], s[i]
			leftSk++
		}
	}
	s[sk], s[leftSk] = s[leftSk], s[sk]
	Quicksort(s[0:leftSk])
	Quicksort(s[leftSk:])
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0 // left和right的index位置
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
			continue
		}
		result = append(result, left[m])
		m++
	}
	result = append(result, right[n:]...) // 这里竟然没有报数组越界的异常？
	result = append(result, left[m:]...)
	return result
}

/* 归并排序 */
func Mergesort(s []int) {
	if len(s) < 2 {
		return
	}
	part := len(s) / 2
	fmt.Println("pa", part)
	Mergesort(s[:part])
	Mergesort(s[part:])
	r := merge(s[:part], s[part:])
	fmt.Println("r:", r)
	copy(s, r)
}

/* 初始化堆 */
func stackinit(s []int) {
	for i := len(s) - 1; i > 0; i-- {
		var p int
		if i%2 != 0 {
			p = (i) / 2
		} else {
			p = (i)/2 - 1
		}
		if s[p] < s[i] {
			s[p], s[i] = s[i], s[p]
		}
	}
}

/* 堆排序 */
func Stacksort(s []int) {
	for i := len(s); i > 0; i-- {
		stackinit(s[:i])
		s[0], s[i-1] = s[i-1], s[0]
	}
}
