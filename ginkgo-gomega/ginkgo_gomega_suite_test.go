package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGinkgoGomega(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoGomega Suite")
}
