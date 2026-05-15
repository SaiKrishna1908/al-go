package search

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr)

	for low < high {
		mid := low + (high-low)/2

		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] == target {
			return mid
		}
	}

	return low
}
