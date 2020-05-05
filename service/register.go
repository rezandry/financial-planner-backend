package service

import (
	"financial-planner/constanta"
	model "financial-planner/model/vo"
	"financial-planner/repository"

	"github.com/gin-gonic/gin"
)

// Register route
func Register(c *gin.Context) {
	res := BuildResponse(constanta.SuccessCode, constanta.SuccessMessage)
	httpStatus := 200

	var userData model.User
	c.Bind(&userData)

	statusInsert := repository.CreateUser(userData)
	if !statusInsert {
		res = BuildResponse(constanta.UserExistCode, constanta.UserExistMessage)
	}

	c.JSON(httpStatus, res)
}
