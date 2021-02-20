package chapter1

import (
	"reflect"
	"testing"
)

func Test_rotateImage(t *testing.T) {
	type args struct {
		image [][]byte
	}
	type wants struct {
		image [][]byte
	}
	tests := []struct {
		name    string
		args    args
		wants   wants
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				image: [][]byte{},
			},
			wantErr: true,
		},
		{
			name: "case 2",
			args: args{
				image: [][]byte{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
				},
			},
			wantErr: true,
		},
		{
			name: "case 3",
			args: args{
				image: [][]byte{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
					{13, 14, 15, 16},
				},
			},
			wants: wants{
				image: [][]byte{
					{13, 9, 5, 1},
					{14, 10, 6, 2},
					{15, 11, 7, 3},
					{16, 12, 8, 4},
				},
			},
			wantErr: false,
		},
		{
			name: "case 4",
			args: args{
				image: [][]byte{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 10},
					{11, 12, 13, 14, 15},
					{16, 17, 18, 19, 20},
					{21, 22, 23, 24, 25},
				},
			},
			wants: wants{
				image: [][]byte{
					{21, 16, 11, 6, 1},
					{22, 17, 12, 7, 2},
					{23, 18, 13, 8, 3},
					{24, 19, 14, 9, 4},
					{25, 20, 15, 10, 5},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := rotateImage(tt.args.image); err != nil {
				if !tt.wantErr {
					t.Errorf("rotateImage() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(tt.args.image, tt.wants.image) {
				t.Errorf("rotateImage() unmatch got = %v, want %v", tt.args.image, tt.wants.image)
			}
		})
	}
}
