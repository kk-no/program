package chapter1

// We need to check the behavior when the given strings are identical.

func isSimilar(s1, s2 string) bool {
	if s1 == s2 {
		return false
	}

	var (
		boundary = 1
		diff     = 0
	)

	// bug fixed
	s, l := swapShortest(s1, s2)

	m := make(map[int32]int32, len(s1)+len(s2))
	for _, v := range s {
		_, ok := m[v]
		if !ok {
			m[v] = 1
			continue
		}
		m[v]++
	}

	for _, v := range l {
		_, ok := m[v]
		if !ok {
			diff++
			continue
		}

		m[v]--

		if m[v] < 0 {
			diff++
		}
	}

	return diff <= boundary
}

func swapShortest(s1, s2 string) (short, long string) {
	if len(s1) > len(s2) {
		return s2, s1
	}
	return s1, s2
}

// There is a slightly better implementation.
