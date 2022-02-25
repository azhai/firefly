package webapi

import (
	"gitee.com/azhai/fiber-u8l/v2"
	"github.com/astro-bug/gondor/webapi/handlers"
)

func AddRoutes(r fiber.Router) {
	// 用户登录
	r.Post("/user/login", handlers.UserLoginHandler)
	// 用户退出
	r.Post("/user/logout", handlers.UserLogoutHandler)
	// 用户资料
	r.Get("/user/info", handlers.UserInfoHandler)

	// 查找用户名
	r.Get("/search/user", handlers.SearchUserHandler)
	// 订单列表
	r.Get("/transaction/list", handlers.OrderListHandler)

	// 完整菜单
	r.Get("/routes", handlers.AllMenuHandler)
	// 角色列表
	r.Get("/roles", handlers.RoleListHandler)
	// 添加角色
	r.Post("/role", handlers.RoleAddHandler)
	// 修改角色
	r.Put("/role/:name", handlers.RoleModHandler)
	// 删除角色
	r.Delete("/role/:name", handlers.RoleModHandler)

	// 文章列表
	r.Get("/article/list", handlers.ArticleListHandler)
	// 文章详情
	r.Get("/article/detail", handlers.ArticleDetailHandler)
	// 文章阅读量
	r.Get("/article/pv", handlers.ArticleReadHandler)
	// 添加文章
	r.Post("/article/create", handlers.ArticleModHandler)
	// 修改文章
	r.Post("/article/update", handlers.ArticleModHandler)
}
