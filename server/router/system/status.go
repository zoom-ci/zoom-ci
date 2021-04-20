// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci"
	"github.com/zoom-ci/zoom-ci/server/render"
)

func Status(c *gin.Context) {
	render.JSON(c, gin.H{
		"localSpacePath": zoom.App.Config.Zoom.LocalSpace,
		"remoteSpacePath": zoom.App.Config.Zoom.RemoteSpace,
		"currentConfigFilePath": zoom.App.ConfigFileHandle,
		"currentZoomVersion": zoom.Version,
	})
}
