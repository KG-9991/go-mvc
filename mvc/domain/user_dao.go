package domain

import (
	"fmt"
	"net/http"

	"github.com/go-mvc/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "K", LastName: "G", Email: "myemail.gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
