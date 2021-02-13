package chapter1

import "testing"

func Test_isPalindromePermutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				s: "taco cat",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				s: "cat cat",
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				s: "abcd abcf",
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				s: "a",
			},
			want: true,
		},
		// Depending on assumptions, the implementation and the outcome of this case will vary.
		{
			name: "case 5",
			args: args{
				s: "Tac Cat",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromePermutation(tt.args.s); got != tt.want {
				t.Errorf("isPalindromePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
