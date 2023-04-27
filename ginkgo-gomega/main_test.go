package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GinkgoGomega", func() {
	Describe("Test EvenorOdd Function", func() {
		Context("When Number is Even", func() {
			It("Check 4 is even or odd", func ()  {
				got := EvenOrOdd(4)
				Expect(got).To(Equal("EVEN"))
			})
		})
		Context("When Number is Odd", func() {
			It("Check 7 is even or odd", func ()  {
				got := EvenOrOdd(7)
				Expect(got).To(Equal("ODD"))
			})
		})
		Context("Error", func() {
			It("Check 7 is even or odd", func ()  {
				got := EvenOrOdd(7)
				Expect(got).To(Equal("ODD"))
			})
		})
	})
})


func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_hello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hello(); got != tt.want {
				t.Errorf("hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
