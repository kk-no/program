package chapter1

import "strconv"

// It is more memory efficient to use []string.

func compressStr(s string) string {
	var (
		current int32
		count   int

		r string
	)

	// init
	current = int32(s[0])
	count = 1

	for _, v := range s[1:] {
		if v != current {
			// push
			r += string(current) + strconv.Itoa(count)

			// next
			current = v
			count = 1
			continue
		}
		count++
	}

	// last push
	r += string(current) + strconv.Itoa(count)

	if len(r) < len(s) {
		return r
	}
	return s
}

// It works as expected, but not very fast.
// We need to look into StringBuilder in Go.
