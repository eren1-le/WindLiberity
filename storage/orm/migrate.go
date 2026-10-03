/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-30 01:29:37
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-09-30 01:29:56
 * @FilePath: /WindLiberity/storage/orm/migrate.go
 * @Description:
 *
 */
package orm

import "time"

// Migration 数据迁移
type Migration struct {
	Version   string    `gorm:"primaryKey"`
	ApplyTime time.Time `gorm:"autoCreateTime"`
}

// TableName 表名
func (Migration) TableName() string {
	return "migration"
}
