package selsort

import (
	"reflect"
	"testing"
)

func Test_indexOfMinimal(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	{"Minimal at first index", args{[]int{1,5,6}}, 0},
	{"Minimal at last index", args{[]int{7,9,1}}, 2},
	{"Multiple minimal elements", args{[]int{7,4,4}}, 1},
	{"All elements identical", args{[]int{5,5,5}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indexOfMinimal(tt.args.arr); got != tt.want {
				t.Errorf("indexOfMinimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swap(t *testing.T) {
	type args struct {
		arr    []int
		index1 int
		index2 int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
	{"Swap middle elements", args{[]int{1, 2, 3, 4}, 1, 2}, []int{1, 3, 2, 4}},
	{"Swap first and last elements", args{[]int{1,2,3,4}, 0, 3}, []int{4,2,3,1}},
	{"Swap on same element", args{[]int{1, 2, 3, 4}, 1, 1}, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swap(tt.args.arr, tt.args.index1, tt.args.index2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
	{"Sort normal", args{[]int{6,5,4,6,2,1}}, []int{1,2,4,5,6,6}},
	{"Sort single element", args{[]int{2}}, []int{2}},
	{"Sort empty slice", args{[]int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
