package comb_sort

func CombSort(arr []int) []int {
	arrlen := len(arr)
	width := nextGap(arrlen)
	sorted := false
	for ; !sorted; {
		if width == 1 { // This is last iteration
			sorted = true
		}
		for j := 0; j < arrlen-width; j++ {
			next := j + width
			if arr[j] > arr[next] {
				arr[j], arr[next] = arr[next], arr[j]
			}
		}
		width = nextGap(width)
	}
	return arr
}

func nextGap(width int) int {
	gap := width * 10 / 13
	if gap == 0 {
		gap = 1
	}
	return gap
}
