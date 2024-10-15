package goutils

import "github.com/gin-gonic/gin"

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
