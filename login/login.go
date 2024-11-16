package login

import (
	"context"
	"fmt"
	"test-bot/constants"
	swagger "test-bot/go_api_client"
)

func NewLogin() {
	cfg := swagger.NewConfiguration()
	fmt.Println(cfg)
	upstoxClient := swagger.NewAPIClient(cfg)
	opts := swagger.LoginApiAuthorizeOpts{}
	res, err := upstoxClient.LoginApi.Authorize(context.TODO(), constants.UPSTOX_CLIENT_ID, constants.UPSTOX_REDIRECT_URI, &opts)
	fmt.Println(err)
	fmt.Println(res.Status)

}
