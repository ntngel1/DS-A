package sort

import (
	"cmp"
	"fmt"
	"slices"
	"testing"
)

func Test_sort(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		args []T
		want []T
	}
	tests := []testCase[int]{
		{
			args: []int{},
			want: []int{},
		},
		{
			args: []int{1},
			want: []int{1},
		},
		{
			args: []int{20, 1},
			want: []int{1, 20},
		},
		{
			args: []int{3, 2, 3},
			want: []int{2, 3, 3},
		},
		{
			args: []int{1, 5, 9, 10},
			want: []int{1, 5, 9, 10},
		},
		{
			args: []int{5, 1, 2, 4, 1, 2, 9, 20, 10, 3, 3},
			want: []int{1, 1, 2, 2, 3, 3, 4, 5, 9, 10, 20},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.args), func(t *testing.T) {
			got := slices.Clone(tt.args)
			bubbleSort(got)
			if compareSlices(got, tt.want) != true {
				t.Error("Bubble Sort: got", got, ", want", tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.args), func(t *testing.T) {
			got := slices.Clone(tt.args)
			selectionSort(got)
			if compareSlices(got, tt.want) != true {
				t.Error("Selection Sort: got", got, ", want", tt.want)
			}
		})
	}
}

func compareSlices[T cmp.Ordered](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
