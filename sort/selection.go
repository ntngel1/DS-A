package sort

import "cmp"

func selectionSort[T cmp.Ordered](items []T) {
	if len(items) <= 1 {
		return
	}

	for a := 0; a < len(items)-1; a++ {
		minIdx := a
		minValue := items[a]
		for b := a + 1; b < len(items); b++ {
			if items[b] <= minValue {
				minValue = items[b]
				minIdx = b
			}
		}

		if items[a] != minValue {
			temp := items[a]
			items[a] = minValue
			items[minIdx] = temp
		}
	}
}
