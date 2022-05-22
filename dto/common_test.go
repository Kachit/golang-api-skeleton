package dto

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DTO_BindIdUriParameterDTO(t *testing.T) {
	m := make(map[string][]string)
	m["id"] = []string{"123"}
	var dto IdUriParameterDTO
	b := binding.Uri
	assert.NoError(t, b.BindUri(m, &dto))
	assert.Equal(t, uint64(123), dto.ID)
}

func Test_DTO_BindHashIdUriParameterDTO(t *testing.T) {
	m := make(map[string][]string)
	m["id"] = []string{"foo"}
	var dto HashIdUriParameterDTO
	b := binding.Uri
	assert.NoError(t, b.BindUri(m, &dto))
	assert.Equal(t, "foo", dto.ID)
}
