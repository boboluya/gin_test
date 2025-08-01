太好了！来点**贴近实际项目的实战练习**！

我们设计一个 **“用户管理中心”** 的简化版，功能完整、结构清晰，适合练手 Gin 的核心能力：

---

## 🎯 项目名称：`用户管理 API 服务（模拟版）`

> 不连数据库，数据存在内存（`map`），支持 CRUD + 分页 + 搜索 + 登录鉴权 + 文件上传头像 + 中间件

---

## ✅ 功能需求（贴近真实项目）

| 功能 | 说明 |
|------|------|
| 🔐 用户登录 | `POST /api/login`，返回 token（模拟） |
| 📋 获取用户列表 | `GET /api/users`，支持分页、关键字搜索 |
| 👤 获取用户详情 | `GET /api/users/:id` |
| ➕ 创建用户 | `POST /api/users`，带头像上传 |
| ✏️ 更新用户 | `PUT /api/users/:id` |
| ❌ 删除用户 | `DELETE /api/users/:id` |
| 🛡️ 权限控制 | 非登录用户不能操作（除登录外） |
| 📁 头像上传 | 上传文件并保存到本地 `/uploads` |
| 📄 静态文件访问 | 上传的头像可直接浏览器访问 |

---

## 🧱 数据结构设计（模拟）

```go
type User struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Role     string `json:"role"`     // admin/user
    Avatar   string `json:"avatar"`   // 头像文件名
    CreatedAt string `json:"created_at"`
}
```

内存存储：
```go
var users = map[string]User{
    "1": {ID: "1", Name: "admin", Email: "admin@example.com", Role: "admin", CreatedAt: "2024-01-01"},
}
```

---

## 📡 接口列表

| 方法 | 路径 | 说明 | 需要登录 |
|------|------|------|--------|
| POST | `/api/login` | 登录，返回 token | ❌ |
| GET  | `/api/users` | 分页查询用户，支持 `?keyword=xxx&page=1&limit=10` | ✅ |
| GET  | `/api/users/:id` | 查看用户详情 | ✅ |
| POST | `/api/users` | 创建用户（含头像上传） | ✅（仅 admin） |
| PUT  | `/api/users/:id` | 更新用户信息 | ✅（仅 admin 或自己） |
| DELETE | `/api/users/:id` | 删除用户 | ✅（仅 admin） |
| GET  | `/uploads/*filepath` | 访问上传的头像 | ❌ |

---

## 🔐 鉴权规则（模拟）

- 登录后返回一个 `token=mock-token-123`
- 所有请求通过 `Authorization: Bearer mock-token-123` 验证
- 中间件校验 token 是否有效
- `Role = admin` 才能创建/删除用户
- 普通用户只能查看和更新自己（进阶可做，先做 admin 全权）

---

## 📁 项目结构建议

```
project/
├── main.go
├── uploads/          // 存放上传头像
└── data/             // 可选：模拟数据初始化
```

---

## 🧩 技术点覆盖（你将练习到）

| 技术点 | 是否包含 |
|--------|----------|
| 路由分组 | ✅ |
| 中间件（鉴权） | ✅ |
| JSON 绑定与校验 | ✅ |
| 文件上传 | ✅ |
| 静态文件服务 | ✅ |
| 路径参数 `:id` | ✅ |
| 查询参数 `?page=` | ✅ |
| 模拟数据存储 | ✅ |
| 模拟 token 鉴权 | ✅ |
| 错误处理 | ✅ |

---

## 🚀 启动命令建议

```bash
go run main.go
```

服务启动后：

- 访问：`http://localhost:8080/api/users`（需先登录）
- 上传头像：用 Postman 传 `multipart/form-data`，字段名 `avatar`

---

## 🧪 测试建议顺序

1. ✅ 先实现 `/api/login`，拿到 token
2. ✅ 实现 `/api/users` 列表（先不分页）
3. ✅ 加分页和搜索
4. ✅ 实现 `/api/users/:id` 查看
5. ✅ 实现 POST 创建用户 + 上传头像
6. ✅ 加中间件鉴权
7. ✅ 剩下 PUT / DELETE

---

## 💡 提示：token 鉴权中间件模板

```go
func authMiddleware(c *gin.Context) {
    token := c.GetHeader("Authorization")
    if token != "Bearer mock-token-123" {
        c.AbortWithStatusJSON(401, gin.H{"error": "未登录"})
        return
    }
    // 可以在这里 set user 信息到 Context
    c.Set("user_id", "1") // 模拟登录用户
    c.Next()
}
```

---

## 🎁 Bonus：进阶挑战（做完基础后可选）

- 支持“只能修改自己”的逻辑
- 上传头像时限制格式（jpg/png）和大小（<5MB）
- 返回统一响应格式：`{ "code": 0, "msg": "ok", "data": {} }`
- 使用 `uuid` 生成用户 ID
- 加日志中间件

---

## ✅ 下一步

你可以开始写了！  
写完某个接口，可以发给我看，我帮你 review。  
或者你卡住了，告诉我哪一步，我给你提示或代码片段。

> 想要我先给你一个 **骨架代码模板**（带路由和结构体）吗？随时说！

这个项目练完，你就能自信地说：**“我会用 Gin 做一个完整的后端服务了！”** 💪