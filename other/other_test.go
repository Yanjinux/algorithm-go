package other

import "testing"

func TestLongestPalindrome(t *testing.T) {
	exam := []struct {
		Val string
		Res string
	}{
		{"cbbd", "bb"}, {"babad", "bab"},
	}
	for i, _ := range exam {
		res := longestPalindrome(exam[i].Val)
		if res != exam[i].Res {
			t.Errorf("val string %v expect %v but return %v", exam[i].Val, exam[i].Res, res)
		}
	}
}
