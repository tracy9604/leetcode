package recursion

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / MyPow(x, -n)
	}

	if n%2 == 1 {
		return x * MyPow(x, n/2) * MyPow(x, n/2)
	}

	return MyPow(x, n/2) * MyPow(x, n/2)
}
