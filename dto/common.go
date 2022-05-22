package dto

type IdUriParameterDTO struct {
	ID uint64 `uri:"id"`
}

type HashIdUriParameterDTO struct {
	ID string `uri:"id"`
}
