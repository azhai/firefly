package handlers

import (
	"net/http"

	"gitee.com/azhai/fiber-u8l/v2"
	"github.com/astro-bug/gondor/webapi/models/db"
	"github.com/astro-bug/gondor/webapi/utils"
	"github.com/azhai/gozzo-db/session"
)

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUserInfo(user *db.User) map[string]interface{} {
	info := map[string]interface{}{
		"username": user.Username,
	}
	if user.Realname != "" {
		info["realname"] = user.Realname
	}
	if user.Mobile != "" {
		info["mobile"] = user.Mobile
	}
	if user.Avatar != "" {
		info["avatar"] = user.Avatar
	}
	if user.Introduction != "" {
		info["introduction"] = user.Introduction
	}
	return info
}

// 用户登录
func UserLoginHandler(ctx *fiber.Ctx) (err error) {
	// 获取参数
	var data UserData
	if err = ctx.BodyParser(&data); err != nil {
		err = ctx.JSON(fiber.Map{
			"code":    510,
			"message": err.Error(),
		})
	}
	// 查询数据
	user, token, err := db.UserSignin(data.Username, data.Password)
	if err != nil || token == "" {
		err = ctx.JSON(fiber.Map{
			"code":    510,
			"message": "失败，密码不正确！",
		})
		return
	}
	// 写入Session
	roles := db.GetUserRoles(user.Uid)
	sess := db.Session(token)
	sess.BindRoles(user.Uid, roles, true)
	sess.SaveMap(GetUserInfo(user), false)
	err = ctx.JSON(fiber.Map{
		"code": 200,
		"data": fiber.Map{
			"token": token,
		},
	})
	return
}

// 用户退出
func UserLogoutHandler(ctx *fiber.Ctx) (err error) {
	token := ctx.Cookies("access_token")
	if token == "" {
		utils.Abort(ctx, http.StatusUnauthorized)
		return
	}
	db.Registry().DelSession(token)
	result := fiber.Map{
		"code": 200,
		"data": "成功退出",
	}
	err = ctx.JSON(result)
	return
}

// 用户资料
func UserInfoHandler(ctx *fiber.Ctx) (err error) {
	token := ctx.Query("token")
	sess := db.Session(token)
	if uid, err := sess.GetString("uid"); err != nil || uid == "" {
		result := fiber.Map{
			"code":    508,
			"message": "失败，用户不存在！",
		}
		err = ctx.JSON(result)
	} else {
		data, _ := sess.GetAllString()
		result := fiber.Map{
			"code": 200,
			"data": fiber.Map{
				"roles":        session.SessListSplit(data["roles"]),
				"name":         data["name"],
				"avatar":       data["avatar"],
				"introduction": data["introduction"],
			},
		}
		err = ctx.JSON(result)
	}
	return
}
