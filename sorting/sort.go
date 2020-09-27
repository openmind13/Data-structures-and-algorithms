package sorting

// ReverseArray ...
func ReverseArray(buffer []int) []int {
	newBuffer := make([]int, len(buffer))
	for i := range buffer {
		newBuffer[len(newBuffer)-1-i] = buffer[i]
	}

	return newBuffer
}

// BubbleSort ...
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

// QuickSort ...
func QuickSort(buffer []int) []int {
	if len(buffer) < 2 {
		return buffer
	}

	left := 0
	right := len(buffer) - 1
	pivot := left

	buffer[pivot], buffer[right] = buffer[right], buffer[pivot]

	for i := range buffer {
		if buffer[i] < buffer[right] {
			buffer[i], buffer[left] = buffer[left], buffer[i]
			left++
		}
	}

	buffer[left], buffer[right] = buffer[right], buffer[left]

	QuickSort(buffer[:left])
	QuickSort(buffer[left+1:])

	return buffer
}
