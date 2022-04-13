package main

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func IntSliceMax(s []int) (max int) {
	max = s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return
}

func IntSliceMin(s []int) (min int) {
	min = s[0]
	for _, v := range s {
		if v > min {
			min = v
		}
	}
	return
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
