package service

import (
	model "financial-planner/model/vo"
	"financial-planner/repository"

	"github.com/gin-gonic/gin"
)

// Register route
func Register(c *gin.Context) {
	res := BuildResponse(1, "Success")
	httpStatus := 200

	var userData model.User
	c.Bind(&userData)

	statusInsert := repository.CreateUser(userData)
	if !statusInsert {
		res = BuildResponse(11, "User have been used")
	}

	c.JSON(httpStatus, res)
}
