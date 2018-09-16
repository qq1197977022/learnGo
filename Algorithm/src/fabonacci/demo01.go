package fabonacci

func Recursion(n int) (ret int) { //递归
	if n == 0 || n == 1 {
		ret = 1
	} else {
		ret = Recursion(n-1) + Recursion(n-2)
	}

	return
}

func NoRecusion(n int) int {
	x, y := 1, 1

	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
