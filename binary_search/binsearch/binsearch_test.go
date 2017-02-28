package binsearch

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
	{"In array", args{[]int{1,2,3,4,5}, 4}, 3, 2},
	{"First", args{[]int{1,2,3,4,5}, 1}, 0, 2},
	{"Middle", args{[]int{1,2,3,4,5}, 3}, 2, 1},
	{"Last", args{[]int{1,2,3,4,5}, 5}, 4, 3},
	{"Not in array", args{[]int{1,2,3,4,5}, 6}, -1, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BinarySearch(tt.args.arr, tt.args.target)
			if got != tt.want {
				t.Errorf("BinarySearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BinarySearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
