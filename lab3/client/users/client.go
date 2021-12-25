package users

import (
	"github.com/zavad4/go/tree/main/lab3/client/common"
	"github.com/zavad4/go/tree/main/lab3/client/dto"
)

func RegistrateUser(client *common.Client, userInfo *dto.User) (int, error) {
	return client.Post("/users", userInfo)
}
