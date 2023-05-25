package authentication_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/arbitrage/apis/ftx/authentication"
)

var _ = Describe("FTX Authentication", func() {
	Describe(".Sign", func() {
		It("replicates GET signature from docs", func() {
			// https://blog.ftx.com/blog/api-authentication/
			signature := authentication.Sign(
				"GET",
				"/api/markets",
				"",
				1588591511721,
				"T4lPid48QtjNxjLUFOcUZghD7CUJ7sTVsfuvQZF2",
			)

			Ω(signature).Should(Equal(
				"dbc62ec300b2624c580611858d94f2332ac636bb86eccfa1167a7777c496ee6f",
			))
		})

		It("replicates POST signature from docs", func() {
			// https://blog.ftx.com/blog/api-authentication/
			signature := authentication.Sign(
				"POST",
				"/api/orders",
				`{"market": "BTC-PERP", "side": "buy", "price": 8500, "size": 1, "type": "limit", "reduceOnly": false, "ioc": false, "postOnly": false, "clientId": null}`,
				1588591856950,
				"T4lPid48QtjNxjLUFOcUZghD7CUJ7sTVsfuvQZF2",
			)

			Ω(signature).Should(Equal(
				"c4fbabaf178658a59d7bbf57678d44c369382f3da29138f04cd46d3d582ba4ba",
			))
		})

		It("panics with Unrecognized method", func() {
			Ω(func() {
				authentication.Sign("UPLOAD", "/api/orders", "", 1588591856950, "T4lPid48QtjNxjLUFOcUZghD7CUJ7sTVsfuvQZF2")
			}).Should(PanicWith("Unrecognized method UPLOAD"))
		})
	})
})
