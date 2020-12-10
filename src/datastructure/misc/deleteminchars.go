package misc

func DeleteMinChars(arr []int) int {
	total := 0
	previous := 0
	for i := 0; i < len(arr)-1; i++ {
		if previous == arr[i+1] {
			total += arr[i+1]
		} else if arr[i] == arr[i+1] && arr[i]-1 == previous {
			total += arr[i]
			previous = arr[i]
		} else if arr[i] == arr[i+1] {
			total += 1
			previous = arr[i]
		}
	}
	return total
}

func DeleteMinCharsMap(arr map[int]int) int {
	total := 0
	for k, v := range arr {
		if _, ok := arr[k-1]; ok && v > 1 {
			total += (v - 1) * k
		} else if v > 1 {
			total += 1 + (v-2)*k
		}
	}
	return total
}
