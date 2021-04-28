// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci"
	"github.com/zoom-ci/zoom-ci/server"
	"github.com/zoom-ci/zoom-ci/server/router/route"
	"github.com/zoom-ci/zoom-ci/server/router/system"
	"github.com/zoom-ci/zoom-ci/util/gopath"
	"log"
	"os"
)

var (
	helpFlag    bool
	zoomIniFlag string
	versionFlag bool
	configFile  *goconfig.ConfigFile
	upgradeFlag string
)

func init() {
	gin.SetMode(gin.ReleaseMode)

	flag.BoolVar(&helpFlag, "h", false, "This help")
	flag.StringVar(&zoomIniFlag, "c", "", "Set configuration file `file`")
	flag.BoolVar(&versionFlag, "v", false, "Version number")
	flag.StringVar(&upgradeFlag, "upgrade", "", "Upgrade")

	flag.Usage = usage
	flag.Parse()
}

func usage() {
	fmt.Printf("Usage: zoom [-c filename]\n\nOptions:\n")
	flag.PrintDefaults()
}

func initCfg() {
	var err error
	zoomIni := findZoomIniFile()
	configFile, err = goconfig.LoadConfigFile(zoomIni)
	if err != nil {
		log.Fatalf("load config file failed, %s\n", err.Error())
	}
	zoom.App.ConfigFileHandle = zoomIni
	zoom.App.ConfigHandle = configFile
	outputInfo("Config Loaded", zoomIni)
}

func configIntOrDefault(section, key string, useDefault int) int {
	val, err := configFile.Int(section, key)
	if err != nil {
		return useDefault
	}
	return val
}

func configOrDefault(section, key, useDefault string) string {
	val, err := configFile.GetValue(section, key)
	if err != nil {
		return useDefault
	}
	return val
}

func findZoomIniFile() string {
	if zoomIniFlag != "" {
		return zoomIniFlag
	}
	currPath, _ := gopath.CurrentPath()

	iniFile := currPath + "/zoom.ini"
	if gopath.Exists(iniFile) && gopath.IsFile(iniFile) {
		return iniFile
	}
	defaultConfigFile, _ := zoom.DefaultConfigIniFile.ReadFile("resource/zoom.sample.ini")
	saveDefaultConfigResult := gopath.SaveFile(iniFile, defaultConfigFile)
	if saveDefaultConfigResult != nil {
		outputInfo("Default Config", "load error")
	} else {
		outputInfo("Default Config", "load success")
	}
	return iniFile
}

func outputInfo(tag string, value interface{}) {
	fmt.Printf("%-18s    %v\n", tag+":", value)
}

func welcome() {
	fmt.Println(" _____________________________")
	fmt.Println("       ___                    ")
	fmt.Println("         /                    ")
	fmt.Println(" -------/-----__----__---_--_-")
	fmt.Println("       /    /   ) /   ) / /  )")
	fmt.Println(" ____(_____(___/_(___/_/_/__/_")
	fmt.Println("     /   Make CI/CD Easier.   ")
	fmt.Println(" (_ /                         ")
	fmt.Println("")
	outputInfo("Service", "zoom")
	outputInfo("Version", zoom.Version)
}

func main() {
	if helpFlag {
		flag.Usage()
		os.Exit(0)
	}
	if versionFlag {
		fmt.Printf("zoom/%s\n", zoom.Version)
		os.Exit(0)
	}

	welcome()

	initCfg()

	cfg := &server.Config{
		Serve: &server.ServeConfig{
			Addr:         configOrDefault("serve", "addr", "7002"),
			ReadTimeout:  configIntOrDefault("serve", "read_timeout", 300),
			WriteTimeout: configIntOrDefault("serve", "write_timeout", 300),
			IdleTimeout:  configIntOrDefault("serve", "idle_timeout", 300),
		},
		Db: &server.DbConfig{
			Unix:            configOrDefault("database", "unix", ""),
			Host:            configOrDefault("database", "host", ""),
			Port:            configIntOrDefault("database", "port", 3306),
			Charset:         "utf8mb4",
			User:            configOrDefault("database", "user", ""),
			Pass:            configOrDefault("database", "password", ""),
			DbName:          configOrDefault("database", "dbname", ""),
			Prefix:          configOrDefault("database", "prefix", "zoom_"),
			MaxIdleConns:    configIntOrDefault("database", "max_idle_conns", 100),
			MaxOpenConns:    configIntOrDefault("database", "max_open_conns", 200),
			ConnMaxLifeTime: configIntOrDefault("database", "conn_max_life_time", 500),
		},
		Log: &server.LogConfig{
			Path: configOrDefault("log", "path", "stdout"),
		},
		Zoom: &server.ZoomConfig{
			LocalSpace:  configOrDefault("zoom", "local_space", "/tmp/zoom_data"),
			RemoteSpace: configOrDefault("zoom", "remote_space", "~/.zoom"),
			Cipher:      configOrDefault("zoom", "cipher_key", ""),
			AppHost:     configOrDefault("zoom", "app_host", ""),
			AppName:     configOrDefault("zoom", "app_name", "Zoom-CI"),
		},
		Mail: &server.MailConfig{
			Enable: configIntOrDefault("mail", "enable", 0),
			Smtp:   configOrDefault("mail", "smtp_host", ""),
			Port:   configIntOrDefault("mail", "smtp_port", 465),
			User:   configOrDefault("mail", "smtp_user", ""),
			Pass:   configOrDefault("mail", "smtp_pass", ""),
		},
	}

	outputInfo("Log", cfg.Log.Path)
	outputInfo("Mail Enable", cfg.Mail.Enable)
	outputInfo("HTTP Service", cfg.Serve.Addr)

	if err := zoom.App.Init(cfg); err != nil {
		log.Fatal(err)
	}

	// process upgrade logic
	if upgradeFlag != "" {
		if upgradeFlag == "syncd" {
			system.UpgradeFromSyncd()
		}
	}

	route.RegisterRoute()

	fmt.Println("Start Running...")

	if err := zoom.App.Start(); err != nil {
		log.Fatal(err)
	}
}
