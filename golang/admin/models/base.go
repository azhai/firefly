package models

import (
	"time"

	"github.com/astro-bug/gondor/webapi/config"
	"github.com/astro-bug/gondor/webapi/config/dialect"
	"github.com/azhai/gozzo-db/session"
	"github.com/azhai/gozzo-utils/redisw"
	"github.com/k0kubun/pp"
	"xorm.io/xorm"
)

func InitCache(cfg *config.Settings, name string, verbose bool) (*session.SessionRegistry, error) {
	var sessreg *session.SessionRegistry
	c, ok := cfg.Connections[name]
	if !ok || c.DriverName != "redis" {
		return nil, nil
	}
	if verbose {
		d := dialect.GetDialectByName(c.DriverName)
		pp.Println(d.Name(), d.GetDSN(c.Params))
		pp.Println(d.(*dialect.Redis).Options)
	}
	sessreg = session.NewRegistry(redisw.ConnParams(c.Params))
	return sessreg, nil
}

func InitConn(cfg *config.Settings, name string, verbose bool) (*xorm.Engine, error) {
	var drv, dsn string
	if c, ok := cfg.Connections[name]; ok {
		d := dialect.GetDialectByName(c.DriverName)
		if d != nil {
			drv, dsn = d.Name(), d.GetDSN(c.Params)
		}
	}
	if drv == "" || dsn == "" {
		return nil, nil
	} else if verbose {
		pp.Println(drv, dsn)
	}
	engine, err := xorm.NewEngine(drv, dsn)
	if err == nil {
		engine.ShowSQL(verbose)
	}
	return engine, err
}

/**
 * 过滤查询
 * 使用方法 query = query.Scopes(filters ...FilterFunc)
 */
type FilterFunc = func(query *xorm.Session) *xorm.Session

/**
 * 数据表名
 */
type ITableName interface {
	TableName() string
}

/**
 * 数据表注释
 */
type ITableComment interface {
	TableComment() string
}

/**
 * 带自增主键的基础Model
 */
type BaseModel struct {
	Id uint `json:"id" xorm:"not null pk autoincr INT(10)"`
}

func (BaseModel) TableComment() string {
	return ""
}

/**
 * 时间相关的三个典型字段
 */
type TimeModel struct {
	CreatedAt time.Time `json:"created_at" xorm:"created comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated comment('更新时间') TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" xorm:"deleted comment('删除时间') index TIMESTAMP"`
}

/**
 * 翻页查询，out参数需要传引用
 * 使用方法 total, err := Paginate(query, &rows, pageno, pagesize)
 */
func Paginate(query *xorm.Session, out interface{}, pageno, pagesize int) (int, error) {
	var (
		total int64
		err   error
	)
	offset, limit := 0, -1 // 初始值
	if pageno < 0 {
		total, err = query.Count()
		if err != nil || total <= 0 {
			return -1, err
		}
		offset = int(total) + pageno*pagesize
	}
	// 参数校正
	if pagesize >= 0 {
		limit = pagesize
		offset = (pageno - 1) * pagesize
	}
	if limit >= 0 && offset >= 0 {
		query = query.Limit(limit, offset)
	}
	if total > 0 {
		err = query.Find(out)
	} else {
		total, err = query.FindAndCount(out)
	}
	return int(total), err
}
