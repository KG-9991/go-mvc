package services

import (
	"github.com/go-mvc/mvc/utils"

	"github.com/go-mvc/mvc/domain"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
