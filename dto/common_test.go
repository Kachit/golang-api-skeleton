package dto

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DTO_IdParameterPathDTO_Bind(t *testing.T) {
	m := make(map[string][]string)
	m["id"] = []string{"123"}
	var dto IdParameterPathDTO
	b := binding.Uri
	assert.NoError(t, b.BindUri(m, &dto))
	assert.Equal(t, uint64(123), dto.ID)
}

func Test_DTO_FilterParameterQueryStringDTO_HasSearch(t *testing.T) {
	d := &FilterParameterQueryStringDTO{}
	assert.False(t, d.HasSearch())
	d.Search = "foo"
	assert.True(t, d.HasSearch())
}

func Test_DTO_FilterParameterQueryStringDTO_GetSearch(t *testing.T) {
	d := &FilterParameterQueryStringDTO{}
	d.Search = "Foo "
	assert.Equal(t, "%Foo%", d.GetSearch())
}
