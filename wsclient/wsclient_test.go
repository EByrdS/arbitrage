package wsclient_test

import (
	"fmt"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	client "github.com/arbitrage/wsclient"
	"github.com/arbitrage/wsclient/connection"
	"github.com/arbitrage/wsclient/lane"
)

func newMockConnector(succeed bool) mockConnector {
	return mockConnector{
		Succeed: succeed,
		bytesC:  make(chan []byte, 5),
		errorC:  make(chan error),
	}
}

type mockConnector struct {
	Succeed bool
	bytesC  chan []byte
	errorC  chan error
}

func errIfFalse(succeed bool) error {
	if succeed {
		return nil
	}
	return fmt.Errorf("error")
}

func (m mockConnector) Open(_ string) error {
	return errIfFalse(m.Succeed)
}
func (m mockConnector) Close(_ string) error {
	return errIfFalse(m.Succeed)
}
func (m mockConnector) SendJSON(_ interface{}) error {
	return errIfFalse(m.Succeed)
}
func (m mockConnector) SetPing(_ []byte, _ time.Duration) {}
func (m mockConnector) Stop()                             {}
func (m mockConnector) BytesC() chan []byte               { return m.bytesC }
func (m mockConnector) ErrorC() chan error                { return m.errorC }

// Check interface is implemented
var _ connection.Connector = (newMockConnector)(true)

var _ = Describe("WSClient concurrency", func() {
	var testClient *client.WSClient

	var saysHello = func(byteMessage []byte) bool {
		return strings.Contains(string(byteMessage), "hello")
	}

	BeforeEach(func() {
		testClient = &client.WSClient{
			Connector: newMockConnector(true),
			Lanes:     map[string]lane.Lane{},
		}

		go client.AssignChannels(testClient)
	})

	Describe(".Unsubscribe", func() {
		BeforeEach(func() {
			testClient.Lanes["helloLane"] = lane.New(saysHello)
		})

		It("does not interrupt other channels", func() {
			testClient.Lanes["holaLane"] = lane.New(func(byteMessage []byte) bool {
				return strings.Contains(string(byteMessage), "hola")
			})
			testClient.Lanes["bonjourLane"] = lane.New(func(byteMessage []byte) bool {
				return strings.Contains(string(byteMessage), "bonjour")
			})
			var helloMsg = []byte("hello friend")
			var holaMsg = []byte("hola amigo")
			var bonjourMsg = []byte("bonjour mon ami")

			Ω(testClient.Lanes).Should(HaveLen(3))

			testClient.BytesC() <- helloMsg
			testClient.BytesC() <- holaMsg
			testClient.BytesC() <- bonjourMsg

			Eventually(testClient.Lanes["helloLane"].C).
				Should(Receive(Equal(helloMsg)))
			Eventually(testClient.Lanes["holaLane"].C).
				Should(Receive(Equal(holaMsg)))
			Eventually(testClient.Lanes["bonjourLane"].C).
				Should(Receive(Equal(bonjourMsg)))

			message := map[string]string{"key": "value"}
			_ = testClient.Unsubscribe(message, "bonjourLane")

			testClient.BytesC() <- helloMsg
			testClient.BytesC() <- holaMsg
			testClient.BytesC() <- bonjourMsg

			Eventually(testClient.Lanes["helloLane"].C).
				Should(Receive(Equal(helloMsg)))
			Eventually(testClient.Lanes["holaLane"].C).
				Should(Receive(Equal(holaMsg)))
			Eventually(testClient.ErrorC()).
				Should(Receive(Equal(
					fmt.Errorf("message without lane (out of 2): bonjour mon ami"),
				)))

			Ω(testClient.Lanes).Should(HaveLen(2))
		})
	})
})

var _ = Describe("WSClient", func() {
	var testClient *client.WSClient
	var saysHello = func(byteMessage []byte) bool {
		return strings.Contains(string(byteMessage), "hello")
	}

	BeforeEach(func() {
		testClient = client.New()
	})

	Describe(".New", func() {
		It("creates empty Lanes map", func() {
			Ω(testClient.Lanes).Should(BeEmpty())
		})
	})

	Describe("Message distribution", func() {

		var acceptedMsg = []byte("hello friend")
		var declinedMsg = []byte("good bye friend")

		BeforeEach(func() {
			testClient.Lanes["helloLane"] = lane.New(saysHello)
		})

		It("sends message to right channel", func() {
			testClient.BytesC() <- acceptedMsg

			Consistently(testClient.ErrorC()).ShouldNot(Receive())
			Eventually(testClient.Lanes["helloLane"].C).
				Should(Receive(Equal(acceptedMsg)))
		})

		It("sends error when no channel available", func() {
			testClient.BytesC() <- declinedMsg

			Consistently(testClient.Lanes["helloLane"].C).ShouldNot(Receive())
			Eventually(testClient.ErrorC()).
				Should(Receive(MatchError(
					fmt.Errorf("message without lane (out of 1): good bye friend"),
				)))
		})
	})

	Describe(".Subscribe", func() {
		BeforeEach(func() {
			testClient.Connector = mockConnector{
				Succeed: true,
			}
		})

		It("creates new lane", func() {
			Ω(testClient.Lanes).Should(BeEmpty())

			message := map[string]string{"key": "value"}
			laneName := "newLane"

			newLane, _ := testClient.Subscribe(message, laneName, saysHello)
			Ω(newLane).Should(BeAssignableToTypeOf(lane.Lane{}))
			Ω(testClient.Lanes).Should(HaveLen(1))
		})

		It("returns new lane", func() {
			message := map[string]string{"key": "value"}
			laneName := "newLane"

			newLane, _ := testClient.Subscribe(message, laneName, saysHello)
			Ω(newLane).Should(BeAssignableToTypeOf(lane.Lane{}))
			Expect(testClient.Lanes[laneName].C).To(BeIdenticalTo(newLane.C))
		})

		It("fails with duplicated lane name", func() {
			message := map[string]string{"key": "value"}
			secondMessage := map[string]string{"key": "other value"}
			laneName := "newLane"

			newLane, _ := testClient.Subscribe(message, laneName, saysHello)
			Ω(newLane).Should(BeAssignableToTypeOf(lane.Lane{}))
			Ω(testClient.Lanes).Should(HaveLen(1))

			gotLane, err := testClient.Subscribe(secondMessage, laneName, saysHello)

			Ω(testClient.Lanes).Should(HaveLen(1))
			Ω(gotLane.C).Should(BeNil())
			Ω(err).Should(MatchError(
				fmt.Errorf("lane already exists: %s", laneName),
			))
		})
	})

	Describe(".Unsubscribe", func() {
		BeforeEach(func() {
			testClient = &client.WSClient{
				Connector: mockConnector{
					Succeed: true,
				},
				Lanes: map[string]lane.Lane{},
			}

			go client.AssignChannels(testClient)
			testClient.Lanes["helloLane"] = lane.New(saysHello)
		})

		It("deletes existing lane", func() {
			Ω(testClient.Lanes).Should(HaveLen(1))

			message := map[string]string{"key": "value"}
			_ = testClient.Unsubscribe(message, "helloLane")
			// Race condition: checking lanes but deleting one

			Ω(testClient.Lanes).Should(BeEmpty())
		})
	})
})
