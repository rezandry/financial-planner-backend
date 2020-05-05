package service

import model "financial-planner/model/vo"

// BuildResponse used for build response message
func BuildResponse(code int, message string) model.Response {
	var res model.Response
	res.SetMessage(message)
	res.SetCode(code)
	return res
}
