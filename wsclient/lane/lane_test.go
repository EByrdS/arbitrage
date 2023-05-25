package lane_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/arbitrage/wsclient/lane"
)

var _ = Describe("Lane", func() {
	Describe(".New", func() {
		var (
			verifier func([]byte) bool
			testLane lane.Lane
		)

		BeforeEach(func() {
			verifier = func(byteMessage []byte) bool {
				return true
			}
			testLane = lane.New(verifier)
		})

		It("creates a capacity 5 buffered channel", func() {
			Ω(testLane.C).Should(HaveCap(5))
		})
	})
})
