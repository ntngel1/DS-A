package dsa

import "cmp"

func binarySearch[K cmp.Ordered](items []K, value K) int {
	left := 0
	right := len(items) - 1

	for left <= right {
		middle := left + (right-left)/2

		if items[middle] == value {
			return middle
		}

		if items[middle] < value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}
