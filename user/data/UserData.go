package data

type User struct {
	Id       int64  `json:"id" form:"id"`
	UserName string `json:"userName" form:"userName"`
	Password string `json:"password" form:"password"`
	RoleId   int64  `json:"RoleId" form:"RoleId"`
	RoleName string `json:"roleName" form:"roleName"`
	Avatar   string `json:"avatar" form:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
}

var Users []*User

func init() {
	Users = []*User{
		{
			Id:       1,
			UserName: "admin",
			Password: "admin123", // 注意：真实项目请加密
			RoleId:   1,
			RoleName: "管理员",
			Avatar:   "https://example.com/avatar1.png",
			Mobile:   "13800138000",
		},
		{
			Id:       2,
			UserName: "lulu",
			Password: "lulu123",
			RoleId:   2,
			RoleName: "普通用户",
			Avatar:   "https://example.com/avatar2.png",
			Mobile:   "13900139000",
		},
		{
			Id:       3,
			UserName: "guest",
			Password: "guest123",
			RoleId:   3,
			RoleName: "游客",
			Avatar:   "https://example.com/avatar3.png",
			Mobile:   "13700137000",
		},
	}
}
