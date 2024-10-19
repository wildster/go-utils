package goutils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ServerSuccessResponse(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"error":   "",
	})
}

func ServerErrorResponse(c *gin.Context, err error) {
	Sugar.Errorln(err)

	c.JSON(200, gin.H{
		"success": false,
		"error":   "Internal Error",
	})
}

type AuthRole int

const (
	User AuthRole = iota + 1
	SuperAdmin
)

func (r AuthRole) String() string {
	switch r {
	case SuperAdmin:
		return "super_admin"
	case User:
		return "user"
	default:
		return "unknown"
	}
}

func AuthRoleFromString(s string) AuthRole {
	switch s {
	case "super_admin":
		return SuperAdmin
	case "user":
		return User
	default:
		return User
	}
}

type Auth struct {
	UserId int64
	Role   AuthRole
}

func ExtractAuth(c *gin.Context) *Auth {
	role := c.GetHeader("x-role")
	userId, err := strconv.ParseInt(c.GetHeader("x-user-id"), 10, 64)

	if err != nil || userId <= 0 {
		return nil
	}

	return &Auth{UserId: userId, Role: AuthRoleFromString(role)}
}
