package authentication_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFtx(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authentication Suite")
}
