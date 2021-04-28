// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci"
	"github.com/zoom-ci/zoom-ci/server/model"
	"github.com/zoom-ci/zoom-ci/server/module/system"
	"github.com/zoom-ci/zoom-ci/server/render"
)

type InstallBind struct {
	UserName     string `form:"user_name" binding:"required"`
	UserPassword string `form:"user_password" binding:"required"`
	UserEmail    string `form:"user_email" binding:"required"`

	MysqlHost     string `form:"mysql_host" binding:"required"`
	MysqlPort     int    `form:"mysql_port" binding:"required"`
	MysqlUsername string `form:"mysql_username" binding:"required"`
	MysqlPassword string `form:"mysql_password" binding:"required"`
	MysqlDbname   string `form:"mysql_dbname" binding:"required"`
}

func Install(c *gin.Context) {
	if zoom.App.ZoomInstalled == true {
		render.AppError(c, "installed")
		return
	}
	var form InstallBind
	if err := c.ShouldBind(&form); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	install := &system.Install{
		UserName:     form.UserName,
		UserPassword: form.UserPassword,
		UserEmail:    form.UserEmail,

		MysqlHost:     form.MysqlHost,
		MysqlPort:     form.MysqlPort,
		MysqlUsername: form.MysqlUsername,
		MysqlPassword: form.MysqlPassword,
		MysqlDbname:   form.MysqlDbname,
	}

	if err := install.Install(); err != nil {
		render.CustomerError(c, render.CODE_ERR_INSTALL_FAILED, err.Error())
		return
	}

	zoom.App.ZoomInstalled = true

	systemInfo := map[string]interface{}{
		"is_installed": zoom.App.ZoomInstalled,
	}

	render.JSON(c, systemInfo)
}

func InstallStatus(c *gin.Context) {
	// check the user whose id is 1 for judge the install status
	if zoom.App.ZoomInstalled {
		user := &model.User{}
		if ok := user.Get(1); !ok || user.ID == 0 {
			zoom.App.ZoomInstalled = false
		}
	}
	render.JSON(c, gin.H{
		"is_installed": zoom.App.ZoomInstalled,
	})
}

func UpgradeFromSyncd() {
	fmt.Print("Start Upgrading from Syncd2.0 ... ")
	result := system.UpgradeFromSyncd()
	if result{
		fmt.Println("OK")
	} else {
		fmt.Println("no need!")
	}
}
