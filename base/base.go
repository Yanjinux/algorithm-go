package base

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
