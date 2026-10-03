package app

import (
	"log"
	"time"

	"WindLiberity/pkg/config"

	"github.com/binbinly/pkg/storage/orm"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Default orm.Config
}


// InitDB init dbs
func InitDB() *gorm.DB {
	var cfg = &DBConfig{}
	if err := LoadDBConf(cfg); err != nil {
		log.Fatalf("load db conf err: %v", err)
	}

	DB = orm.NewDB(&cfg.Default)
	return DB
}

func InitBasicDB(driver, dsn string) *gorm.DB {
	DB = orm.NewDB(&orm.Config{
		Driver: driver,
		Dsn: 	dsn,
	})

	return DB
}


// LoadDBConf load dbs config
func LoadDBConf(cfg *DBConfig) error {
	if err := config.Load("database", cfg,func(v *viper.Viper) {
		v.SetDefault("default", map[string]any{
			"Driver":          "mysql",
			"Host":            "127.0.0.1",
			"Port":            3306,
			"User":            "root",
			"Password":        "root",
			"Database":        "WindLiberity",
			"Debug":           true,
			"MaxIdleConn":     10,
			"MaxOpenConn":     100,
			"ConnMaxLifeTime": 100 * time.Second,
		})
		v.BindEnv("default.driver", "WINDLIBERITY_DB_DRIVER")
		v.BindEnv("default.dsn", "WINDLIBERITY_DB_DSN")
		v.BindEnv("default.host", "WINDLIBERITY_DB_HOST")
		v.BindEnv("default.port", "WINDLIBERITY_DB_PORT")
		v.BindEnv("default.user", "WINDLIBERITY_DB_USER")
		v.BindEnv("default.password", "WINDLIBERITY_DB_PASSWORD")
		v.BindEnv("default.database", "WINDLIBERITY_DB_DATABASE")
	}); err != nil {
		return err
	}
	return nil
}
