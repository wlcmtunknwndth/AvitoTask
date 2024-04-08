package auth

type Storage interface {
	GetUserToken()
	GetAdminToken()
	RegisterAdminToken()
	RegisterUserToken()
}
