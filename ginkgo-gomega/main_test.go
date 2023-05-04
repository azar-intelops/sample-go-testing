package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GinkgoGomega", func() {
	// EvenorOdd
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
		// Makes it pending!
		XContext("Error", func() {
			It("Check 0 is even or odd", func ()  {
				got := EvenOrOdd(0)
				Expect(got).To(Equal("Invalid"))
			})
		})
	})

	// Hello Function
	Describe("Test Hello Function", func() {
		Context("Just Says Hello World", func() {
			It("Check 4 is even or odd", func ()  {
				got := hello()
				Expect(got).To(Equal("Hello World"))
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
