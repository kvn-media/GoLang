package main
// Binary Search Algorithm - Iterative Way
func BinarySearch(data []int, value int) bool {
	size := len(data)
	low := 0
	high := size - 1
	mid := 0

	for low <= high {
		mid = low + (high + low) / 2
		if data[mid] == value {
			return true
		} else if data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

// Binary Search Algorithm - Recursive Way
func BinarySearchRecursive(data []int, low int, high int, value int) bool {
	if low > high {
		return false
	}
	mid := low + (high - low) / 2
	if data[mid] == value {
		return true
	} else if data[mid] < value {
		return BinarySearchRecursive(data, mid + 1, high, value)
	} else {
		return BinarySearchRecursive(data, low, mid - 1, value)
	}
}

func main() {
	
}