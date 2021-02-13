package chapter1

// If something is not treated as a character, you need to check it beforehand.
// Need to make sure it's case-sensitive.

func isPalindromePermutation(s string) bool {
	ignore := " "

	m := make(map[string]int32, len(s))
	for _, v := range s {
		sv := string(v)
		_, ok := m[sv]
		if !ok {
			m[sv] = 1
			continue
		}
		m[sv]++
	}

	// Odd count
	var c int

	for k, v := range m {
		if k == ignore {
			continue
		}
		if v%2 == 0 {
			continue
		}
		if c++; c > 1 {
			return false
		}
	}
	return true
}
