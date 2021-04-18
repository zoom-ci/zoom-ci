// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package common

import (
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci/server/module/project"
	"github.com/zoom-ci/zoom-ci/server/render"
)

func InSpaceCheck(c *gin.Context, spaceId int) bool {
	member := &project.Member{
		UserId:  c.GetInt("user_id"),
		SpaceId: spaceId,
	}
	if in := member.MemberInSpace(); !in {
		render.CustomerError(c, render.CODE_ERR_NO_PRIV, "user is not in the project space")
		return false
	}
	return true
}
