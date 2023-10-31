package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserId(c *gin.Context) {
	fmt.Println(c.Query("foo"))
	fmt.Println(c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, id)
}
