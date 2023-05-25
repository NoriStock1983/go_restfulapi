package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AllUsers(c *gin.Context) {
	fmt.Print("All Users")
}
