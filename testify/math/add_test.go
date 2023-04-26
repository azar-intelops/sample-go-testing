package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name:"Add 2, 4",
			args:args{
				a:2,
				b:4,
			},
			want:6,
		},
		{
			name:"Add 2, 8",
			args:args{
				a:2,
				b:8,
			},
			want:6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Add(tt.args.a, tt.args.b) )
			// if got := Add(tt.args.a, tt.args.b); got != tt.want {
			// 	t.Errorf("Add() = %v, want %v", got, tt.want)
			// }
		})
	}
}
