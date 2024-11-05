package sort

import "cmp"

func bubbleSort[T cmp.Ordered](items []T) {
	if len(items) <= 1 {
		return
	}

	for a := 0; a < len(items)-1; a++ {
		for b := 0; b < (len(items) - 1 - a); b++ {
			if items[b] > items[b+1] {
				temp := items[b]
				items[b] = items[b+1]
				items[b+1] = temp
			}
		}
	}
}
