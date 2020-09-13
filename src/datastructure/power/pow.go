package power

func pow(a int, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return a
	} else {
		return pow(a, n/2) * pow(a, n/2)
	}
}

func nearestPow2(n int) int {
	nearest := 2
	for ; nearest*2 <= n; {
		nearest *= 2
	}
	return nearest
}
