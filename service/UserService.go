package UserService

import (
	"web-api/repo"
	"web-api/models/user"
)

func Save(user *user.User) int64 {
	res, err := UserRepository.Save(user)
	if err != nil {
		panic(err)
		return res
	}
	return res
}

func GetUserById(id int64) user.User{
	user, err := UserRepository.GetById(id)
	if err != nil {
		panic(err)
		return user
	}
	return user
}
