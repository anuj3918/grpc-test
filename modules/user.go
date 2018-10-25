package grpctest

import (
	"fmt"
)

// Set user
func Set(userID string, properties interface{}) (e *error) {
	fmt.Println("i am inside set function of user module")
	return nil
}
