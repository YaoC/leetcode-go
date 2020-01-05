package _04_MedianOfTwoSortedArrays

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "right-1",
			args: args{
				nums1: []int{1, 2, 4, 5, 11, 22},
				nums2: []int{},
			},
			want: 4.5,
		},
		{
			name: "right-2",
			args: args{
				nums1: []int{1, 2, 4, 5, 22},
				nums2: []int{},
			},
			want: 4,
		},
		{
			name: "right-3",
			args: args{
				nums1: []int{},
				nums2: []int{3, 6, 7, 8, 15, 21, 25},
			},
			want: 8,
		},
		{
			name: "right-4",
			args: args{
				nums1: []int{},
				nums2: []int{6, 7, 8, 15, 21, 25},
			},
			want: 11.5,
		},
		{
			name: "right-5",
			args: args{
				nums1: []int{1, 2, 5, 9, 22},
				nums2: []int{1, 3, 4, 6, 21, 25},
			},
			want: 5,
		},
		{
			name: "right-6",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
