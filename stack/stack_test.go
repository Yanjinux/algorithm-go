package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	var l List
	exam := []string{"cbbd", "babad", "bab"}
	mc := len(exam)
	for i := 0; i < mc; i++ {
		l.Push(exam[i])
	}
	if l.GetSize() != mc {
		t.Errorf("Size error: expect %d but ret %d.", mc, l.GetSize())
		return
	}
	for i := mc - 1; i >= 0; i-- {
		pv := l.Pop()
		if pv != exam[i] {
			t.Errorf("Val error: expect %s but ret %s.", exam[i], pv)
			return
		}
	}
	if l.GetSize() != 0 {
		t.Errorf("Size error: expect %d but ret %d.", 0, l.GetSize())
		return
	}
}

func TestisValid(t *testing.T) {
	exam := []struct {
		val string
		res bool
	}{{"()", true}, {"([{}])", true}, {"()[]{}", true}, {"([", false}}

	for _, v := range exam {
		r := isValid(v.val)
		if r != v.res {
			t.Errorf("error: %s", v.val)
			return
		}
	}
}
