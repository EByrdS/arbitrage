package authentication_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/arbitrage/apis/binance/authentication"
)

var _ = Describe("Binance Authentication", func() {
	Describe(".Sign", func() {
		It("replicates example 1", func() {
			// https://binance-docs.github.io/apidocs/spot/en/#signed-trade-user_data-and-margin-endpoint-security
			signature := authentication.Sign(
				"symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC&"+
					"quantity=1&price=0.1&recvWindow=5000",
				1499827319559,
				"NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j",
			)

			Ω(signature).Should(Equal(
				"c8db56825ae71d6d79447849e617115f4a920fa2acdcab2b053c4b2838bd6b71",
			))
		})

		It("replicates example 2", func() {
			signature := authentication.Sign(
				"symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC"+ // no &
					"quantity=1&price=0.1&recvWindow=5000",
				1499827319559,
				"NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j",
			)

			Ω(signature).Should(Equal(
				"0fd168b8ddb4876a0358a8d14d0c9f3da0e9b20c5d52b2a00fcf7d1c602f9a77",
			))
		})
	})
})
