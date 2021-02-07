package chapter1

// NOTE:
//  We need to check the character encoding first.

func isUnique(s string) bool {
	m := make(map[int32]struct{}, len(s))
	for _, v := range s {
		if _, exist := m[v]; exist {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}
