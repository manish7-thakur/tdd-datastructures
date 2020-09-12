package fibonacci

type Fib struct {
	mem []int
}

func New(capacity int) Fib {
	return Fib{make([]int, capacity)}
}
func (f Fib) generateRec(n int) int {
	if n <= 1 {
		return n
	} else {
		f1 := f.mem[n-1]
		f2 := f.mem[n-2]
		if f1 == 0 {
			f1 = f.generateRec(n - 1)
		}
		if f2 == 0 {
			f2 = f.generateRec(n - 2)
		}
		f.mem[n-1] = f1
		f.mem[n-2] = f2
		return f1 + f2
	}
}
