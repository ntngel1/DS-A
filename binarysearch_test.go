package dsa

import (
	"cmp"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	type args[K cmp.Ordered] struct {
		items []K
		value K
	}
	type testCase[K cmp.Ordered] struct {
		name string
		args args[K]
		want int
	}

	tests := []testCase[string]{
		{
			args: args[string]{
				items: []string{"Casey Neistat", "Homer Simpson", "Captain America", "Gordon Freeman"},
				value: "Gordon Freeman",
			},
			want: 3,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat", "Homer Simpson", "Captain America"},
				value: "Homer Simpson",
			},
			want: 1,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat", "Homer Simpson", "Captain America"},
				value: "Casey Neistat",
			},
			want: 0,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat", "Homer Simpson", "Captain America"},
				value: "Nobody Knows This Guy",
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.items, tt.args.value); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
