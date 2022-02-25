package handlers

import (
	"net/http"
	"strings"

	"gitee.com/azhai/fiber-u8l/v2"
	"github.com/astro-bug/gondor/webapi/models/db"
	"github.com/astro-bug/gondor/webapi/utils"
)

// 用户鉴权
func UserAuth(ctx *fiber.Ctx) {
	currUrl := ctx.Path()
	// 1. 静态资源，直接放行
	if utils.IsStaticResourceUrl(currUrl) {
		ctx.Next()
		return
	}

	userType, err := utils.GetUserType(ctx)
	if err != nil { // 出错了
		utils.AccessDenied(ctx, err.Error())
		return
	}

	// 2. 匿名用户，如果是公开资源放行，否则失败
	if userType == utils.USER_ANONYMOUS {
		var urls = utils.GetAnonymousOpenUrls()
		if !utils.InStringList(currUrl, urls, utils.CMP_STRING_STARTSWITH) {
			utils.AccessDenied(ctx, "已注册用户可访问，请您先登录！")
		} else {
			ctx.Next()
		}
		return // 匿名用户到此为止
	}

	// 3. 受限用户，优先判断黑名单，此网址在黑名单中则失败
	if userType == utils.USER_LIMITED {
		if urls := utils.GetLimitedBlackListUrls(); len(urls) > 0 { // 二选一
			if utils.InStringList(currUrl, urls, utils.CMP_STRING_STARTSWITH) {
				utils.AccessDenied(ctx, "您的账号无权限访问，请联系客服！")
				return
			}
		} else if urls := utils.GetLimitedWhiteListUrls(); len(urls) > 0 { // 二选一
			if utils.InStringList(currUrl, urls, utils.CMP_STRING_STARTSWITH) {
				ctx.Next()
				return
			}
		}
	}

	// 4. 超级用户，如果有此网址权限则放行
	if urls := utils.GetSuperPermissionUrls(ctx); len(urls) > 0 {
		if utils.InStringList(currUrl, urls, utils.CMP_STRING_STARTSWITH) {
			ctx.Next()
			return
		}
	}

	// 5. 正常用户，如果有此网址权限则放行，内容最多，放在最后
	if urls := utils.GetRegularPermissionUrls(ctx); len(urls) > 0 {
		if utils.InStringList(currUrl, urls, utils.CMP_STRING_STARTSWITH) {
			ctx.Next()
			return
		}
	}

	// 6. 权限不明确网址，失败
	utils.AccessDenied(ctx, "找不到此网址，请核实后访问！")
}

// 用户认证与角色控制
func RoleAuth(ctx *fiber.Ctx) (err error) {
	// 匿名可访问页面
	path := ctx.Path()
	if strings.HasPrefix(path, "/api/user/log") {
		ctx.Next() // 登录(/api/user/login)或退出(/api/user/logout)
		return
	} else if strings.HasPrefix(path, "/error/") {
		ctx.Next() // 各种错误页面，例如404错误(/error/404)
		return
	}

	// 用户认证
	token := ctx.Cookies("access_token")
	if token == "" {
		utils.Abort(ctx, http.StatusUnauthorized)
		return
	}
	sess := db.Session(token)
	if sess.GetTimeout(false) < -1 {
		utils.Abort(ctx, http.StatusUnauthorized)
		return
	}

	// 鉴权
	ctx.Next()
	return
}
