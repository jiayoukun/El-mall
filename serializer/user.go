package serializer

import "test15/model"

type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
}

func Builduser(user *model.User) User {
	return User{
		ID: user.ID,
		UserName: user.Username,
		Nickname: user.Nickname,
		Email: user.Email,
		CreatedAt: user.CreatedAt.Unix(),
	}
}