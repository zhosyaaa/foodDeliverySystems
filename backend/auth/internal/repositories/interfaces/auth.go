package interfaces

type AuthRepository interface {
	RegisterUser(username string, password string) error
	AuthenticateUser(username string, password string) (bool, error)
}
