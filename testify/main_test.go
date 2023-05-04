package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "Main Test Case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestEvenOrOdd(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"Check 4 is even or odd",
			args:args{
				num:4,
			},
			want:"EVEN",
		},
		{
			name:"Check 7 is even or odd",
			args:args{
				num:7,
			},
			want:"ODD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, EvenOrOdd(tt.args.num) )
			if got := EvenOrOdd(tt.args.num); got != tt.want {
				t.Errorf("EvenOrOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name:"Test Hello Function",
			want:"Hello World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hello(); got != tt.want {
				t.Errorf("hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
