package sliceexample

func AppendInt(x []int, y int) []int {
	var z []int
	zLen := len(x) + 1
	if cap(x) >= zLen {
		z = x[:zLen]
	} else {
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		// same as
		// z = make([]int, zCap)[:zLen]
		z = make([]int, zLen, zCap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func AppendSlice(x []int, y ...int) []int {
	var z []int
	zLen := len(x) + len(y)
	if cap(x) >= zLen {
		z = x[:zLen]
	} else {
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zCap)[:zLen]
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
