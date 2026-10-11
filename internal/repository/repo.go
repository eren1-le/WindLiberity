/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-28 21:06:42
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-10-08 09:54:33
 * @FilePath: /WindLiberity/internal/repository/repo.go
 * @Description:
 *
 */
package repository

import (
	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/repo"
	"gorm.io/gorm"
)

var _ IRepo = (*Repo)(nil)

// IRepo 数据仓库接口
type IRepo interface {
	User
}

// Repo dbs struct
type Repo struct {
	repo.Repo
}

// New new a Dao and return
func New(db *gorm.DB, c cache.Cache) IRepo {
	return &Repo{repo.Repo{
		DB:    db,
		Cache: c,
	}}
}
