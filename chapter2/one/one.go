package one

import "container/list"

func RemoveDups(l *list.List) {
	m := map[int]int{}
	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.(int)
		m[v]++
	}
	toRemove := []*list.Element{}
	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.(int)
		if count := m[v]; count > 1 {
			toRemove = append(toRemove, e)
			m[v]--
		}
	}
	for _, e := range toRemove {
		l.Remove(e)
	}
}
