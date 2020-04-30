package geometry

import "errors"

// Volume calculate volume of cube
func Volume(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	}
	return 0, errors.New("Zero length edge is not allowed")
}
