package main

import "testing"

func TestSubstr(t *testing.T) {
	test := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbb", 1},
		{"abcabcd", 4},

		// Chinese support
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发胡发黑会飞花", 6},
	}

	for _, tt := range test {
		if actual := noneRepeating(tt.s); actual != tt.ans {
			t.Errorf("noneRepeating(%s); got %d; expected %d",
				tt.s, actual, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发胡发黑会飞花"
	ans := 6

	for i := 0; i < b.N; i++ {
		if actual := noneRepeating(s); actual != ans {
			b.Errorf("noneRepeating(%s); got %d; expected %d",
				s, actual, ans)
		}

	}
}
