package db

import (
	"github.com/astro-bug/gondor/webapi/config"
	"github.com/astro-bug/gondor/webapi/models"
	"github.com/azhai/gozzo-db/session"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	engine  *xorm.Engine
	sessreg *session.SessionRegistry
)

// 初始化、连接数据库和缓存
func Initialize(cfg *config.Settings, verbose bool) {
	var err error
	sessreg, err = models.InitCache(cfg, "cache", verbose)
	if err != nil {
		panic(err)
	}
	engine, err = models.InitConn(cfg, "default", verbose)
	if err != nil || engine == nil {
		panic(err)
	}
}

// 查询某张数据表
func Engine() *xorm.Engine {
	return engine
}

// 查询某张数据表
func Table(name interface{}) *xorm.Session {
	if engine == nil {
		return nil
	}
	return engine.Table(name)
}

// 获得当前会话管理器
func Registry() *session.SessionRegistry {
	return sessreg
}

// 获得用户会话
func Session(token string) *session.Session {
	if sessreg == nil {
		return nil
	}
	return sessreg.GetSession(token)
}
