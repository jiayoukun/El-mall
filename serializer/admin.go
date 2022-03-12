package serializer

import "test15/model"

type Admin struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
}


func Builduadmin(user *model.Admin) Admin {
	return Admin{
		ID: user.ID,
		UserName: user.Username,
		Nickname: user.Nickname,
		Email: user.Email,
		CreatedAt: user.CreatedAt.Unix(),
	}
}