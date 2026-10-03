package repository

import (
	"context"
	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/repo"
	"gorm.io/gorm"

	"WindLiberity/internal/model"
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
