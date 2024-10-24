package leetcode_go

import (
	"strings"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	longestLength := 0
	substr := ""
	for _, v := range s {
		vv := string(v)
		idx := strings.Index(substr, vv)
		if idx == -1 {
			// not hit
		} else {
			// hit
			if len(substr) > longestLength {
				longestLength = len(substr)
			}

			if idx == len(substr)-1 {
				// last hit
				substr = ""
			} else {
				// middle hit
				substr = substr[idx+1:]
			}
		}
		substr += vv
	}

	if len(substr) > longestLength {
		longestLength = len(substr)
	}

	return longestLength
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "abcabcbb",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "bbbbb",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "pwwkew",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "dvdf",
			args: args{
				s: "dvdf",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
