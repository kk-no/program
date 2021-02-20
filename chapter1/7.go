package chapter1

import (
	"errors"
)

// I couldn't solve the problem...

func rotateImage(image [][]byte) error {
	if len(image) == 0 || len(image) != len(image[0]) {
		return errors.New("rotateImage error")
	}
	n := len(image)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := image[first][i]

			image[first][i] = image[last-offset][first]
			image[last-offset][first] = image[last][last-offset]
			image[last][last-offset] = image[i][last]
			image[i][last] = top
		}
	}
	return nil
}
