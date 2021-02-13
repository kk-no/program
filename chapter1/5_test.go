package chapter1

import "testing"

func Test_isSimilar(t *testing.T) {
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
				s1: "aaa",
				s2: "aaa",
			},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				s1: "askahjg",
				s2: "askadhjg",
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				s1: "aaabbbccc",
				s2: "abc",
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				s1: "abc",
				s2: "cb",
			},
			want: true,
		},
		{
			name: "case 5",
			args: args{
				s1: "abcdefghijk",
				s2: "ecghijkdfabsfkjhgsv",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSimilar(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
