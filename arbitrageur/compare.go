package arbitrageur

import (
	"time"

	"github.com/arbitrage/ticker"
)

type CompareFunc func(left ticker.Ticker, right ticker.Ticker)

func Compare(
	leftChan chan ticker.Ticker,
	rightChan chan ticker.Ticker,
	action CompareFunc,
	interrupter chan int,
) {

	var (
		leftTicker, rightTicker ticker.Ticker
	)

	/*
		actionSparser.
		This channel is empty when the action can be made, and has one message
		when an action is being done.
		As the action is done concurrently we need to avoid executing
		two of the same actions overlapping one another.
	*/
	actionSparser := make(chan int, 1)
	// Ensure to close the actionSparser channel
	defer func() { <-time.After(time.Second); close(actionSparser) }()

	// Anonymous wrapper function.
	// It ensures to free the sparser channel after completion.
	actionNoOverlap := func(left ticker.Ticker, right ticker.Ticker) {
		defer func() { <-actionSparser }()

		action(left, right)
	}

	for {
		select {
		case <-interrupter:
			return
		case leftTicker = <-leftChan:
			/*
						The oncoming message will at least be read even if the action is not
				  	executed, ensuring that the action will be executed with recent data
						next time.

						This also avoids congestion in the channel if messages come
						faster than they leave.

						The logic is different from a MutEx, as a MutEx executes 'when'
						the resource is unlocked, which would lead to blocking this caller.
						Here the action is executed 'only if' the resource is unlocked, and
						passes otherwise.
			*/
			if len(actionSparser) == 0 {
				actionSparser <- 1
				go actionNoOverlap(leftTicker, rightTicker)
			}
		case rightTicker = <-rightChan:
			if len(actionSparser) == 0 {
				actionSparser <- 1
				go actionNoOverlap(leftTicker, rightTicker)
			}
		}
	}
}
