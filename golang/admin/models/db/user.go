package db

import (
	"encoding/hex"
	"time"

	"github.com/azhai/gozzo-utils/cryptogy"
	"github.com/muyo/sno"
)

func NewUser(username, realname string) *User {
	user := &User{Username: username, Realname: realname}
	user.Uid = sno.NewWithTime('U', user.CreatedAt).String()
	user.CreatedAt = time.Now()
	return user
}

// 8位salt值，用$符号分隔开
var saltPasswd = cryptogy.NewSaltPassword(8, "$")

// 设置密码
func (m *User) SetPassword(password string) *User {
	m.Password = saltPasswd.CreatePassword(password)
	return m
}

// 校验密码
func (m User) VerifyPassword(password string) bool {
	return saltPasswd.VerifyPassword(password, m.Password)
}

// 登录
func UserSignin(username, password string) (*User, string, error) {
	user, token := new(User), ""
	has, err := engine.Where("username = ?", username).Get(user)
	if has && err == nil && user.VerifyPassword(password) {
		ticket := sno.New('T').Bytes()[:8]
		tailno := cryptogy.RandSalt(2)
		token = hex.EncodeToString(append(ticket, tailno...)) // 生成token
	}
	return user, token, err
}
