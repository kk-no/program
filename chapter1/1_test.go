package chapter1

import "testing"

func Test_isUnique(t *testing.T) {
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
				s: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				s: "aba",
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				s: "abcdefghijklmnopqrstuvwxyzaBCDEFGHIJKLMNOPQRSTUVWXYZ",
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				s: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnique(tt.args.s); got != tt.want {
				t.Errorf("isUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
