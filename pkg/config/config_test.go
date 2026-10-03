/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-22 18:41:08
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-09-22 18:53:21
 * @FilePath: /WindLiberity/pkg/config/config_test.go
 * @Description:
 *
 */
package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	type config struct {
		Name string
		Addr string
		Username string
		Password string
	}

	var dbConf config
	c := New(WithConfigDir("../../test/config/"))
	if err := c.Load("database", &dbConf, nil); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, dbConf.Name, "WindLiberity")
	assert.Equal(t, dbConf.Addr, "127.0.0.1:3306")
	assert.Equal(t, dbConf.Username, "root")
	assert.Equal(t, dbConf.Password, "root")	
}