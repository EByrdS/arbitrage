package authentication

import "fmt"

func SecurityLevel(name string) (sendKey bool, sign bool) {
	sendKey = false
	sign = false

	if name == "NONE" {
		return
	}

	sendKey = true

	if name == "TRADE" || name == "MARGIN" || name == "USER_DATA" {
		sign = true
	} else if name != "USER_STREAM" && name != "MARKET_DATA" {
		panic(fmt.Sprintf("Unrecognized security level name: %s", name))
	}

	return
}
