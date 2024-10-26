package leetcode_go

import (
	"slices"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	numsAll := slices.Concat(nums1, nums2)
	slices.Sort(numsAll)
	if isEven := len(numsAll)%2 == 0; isEven {
		nearBeforeIdx := (len(numsAll) / 2) - 1
		nearAfterIdx := nearBeforeIdx + 1
		return float64(numsAll[nearBeforeIdx]+numsAll[nearAfterIdx]) / 2
	} else {
		medIdx := (len(numsAll) - 1) / 2
		return float64(numsAll[medIdx])
	}
}

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
			name: "nums1 = [1,3], nums2 = [2]",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2.0,
		},
		{
			name: "nums1 = [1,2], nums2 = [3,4]",
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
