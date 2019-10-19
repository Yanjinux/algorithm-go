package stack

type List struct {
	_val   []interface{}
	_count int // 数据数目
}

func (l *List) Push(val interface{}) {
	if l == nil {
		return
	}
	l._val = append(l._val, val)
	l._count++
}

func (l *List) Pop() interface{} {
	if l == nil || l._count == 0 {
		return nil
	}
	l._count--
	t := l._val[l._count]
	l._val = l._val[:l._count]
	return t
}

func (l *List) GetSize() int {
	return l._count
}

/* 20 有效括号 */
func isValid(s string) bool {
	var l List
	for _, c := range s {
		switch c {
		case '(', '{', '[':
			l.Push(c)
		case ')':
			if l.Pop() != '(' {
				return false
			}
		case '}':
			if l.Pop() != '{' {
				return false
			}
		case ']':
			if l.Pop() != '[' {
				return false
			}
		}
	}
	if l.GetSize() == 0 {
		return true
	}
	return false
}
