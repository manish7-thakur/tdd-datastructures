package fibonacci

type Fib struct {
}

func (f Fib) generate(n int) int {
	if n == 1 {
		return 1
	} else if n == 0 {
		return 0
	} else {
		return f.generate(n-1) + f.generate(n-2)
	}
}
