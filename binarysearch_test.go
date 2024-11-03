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
				items: []string{},
				value: "Somebody",
			},
			want: -1,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat"},
				value: "Casey Neistat",
			},
			want: 0,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat", "Homer Simpson", "Zed"},
				value: "Casey Neistat",
			},
			want: 0,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat", "Homer Simpson", "Zed"},
				value: "Homer Simpson",
			},
			want: 1,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat", "Homer Simpson", "Zed"},
				value: "Zed",
			},
			want: 2,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat", "Homer Simpson", "Zed"},
				value: "Captain America",
			},
			want: -1,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat", "Homer Simpson", "ZZZZZZZZZZED", "Zed"},
				value: "ZZZZZZZZZZED",
			},
			want: 2,
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

func Test_recursiveBinarySearch(t *testing.T) {
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
				items: []string{},
				value: "Somebody",
			},
			want: -1,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat"},
				value: "Casey Neistat",
			},
			want: 0,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat", "Homer Simpson", "Zed"},
				value: "Casey Neistat",
			},
			want: 0,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat", "Homer Simpson", "Zed"},
				value: "Homer Simpson",
			},
			want: 1,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat", "Homer Simpson", "Zed"},
				value: "Zed",
			},
			want: 2,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat", "Homer Simpson", "Zed"},
				value: "Captain America",
			},
			want: -1,
		},
		{
			args: args[string]{
				items: []string{"Casey Neistat", "Homer Simpson", "ZZZZZZZZZZED", "Zed"},
				value: "ZZZZZZZZZZED",
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursiveBinarySearch(tt.args.items, tt.args.value, 0, len(tt.args.items)-1); got != tt.want {
				t.Errorf("recursiveBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
