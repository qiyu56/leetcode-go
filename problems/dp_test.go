package problems

import "testing"

func Test_knapsack(t *testing.T) {
	type args struct {
		v  int
		n  int
		vw [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{v: 10, n: 2, vw: [][]int{{1, 3}, {10, 4}}}, want: 4},
		{name: "2", args: args{v: 10, n: 2, vw: [][]int{{1, 3}, {9, 8}}}, want: 11},
		{name: "3", args: args{v: 10, n: 10, vw: [][]int{{2, 0}, {2, 3}, {7, 5}, {9, 2}, {2, 8}, {9, 7}, {3, 6}, {1, 2}, {9, 3}, {1, 9}}}, want: 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapsack(tt.args.v, tt.args.n, tt.args.vw); got != tt.want {
				t.Errorf("knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBackPack(t *testing.T) {
	type args struct {
		m int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{m: 10, a: []int{3, 4, 8, 5}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BackPack(tt.args.m, tt.args.a); got != tt.want {
				t.Errorf("BackPack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxEnvelopes(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "1", args: args{envelopes: [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}}, want: 3},
		{name: "2", args: args{envelopes: [][]int{{1, 1}, {1, 1}, {1, 1}}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}
