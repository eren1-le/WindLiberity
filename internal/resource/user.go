
package resource

import (
	"WindLiberity/internal/model"
	"WindLiberity/pkg/app"
)

// UserResponse 用户响应结构
type UserResponse struct {
	ID       int    `json:"id"`
	Phone    int64  `json:"phone"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Gender   int8   `json:"gender"`
}

// UserResource

func UserResource(user *model.UserModel) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Phone:    user.Phone,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Avatar:   app.BuildResUrl(user.Avatar),
		Gender:   user.Gender,
	}
}
// UserBasicResource transform
func UserBasicResource(user *model.UserModel) *model.User {
	return userBasic(user)
}


// UserListResource
func UserListResource(users []*model.UserModel) []* UserResponse {
	if len(users) == 0 {
		return []*UserResponse{}

		
	}
	us := make([] *UserResponse, 0, len(users))
	for _, u := range users {
		us = append(us, UserResource(u))
	}

	return us
}

// UserToMap
func UserToMap(users []*model.UserModel) (m map[int]*model.User) {
	m = make(map[int]*model.User, len(users))
	for _, user := range users {
		m[user.ID] = userBasic(user)
	}

	return m
}
// UserBasic
func userBasic(user *model.UserModel) *model.User {
	name := user.Username
	if user.Nickname != "" {
		name = user.Nickname
	}

	return &model.User{
		ID: 	user.ID,
		Name:	name,
		Avatar: app.BuildResUrl(user.Avatar),
	}
}