package forums

import (
	"github.com/zavad4/go/tree/main/lab3/client/common"
	"github.com/zavad4/go/tree/main/lab3/client/dto"
)

func GetForums(client *common.Client) ([]dto.ForumsResponse, error) {
	return client.Get("/forums")
}
