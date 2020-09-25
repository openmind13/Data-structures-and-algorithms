package sorting

// BubbleSort ...
// n^2
func BubbleSort(buffer []int) []int {
	size := len(buffer)
	for i := 0; i < size-1; i++ {
		for j := 0; j < (size - i - 1); j++ {
			if buffer[j] > buffer[j+1] {
				buffer[j], buffer[j+1] = buffer[j+1], buffer[j]
			}
		}
	}
	return buffer
}

// InsertionSort ...
func InsertionSort(buffer []int) []int {
	size := len(buffer)

	for i := 1; i < size; i++ {
		key := buffer[i]
		j := i
		for j > 0 && key < buffer[j-1] {
			buffer[j] = buffer[j-1]
			j--
		}
		buffer[j] = key
	}

	return buffer
}

// SelectionSort ...
func SelectionSort(buffer []int) []int {
	size := len(buffer)

	for i := 0; i < size; i++ {
		min := i
		for j := i + 1; j < size; j++ {
			if buffer[j] < buffer[min] {
				min = j
			}
		}

		if min != i {
			buffer[i], buffer[min] = buffer[min], buffer[i]
		}
	}

	return buffer
}
