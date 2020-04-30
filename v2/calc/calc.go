package calc

func Add(args ...int) (s int) {
	s = 0
	for _, v := range args {
		s += v
	}
	return
}

func Sub(a, b int) int {
	return a - b
}
