package login

import (
	"context"
	"fmt"
	"test-bot/constants"
	swagger "test-bot/go_api_client"
	"test-bot/helpers"

	"github.com/antihax/optional"
)

func (l *loginRepositoryImpl) NewLogin() {
	url := "https://api.upstox.com/v2/login/authorization/dialog?" + "client_id=" + l.conf.UpstoxClientId + "&redirect_uri=" + l.conf.UpstoxRedirectUri + "&state=code"
	err := helpers.OpenBrowser(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Please enter code:")
	var code string
	fmt.Scan(&code)

	cfg := swagger.NewConfiguration()
	upstoxClient := swagger.NewAPIClient(cfg)
	opts := swagger.LoginApiTokenOpts{
		Code:         optional.NewString(code),
		ClientId:     optional.NewString(constants.UPSTOX_CLIENT_ID),
		ClientSecret: optional.NewString(constants.UPSTOX_CLIENT_SECRET),
		RedirectUri:  optional.NewString(constants.UPSTOX_REDIRECT_URI),
		GrantType:    optional.NewString("authorization_code"),
	}
	token, _, err := upstoxClient.LoginApi.Token(context.TODO(), &opts)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token.AccessToken)
	l.Token.SetToken(token.AccessToken)
}
