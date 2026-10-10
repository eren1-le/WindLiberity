package version

import (
	"runtime"

	"WindLiberity/cmd/migrate/migration"
	"WindLiberity/internal/model"
	"WindLiberity/storage/orm"

	"gorm.io/gorm"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _data_model)
}

func _data_model(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Debug().Migrator().AutoMigrate(
			new(model.UserModel),
		)
		if err != nil {
			return err
		}
		return tx.Create(&orm.Migration{
			Version: version,
		}).Error
	})
}