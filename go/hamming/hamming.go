package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	var distance int

	if len(a) != len(b) {
		return 0, errors.New("")
	}

	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
