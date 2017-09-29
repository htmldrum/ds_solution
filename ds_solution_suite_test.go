package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDsSolution(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DsSolution Suite")
}
