package chapter1

// NOTE:
//  For strings, we need to check how case and whitespace are handled.
//  In addition to the following, there are other ways to compare strings that have been sorted identically.

// I couldn't solve the problem in the best way.

func isPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[int32]int)
	for _, v := range s1 {
		if _, exist := m[v]; exist {
			m[v]++
			continue
		}
		m[v] = 1
	}
	for _, v := range s2 {
		i, exist := m[v]
		if !exist {
			return false
		}
		if i-1 < 0 {
			return false
		}
		m[v]--
	}
	return true
}
