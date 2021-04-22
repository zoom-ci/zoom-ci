// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package system

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/zoom-ci/zoom-ci"
	"github.com/zoom-ci/zoom-ci/server"
	"github.com/zoom-ci/zoom-ci/server/model"
	"strconv"
)

type Install struct {
	MysqlHost     string
	MysqlPort     int
	MysqlUsername string
	MysqlPassword string
	MysqlDbname   string
}

func (install *Install) Install() error {
	// do install
	// install db
	zoom.App.DB = server.NewDatabase(&server.DbConfig{
		Unix:            zoom.App.Config.Db.Unix,
		Host:            install.MysqlHost,
		Port:            install.MysqlPort,
		Charset:         zoom.App.Config.Db.Charset,
		User:            install.MysqlUsername,
		Pass:            install.MysqlPassword,
		DbName:          install.MysqlDbname,
		MaxIdleConns:    zoom.App.Config.Db.MaxIdleConns,
		MaxOpenConns:    zoom.App.Config.Db.MaxOpenConns,
		ConnMaxLifeTime: zoom.App.Config.Db.ConnMaxLifeTime,
	})
	err := zoom.App.DB.Open()
	if err != nil {
		fmt.Println(err)
		return err
	}

	zoom.App.DB.DbHandler.Exec("CREATE TABLE IF NOT EXISTS `zoom_deploy_apply` ( " +
		"`id` int UNSIGNED NOT NULL AUTO_INCREMENT," +
		"`space_id` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`project_id` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`name` varchar(100) NOT NULL DEFAULT ''," +
		"`description` varchar(500) NOT NULL DEFAULT ''," +
		"`branch_name` varchar(100) NOT NULL DEFAULT ''," +
		"`commit_version` varchar(100) NOT NULL DEFAULT ''," +
		"`audit_status` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`audit_refusal_reason` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT ''," +
		"`status` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`user_id` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`rollback_id` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`rollback_apply_id` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`is_rollback_apply` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`ctime` int UNSIGNED NOT NULL DEFAULT '0'," +
		"PRIMARY KEY (`id`)," +
		"KEY `idx_space_project` (`space_id`,`project_id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")

	zoom.App.DB.DbHandler.Exec("CREATE TABLE IF NOT EXISTS `zoom_deploy_build` (" +
		"`id` int UNSIGNED NOT NULL AUTO_INCREMENT," +
		"`apply_id` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`start_time` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`finish_time` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`status` int UNSIGNED NOT NULL DEFAULT '1'," +
		"`tar` varchar(2000) NOT NULL DEFAULT ''," +
		"`output` mediumtext NOT NULL," +
		"`errmsg` text NOT NULL," +
		"`ctime` int UNSIGNED NOT NULL DEFAULT '0'," +
		"PRIMARY KEY (`id`)," +
		"UNIQUE KEY `idx_apply_id` (`apply_id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")

	zoom.App.DB.DbHandler.Exec("CREATE TABLE IF NOT EXISTS `zoom_deploy_task` (" +
		"`id` int UNSIGNED NOT NULL AUTO_INCREMENT," +
		"`apply_id` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`group_id` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`status` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`content` text NOT NULL," +
		"`ctime` int UNSIGNED NOT NULL," +
		"PRIMARY KEY (`id`)," +
		"KEY `idx_apply_id` (`apply_id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")

	zoom.App.DB.DbHandler.Exec("CREATE TABLE IF NOT EXISTS `zoom_project` (" +
		"`id` int UNSIGNED NOT NULL AUTO_INCREMENT," +
		"`space_id` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`name` varchar(100) NOT NULL DEFAULT ''," +
		"`description` varchar(500) NOT NULL DEFAULT ''," +
		"`project_type` tinyint  NOT NULL DEFAULT '1'," +
		"`need_audit` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`status` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`repo_url` varchar(500) NOT NULL DEFAULT ''," +
		"`deploy_mode` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`repo_branch` varchar(100) NOT NULL DEFAULT ''," +
		"`pre_release_cluster` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`online_cluster` varchar(2000) NOT NULL DEFAULT ''," +
		"`deploy_user` varchar(50) NOT NULL DEFAULT ''," +
		"`deploy_path` varchar(500) NOT NULL DEFAULT ''," +
		"`build_script` text NOT NULL," +
		"`build_hook_script` text NOT NULL," +
		"`deploy_hook_script` text NOT NULL," +
		"`pre_deploy_cmd` text NOT NULL," +
		"`after_deploy_cmd` text NOT NULL," +
		"`audit_notice` varchar(2000) DEFAULT NULL," +
		"`deploy_notice` varchar(2000) DEFAULT NULL," +
		"`ctime` int UNSIGNED NOT NULL DEFAULT '0'," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")

	zoom.App.DB.DbHandler.Exec("CREATE TABLE IF NOT EXISTS `zoom_project_member` (" +
		"`id` int UNSIGNED NOT NULL AUTO_INCREMENT," +
		"`space_id` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`user_id` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`ctime` int UNSIGNED NOT NULL DEFAULT '0'," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")

	zoom.App.DB.DbHandler.Exec("CREATE TABLE IF NOT EXISTS `zoom_project_space` (" +
		"`id` int UNSIGNED NOT NULL AUTO_INCREMENT," +
		"`name` varchar(100) NOT NULL DEFAULT ''," +
		"`description` varchar(2000) NOT NULL DEFAULT ''," +
		"`ctime` int UNSIGNED NOT NULL DEFAULT '0'," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")

	zoom.App.DB.DbHandler.Exec("CREATE TABLE IF NOT EXISTS `zoom_server` (" +
		"`id` int UNSIGNED NOT NULL AUTO_INCREMENT," +
		"`group_id` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`name` varchar(100) NOT NULL DEFAULT ''," +
		"`ip` varchar(100) NOT NULL DEFAULT ''," +
		"`ssh_port` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`ctime` int UNSIGNED NOT NULL DEFAULT '0'," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")

	zoom.App.DB.DbHandler.Exec("CREATE TABLE IF NOT EXISTS `zoom_server_group` (" +
		"`id` int UNSIGNED NOT NULL AUTO_INCREMENT," +
		"`name` varchar(100) NOT NULL DEFAULT ''," +
		"`ctime` int NOT NULL DEFAULT '0'," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")

	zoom.App.DB.DbHandler.Exec("CREATE TABLE IF NOT EXISTS `zoom_user` (" +
		"`id` int UNSIGNED NOT NULL AUTO_INCREMENT," +
		"`role_id` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`username` varchar(20) NOT NULL DEFAULT ''," +
		"`password` char(32) NOT NULL DEFAULT ''," +
		"`salt` char(10) NOT NULL DEFAULT ''," +
		"`truename` varchar(10) NOT NULL DEFAULT ''," +
		"`mobile` varchar(20) NOT NULL DEFAULT ''," +
		"`email` varchar(500) NOT NULL DEFAULT ''," +
		"`status` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`last_login_time` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`last_login_ip` varchar(50) NOT NULL DEFAULT ''," +
		"`ctime` int UNSIGNED NOT NULL DEFAULT '0'," +
		"PRIMARY KEY (`id`)," +
		"KEY `idx_username` (`username`)," +
		"KEY `idx_email` (`email`(20))" +
		") ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")

	user := &model.User{}
	if ok := user.Get(1); !ok || user.ID == 0 {
		zoom.App.DB.DbHandler.Exec("INSERT INTO `zoom_user` (`id`, `role_id`, `username`, `password`, `salt`, `truename`, `mobile`, `email`, `status`, `last_login_time`, `last_login_ip`, `ctime`) VALUES(1, 1, 'admin', '1583514ddbb5ad4e789f6e664f7814ee', 'e6NukxZ0MX', 'Zoom', '', 'admin@zoom.com', 1, 0, '', 0);")
	}

	zoom.App.DB.DbHandler.Exec("CREATE TABLE IF NOT EXISTS `zoom_user_role` (" +
		"`id` int UNSIGNED NOT NULL AUTO_INCREMENT," +
		"`name` varchar(100) NOT NULL DEFAULT ''," +
		"`privilege` varchar(2000) NOT NULL DEFAULT ''," +
		"`ctime` int UNSIGNED NOT NULL DEFAULT '0'," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")

	userRole := &model.UserRole{}
	if ok := userRole.Get(1); !ok || userRole.ID == 0 {
		zoom.App.DB.DbHandler.Exec("INSERT INTO `zoom_user_role` (`id`, `name`, `privilege`, `ctime`) VALUES(1, '管理员', '2001,2002,2003,2004,2100,2101,2102,2201,2202,2203,2204,2205,2206,2207,2208,3001,3002,3004,3003,3101,3102,3103,3104,4001,4002,4003,4004,4101,4102,4103,4104,1001,1002,1006,1003,1004,1005,5001', 0);")
	}
	zoom.App.DB.DbHandler.Exec("CREATE TABLE IF NOT EXISTS `zoom_user_token` (" +
		"`id` int UNSIGNED NOT NULL AUTO_INCREMENT," +
		"`user_id` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`token` varchar(100) NOT NULL DEFAULT ''," +
		"`expire` int UNSIGNED NOT NULL DEFAULT '0'," +
		"`ctime` int UNSIGNED NOT NULL DEFAULT '0'," +
		"PRIMARY KEY (`id`)," +
		"UNIQUE KEY `idx_user_id` (`user_id`)" +
		") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;")

	// update config
	zoom.App.Config.Db.Host = install.MysqlHost
	zoom.App.Config.Db.Port = install.MysqlPort
	zoom.App.Config.Db.User = install.MysqlUsername
	zoom.App.Config.Db.Pass = install.MysqlPassword
	zoom.App.Config.Db.DbName = install.MysqlDbname

	// update config file
	zoom.App.ConfigHandle.SetValue("database", "host", install.MysqlHost)
	zoom.App.ConfigHandle.SetValue("database", "port", strconv.Itoa(install.MysqlPort))
	zoom.App.ConfigHandle.SetValue("database", "user", install.MysqlUsername)
	zoom.App.ConfigHandle.SetValue("database", "password", install.MysqlPassword)
	zoom.App.ConfigHandle.SetValue("database", "dbname", install.MysqlDbname)

	goconfig.SaveConfigFile(zoom.App.ConfigHandle, zoom.App.ConfigFileHandle)

	return nil
}
