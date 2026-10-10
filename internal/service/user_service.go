/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-10-08 09:38:09
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-10-08 11:09:14
 * @FilePath: /WindLiberity/internal/service/user_service.go
 * @Description:
 *
 */
package service

import (
	"context"


	"WindLiberity/internal/model"
	"WindLiberity/pkg/app"

	"github.com/binbinly/pkg/auth"

	"github.com/pkg/errors"
)

// User 用户服务接口
type User interface {
	// UserRegister 用户注册
	UserRegister(ctx context.Context, username, password string, phone int64) (int, error)
	// UsernameLogin 用户名登录
	UsernameLogin(ctx context.Context, username, password string) (*model.UserModel, string, error)
	// UserPhoneLogin 手机号登录
	UserPhoneLogin(ctx context.Context, phone int64) (*model.UserModel, string, error)
	// UserEditPwd 修改密码
	UserEditPwd(ctx context.Context, id int, password string) error
	// UserEdit 修改用户信息
	UserEdit(ctx context.Context, id int, um map[string]any) error
	// UserInfoByID 获取用户详情
	UserInfoByID(ctx context.Context, id int) (*model.UserModel, error)
	// UserLogout 用户登出
	UserLogout(ctx context.Context, id int) error
	// UserSearch 搜索用户
	UserSearch(ctx context.Context, keyword string) ([]*model.UserModel, error)

}

// UserRegister
func (s *Service) UserRegister(ctx context.Context, username, password string, phone int64) (id int, err error) {
	exist, err := s.repo.UserExist(ctx, username, phone)
	if err != nil {
		return 0, errors.Wrapf(err, "[service.user] user exist")
	}
	if exist {
		return 0, ErrUserExisted
	}
	u := &model.UserModel{
		Username: username,
		Password: password,
		Phone:    phone,
		Status:   model.UserStatusNormal,
	}
	id, err = s.repo.UserCreate(ctx, u)
	if err != nil {
		return 0, errors.Wrapf(err, "[service.user] create user")
	}
	return id, nil
}

// UsernameLogin 用户名密码登录
func (s *Service) UsernameLogin(ctx context.Context, username, password string) (*model.UserModel, string, error) {
	// 如果是已经注册用户，则通过用户名获取用户信息
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, "", errors.Wrapf(err, "[service.user] get user by username: %s", username)
	}

	// 否则新建用户信息, 并取得用户信息
	if user.ID == 0 {
		return nil, "", ErrUserNotFound
	}

	if user.Status != model.UserStatusNormal {
		return nil, "", ErrUserFrozen
	}

	// Compare the login password with the user password.
	err = user.Compare(password)
	if err != nil {
		return nil, "", ErrUserNotMatch
	}

	token, err := s.generateToken(ctx, user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
// UserPhoneLogin 手机登录
func (s *Service) UserPhoneLogin(ctx context.Context, phone int64) (*model.UserModel, string, error) {
	// 如果是已经注册用户，则通过手机号获取用户信息
	user, err := s.repo.GetUserByPhone(ctx, phone)
	if err != nil {
		return nil, "", errors.Wrapf(err, "[service.user] get user by phone: %d", phone)
	}

	// 否则新建用户信息, 并取得用户信息
	if user.ID == 0 {
		return nil, "", ErrUserNotFound
	}

	if user.Status != model.UserStatusNormal {
		return nil, "", ErrUserFrozen
	}

	token, err := s.generateToken(ctx, user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

// UserEdit 修改用户信息
func (s *Service) UserEdit(ctx context.Context, id int, um map[string]any) error {
	return s.repo.UserUpdate(ctx, id, um)
}

// UserEditPwd 修改用户密码
func (s *Service) UserEditPwd(ctx context.Context, id int, password string) error {
	user, err := s.userinfo(ctx, id)
	if err != nil {
		return err
	}
	user.Password = password
	if err = s.repo.UserUpdatePwd(ctx, user); err != nil {
		return errors.Wrapf(err, "[service.user] update user pwd by id:%v", id)
	}
	return nil
}

// UserInfoByID 获取用户信息
func (s *Service) UserInfoByID(ctx context.Context, id int) (*model.UserModel, error) {
	return s.userinfo(ctx, id)
}

// UserLogout 用户登出
func (s *Service) UserLogout(ctx context.Context, id int) error {
	return s.rdb.Del(ctx, app.BuildUserTokenKey(id)).Err()
}

// UserSearch 搜索用户
func (s *Service) UserSearch(ctx context.Context, keyword string) ([]*model.UserModel, error) {
	return s.repo.GetUsersByKeyword(ctx, keyword)
}

// userinfo 获取用户模型
func (s *Service) userinfo(ctx context.Context, id int) (*model.UserModel, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.user] find id: %d", id)
	}
	if user.ID == 0 {
		return nil, ErrUserNotFound
	}
	if user.Status != model.UserStatusNormal {
		return nil, ErrUserFrozen
	}
	return user, nil
}

func (s *Service) batchUserinfo(ctx context.Context, ids []int) ([]*model.UserModel, error) {
	users, err := s.repo.GetUsersByIds(ctx, ids)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.user] users info by ids:%v", ids)
	}
	return users, nil
}

// generateToken 生成token
func (s *Service) generateToken(ctx context.Context, uid int) (string, error) {
	// 签名
	payload := map[string]any{"user_id": uid}
	token, err := auth.Sign(ctx, payload, s.opts.jwtSecret, s.opts.jwtTimeout)
	if err != nil {
		return "", errors.Wrapf(err, "[service.user] gen token sign")
	}
	return token, nil
}