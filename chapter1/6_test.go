package chapter1

import "testing"

func Test_compressStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				s: "aabcccccaaa",
			},
			want: "a2b1c5a3",
		},
		{
			name: "case 2",
			args: args{
				s: "AAbccCCcaAa",
			},
			want: "AAbccCCcaAa",
		},
		{
			name: "case 3",
			args: args{
				s: "AAbCCCCcAAA",
			},
			want: "A2b1C4c1A3",
		},
		{
			name: "case 4",
			args: args{
				s: "abc",
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressStr(tt.args.s); got != tt.want {
				t.Errorf("compressStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
