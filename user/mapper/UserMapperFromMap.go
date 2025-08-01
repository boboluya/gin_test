package mapper

import (
	"errors"
	"fmt"
	"gin_simulate_auth/user/data"
	"time"
)

type UserMapperFromMap struct{}

var dbUsers = data.Users

// Add /*添加用户
func (UserMapperFromMap) Add(user data.User) int64 {
	userName := user.UserName
	hasUser := false
	for _, dbUser := range dbUsers {
		if dbUser.UserName == userName {
			hasUser = true
			break
		}
	}
	if !hasUser {
		fmt.Println("正在添加...")
		time.Sleep(3 * time.Second)
		fmt.Println("添加用户1个成功")
		return 1
	} else {
		fmt.Println("正在添加...")
		time.Sleep(3 * time.Second)
		fmt.Println("添加失败,此用户名已存在")
		return 0
	}
}

func (UserMapperFromMap) BatchAdd(users []data.User) int64 {
	fmt.Println("批量新增...")
	return int64(len(users))
}

// SelectList /*
func (UserMapperFromMap) SelectList(user data.User) []*data.User {
	fmt.Println("正在查询用户:", user.RoleName)
	var results []*data.User
	for _, dbUser := range dbUsers {
		if dbUser.RoleName == user.RoleName {
			results = append(results, dbUser)
		}
	}
	return results
}

func (UserMapperFromMap) SelectOne(user data.User) (*data.User, error) {
	fmt.Println("正在查询用户:", user.UserName)
	for _, dbUser := range dbUsers {
		if dbUser.UserName == user.UserName {
			return dbUser, nil
		}
	}
	return nil, errors.New("此用户不存在")
}

func (UserMapperFromMap) SelectOneById(userId int64) (*data.User, error) {
	fmt.Println("正在查询用户:", userId)
	for _, dbUser := range dbUsers {
		if dbUser.Id == userId {
			return dbUser, nil
		}
	}
	return nil, errors.New("此用户不存在")
}

func (UserMapperFromMap) Update(user data.User) (int64, error) {
	for _, dbUser := range dbUsers {
		if dbUser.Id == user.Id {
			fmt.Println("正在更新...")
			time.Sleep(3 * time.Second)
			fmt.Println("更新用户1个成功: ", dbUser.UserName)
			return 1, nil
		}
	}
	return 0, errors.New("此用户不存在")
}

func (UserMapperFromMap) Delete(user data.User) int64 {
	for _, dbUser := range dbUsers {
		if dbUser.Id == user.Id {
			fmt.Println("正在删除...")
			time.Sleep(3 * time.Second)
			fmt.Println("删除用户1个成功: ", dbUser.UserName)
			return 1
		}
	}
	return 0
}
