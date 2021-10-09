package models

type UserMapper struct {
}

func (pr *UserMapper) MapForInsert(user *User) map[string]interface{} {
	row := map[string]interface{}{
		"name":        user.Name,
		"email":       user.Email,
		"password":    user.Password,
		"created_at":  user.CreatedAt,
		"description": user.Description.String,
	}
	return row
}

func (pr *UserMapper) MapForUpdate(user *User) map[string]interface{} {
	row := map[string]interface{}{
		"name":        user.Name,
		"email":       user.Email,
		"password":    user.Password,
		"created_at":  user.CreatedAt,
		"description": user.Description.String,
	}
	return row
}
