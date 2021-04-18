// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zoom

import (
	"embed"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/zoom-ci/zoom-ci/server"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci/util/golog"
	"github.com/zoom-ci/zoom-ci/util/gopath"
)

var (
	App *zoom
)

//go:embed web/dist
var WebPath embed.FS

//go:embed resource/zoom.sample.ini
var DefaultConfigIniFile embed.FS

const (
	Version = "v1.0.0"
)

func init() {
	App = newZoom()
}

type zoom struct {
	Gin              *gin.Engine
	DB               *server.DB
	Logger           *golog.Logger
	Mail             *server.SendMail
	LocalSpace       string
	LocalTmpSpace    string
	LocalTarSpace    string
	RemoteSpace      string
	CipherKey        []byte
	AppHost          string
	Config           *server.Config
	ConfigFileHandle string
	ConfigHandle     *goconfig.ConfigFile
	ZoomInstalled    bool
}

func newZoom() *zoom {
	return &zoom{
		Gin: gin.New(),
	}
}

func (s *zoom) Init(cfg *server.Config) error {
	s.Config = cfg
	if s.Config.Db.User != "" && s.Config.Db.Pass != "" && s.Config.Db.DbName != "" {
		s.ZoomInstalled = true
		if err := s.RegisterOrm(); err != nil {
			return err
		}
	} else {
		// set not installed by db config
		s.ZoomInstalled = false
	}
	s.registerMail()
	s.registerLog()

	if err := s.initEnv(); err != nil {
		return err
	}
	return nil
}

func (s *zoom) Start() error {
	return s.Gin.Run(s.Config.Serve.Addr)
}

func (s *zoom) RegisterOrm() error {
	s.DB = server.NewDatabase(s.Config.Db)
	return s.DB.Open()
}

func (s *zoom) registerLog() {
	var loggerHandler io.Writer
	switch s.Config.Log.Path {
	case "stdout":
		loggerHandler = os.Stdout
	case "stderr":
		loggerHandler = os.Stderr
	case "":
		loggerHandler = os.Stdout
	default:
		loggerHandler = golog.NewFileHandler(s.Config.Log.Path)
	}
	s.Logger = golog.New(loggerHandler)
}

func (s *zoom) registerMail() {
	sendmail := &server.SendMail{
		Enable: s.Config.Mail.Enable,
		Smtp:   s.Config.Mail.Smtp,
		Port:   s.Config.Mail.Port,
		User:   s.Config.Mail.User,
		Pass:   s.Config.Mail.Pass,
	}
	s.Mail = server.NewSendMail(sendmail)
}

func (s *zoom) initEnv() error {
	s.AppHost = s.Config.Zoom.AppHost
	s.LocalSpace = s.Config.Zoom.LocalSpace
	s.LocalTmpSpace = s.LocalSpace + "/tmp"
	s.LocalTarSpace = s.LocalSpace + "/tar"

	if err := gopath.CreatePath(s.LocalSpace); err != nil {
		return err
	}
	if err := gopath.CreatePath(s.LocalTmpSpace); err != nil {
		return err
	}
	if err := gopath.CreatePath(s.LocalTarSpace); err != nil {
		return err
	}

	s.RemoteSpace = s.Config.Zoom.RemoteSpace
	if s.Config.Zoom.Cipher == "" {
		return errors.New("zoom config 'Cipher' not setting")
	}
	dec, err := base64.StdEncoding.DecodeString(s.Config.Zoom.Cipher)

	if err != nil {
		return errors.New(fmt.Sprintf("decode Cipher failed, %s", err.Error()))
	}
	s.CipherKey = dec

	return nil
}
