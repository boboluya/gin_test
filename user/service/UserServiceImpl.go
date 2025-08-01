package service

import (
	"gin_simulate_auth/user/data"
	"gin_simulate_auth/user/mapper"
)

type UserServiceImpl struct {
	userMapper mapper.IUserMapper
}

func NewUserService(userMapper mapper.IUserMapper) *UserServiceImpl {
	return &UserServiceImpl{userMapper: userMapper}
}

func (u *UserServiceImpl) Add(user data.User) int64 {
	return u.userMapper.Add(user)
}

func (u *UserServiceImpl) BatchAdd(users []data.User) int64 {
	return u.userMapper.BatchAdd(users)
}

func (u *UserServiceImpl) SelectList(user data.User) []*data.User {
	return u.userMapper.SelectList(user)
}

func (u *UserServiceImpl) SelectOne(user data.User) (*data.User, error) {
	return u.userMapper.SelectOne(user)
}

func (u *UserServiceImpl) SelectOneById(userId int64) (*data.User, error) {
	return u.userMapper.SelectOneById(userId)
}

func (u *UserServiceImpl) Update(user data.User) (int64, error) {
	return u.userMapper.Update(user)
}

func (u *UserServiceImpl) Delete(user data.User) int64 {
	return u.userMapper.Delete(user)
}
