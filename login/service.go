package login

type LoginRepository interface {
	NewLogin()
	GetToken() string
}
