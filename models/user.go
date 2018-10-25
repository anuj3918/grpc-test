package grpctest

type UserService struct{}

// UserMethods allowed on these models
type UserMethods interface {
	Set(userID string, properties interface{}) (e *error)
}
