package chapter1

import "testing"

func Test_escapeSpace(t *testing.T) {
	type args struct {
		s string
		l int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: "Mr John Smith ",
				l: 13,
			},
			want: "Mr%20John%20Smith",
		},
		{
			name: "case 2",
			args: args{
				s: " If there is a leading space",
				l: 28,
			},
			want: "%20If%20there%20is%20a%20leading%20space",
		},
		{
			name: "case 3",
			args: args{
				s: "The string will be trimmed and escaped",
				l: 16,
			},
			want: "The%20string%20will%20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := escapeSpace(tt.args.s, tt.args.l); got != tt.want {
				t.Errorf("escapeSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
