package lane_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLane(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lane Suite")
}
