package power

type Pow struct {
	mem []int
}

func New(capacity int) Pow {
	return Pow{mem: make([]int, capacity)}
}

func (p Pow) calc(a int, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return a
	}
	res := p.mem[n]
	if res == 0 {
		nearest := nearestPow2(n)
		rest := n - nearest
		res = p.calc(a, nearest/2) * p.calc(a, nearest/2) * p.calc(a, rest)
	}
	p.mem[n] = res
	return res

}

/*func pow(a int, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return a
	} else {
		res := mem[n]
		if res == 0 {
			nearest := nearestPow2(n)
			rest := n - nearest
			res = pow(a, nearest/2) * pow(a, nearest/2) * pow(a, rest)
		}
		mem[n] = res
		return res
	}
}*/

func nearestPow2(n int) int {
	nearest := 2
	for ; nearest*2 <= n; {
		nearest *= 2
	}
	return nearest
}
