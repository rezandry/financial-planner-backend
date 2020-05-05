package repository

import (
	"encoding/json"
	model "financial-planner/model/vo"
)

// CreateUser : Insert new user data to db
func CreateUser(user model.User) bool {
	redisConn := RedisInit()
	defer redisConn.Close()

	data, _ := redisConn.Do("GET", user.Username)

	if data != nil {
		return false
	}

	userDump, _ := json.Marshal(user)
	redisConn.Do("SET", user.Username, string(userDump))
	return true
}
