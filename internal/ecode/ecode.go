package ecode

import "github.com/binbinly/pkg/errno"

var (
	ErrUserNotFound 				= errno.NewError(20101, "user not exists")
)