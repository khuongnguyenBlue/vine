package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetIDParams(c *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func GetUserID(c *gin.Context) (uint, bool) {
	userID, ok := c.Get("user_id")
	if !ok {
		return 0, false
	}

	uUserId, ok := userID.(uint)
	if !ok {
		return 0, false
	}

	return uUserId, true
}
