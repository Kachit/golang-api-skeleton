package dto

import "strings"

type IdParameterPathDTO struct {
	ID uint64 `uri:"id"`
}

type HashParameterPathDTO struct {
	Code string `uri:"hash"`
}

type FilterParameterQueryStringDTO struct {
	Limit  uint64            `form:"limit"`
	Offset uint64            `form:"offset"`
	Search string            `form:"q"`
	Sort   map[string]string `form:"sort"`
}

func (f *FilterParameterQueryStringDTO) HasSearch() bool {
	return f.Search != ""
}

func (f *FilterParameterQueryStringDTO) GetSearch() string {
	return "%" + strings.TrimSpace(f.Search) + "%"
}
