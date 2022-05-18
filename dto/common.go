package dto

import (
	"github.com/gin-gonic/gin"
)

type IdUriParameterDTO struct {
	ID uint64 `uri:"id"`
}

func BindIdUriParameterDTO(c *gin.Context) (*IdUriParameterDTO, error) {
	var uriDTO IdUriParameterDTO
	if err := c.ShouldBindJSON(&uriDTO); err != nil {
		return nil, err
	}
	return &uriDTO, nil
}
