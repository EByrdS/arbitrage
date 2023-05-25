package connection_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/arbitrage/wsclient/connection"
)

var _ = Describe("Connection", func() {
	var testConn *connection.ExchangeConnection

	Describe(".New", func() {
		It("has a cap 5 BytesC channel", func() {
			testConn = connection.New()
			Ω(testConn.BytesC()).Should(HaveCap(5))
		})
	})
})
