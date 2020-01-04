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
