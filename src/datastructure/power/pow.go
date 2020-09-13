package power

func pow(a int, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return a
	} else {
		nearest := nearestPow2(n)
		rest := n - nearest
		return pow(a, nearest/2) * pow(a, nearest/2) * pow(a, rest)
	}
}

func nearestPow2(n int) int {
	nearest := 2
	for ; nearest*2 <= n; {
		nearest *= 2
	}
	return nearest
}
