package binance

// import (
// 	"fmt"
// 	"net/url"

// 	"github.com/arbitrage/apis/binance/rest"
// )

// type UserInfoResponse struct {
// 	UserId string `json:"userId"`
// 	Email  string `json:"string"`
// }

// func (api Binance) GetUserInfo() (userInfoResponse UserInfoResponse, err error) {
// 	resp, err := rest.Get("/user-info", url.Values{})
// 	if err != nil {
// 		return userInfoResponse, fmt.Errorf("error getting user info : %w", err)
// 	}
// 	defer resp.Body.Close()
// }
