package chapter1

const space = 32

func escapeSpace(s string, l int) string {
	if l == 0 {
		return ""
	}

	var (
		r     = ""
		sub   = s[:l]
		start = 0
		end   = l
	)

	for i, v := range sub {
		if v == space {
			r += sub[start:i] + "%20"
			start = i + 1
		}
	}
	r += sub[start:end]

	// return strings.Replace(s[:l], " ", "%20", -1)
	return r
}

// NOTE:
//  The implementation uses the string as is,
//  but Go strings are immutable, using slice will reduce the number of assignments.
