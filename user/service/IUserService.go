package service

import "gin_simulate_auth/user/data"

type IUserService interface {
	Add(user data.User) int64
	BatchAdd(users []data.User) int64
	SelectList(user data.User) []*data.User
	SelectOne(user data.User) (*data.User, error)
	SelectOneById(userId int64) (*data.User, error)
	Update(user data.User) (int64, error)
	Delete(user data.User) int64
}
