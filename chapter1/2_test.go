package chapter1

import "testing"

func Test_isPermutation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				s1: "dog",
				s2: "god",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				s1: "test case",
				s2: "case test",
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				s1: "Test case",
				s2: "Case test",
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				s1: "test case different string length",
				s2: "test case",
			},
			want: false,
		},
		{
			name: "case 5",
			args: args{
				s1: "",
				s2: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPermutation(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
