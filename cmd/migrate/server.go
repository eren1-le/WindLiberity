/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-10-09 16:16:31
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-10-10 19:47:48
 * @FilePath: /WindLiberity/cmd/migrate/server.go
 * @Description:
 *
 */
package migrate

import (
	"WindLiberity/cmd/migrate/migration"
	_ "WindLiberity/cmd/migrate/migration/version"
	"WindLiberity/pkg/app"
	"WindLiberity/storage/orm"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"text/template"
	"time"

	"github.com/binbinly/pkg/util/xfile"

	"github.com/spf13/cobra"
)

var (
	dsn      string
	driver   string
	generate bool
	StartCmd = &cobra.Command{
		Use:     "migrate",
		Short:   "Initialize the database",
		Example: "WindLiberity migrate",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&dsn, "dsn", "d",
    "root:root@tcp(127.0.0.1:3306)/WindLiberity?charset=utf8mb4&parseTime=true&loc=Local",
    "dbs dsn data source name")
	StartCmd.PersistentFlags().StringVarP(&driver, "driver", "t", "mysql", "db driver")
	StartCmd.PersistentFlags().BoolVarP(&generate, "generate", "g", false, "generate migration file")
}
func run() {
	if !generate {
		fmt.Println(`start init`)
		initDB()
	} else {
		fmt.Println(`generate migration file`)
		err := genFile()
		if err != nil {
			log.Fatal("err", err)
		}
	}
}

func initDB() {
	app.InitBasicDB(driver, dsn)
	fmt.Println("数据库迁移开始")
	_ = migrateModel()
}

func migrateModel() error {
	db := app.DB
	err := db.Debug().AutoMigrate(new(orm.Migration))
	if err != nil {
		return err
	}
	migration.Migrate.SetDb(db.Debug())
	migration.Migrate.Migrate()
	return err
}

func genFile() error {
	t1, err := template.ParseFiles("cmd/migrate/migrate.template")
	if err != nil {
		return err
	}
	m := map[string]string{}
	m["GenerateTime"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	m["Package"] = "version"
	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	if err != nil {
		return err
	}
	return xfile.Create(b1, "./cmd/migrate/migration/version/"+m["GenerateTime"]+"_migrate.go")
}
