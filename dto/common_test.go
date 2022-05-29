package dto

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DTOCommonTestSuite struct {
	suite.Suite
}

func (suite *DTOCommonTestSuite) TestBindIdUriParameterDTO() {
	m := make(map[string][]string)
	m["id"] = []string{"123"}
	var dto IdUriParameterDTO
	b := binding.Uri
	assert.NoError(suite.T(), b.BindUri(m, &dto))
	assert.Equal(suite.T(), uint64(123), dto.ID)
}

func (suite *DTOCommonTestSuite) TestBindHashIdUriParameterDTO() {
	m := make(map[string][]string)
	m["id"] = []string{"foo"}
	var dto HashIdUriParameterDTO
	b := binding.Uri
	assert.NoError(suite.T(), b.BindUri(m, &dto))
	assert.Equal(suite.T(), "foo", dto.ID)
}

func TestDTOCommonTestSuite(t *testing.T) {
	suite.Run(t, new(DTOCommonTestSuite))
}
