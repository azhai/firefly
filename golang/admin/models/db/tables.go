package db

import (
	"time"

	"github.com/astro-bug/gondor/webapi/models"
)

type Access struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	RoleName     string    `json:"role_name" xorm:"not null default '' comment('角色名') index VARCHAR(50)"`
	ResourceType string    `json:"resource_type" xorm:"not null default '' comment('资源类型') VARCHAR(50)"`
	ResourceArgs string    `json:"resource_args" xorm:"comment('资源参数') VARCHAR(255)"`
	PermCode     int       `json:"perm_code" xorm:"not null default 0 comment('权限码') SMALLINT(5)"`
	Actions      string    `json:"actions" xorm:"not null default '' comment('允许的操作') VARCHAR(50)"`
	GrantedAt    time.Time `json:"granted_at" xorm:"comment('授权时间') TIMESTAMP"`
	RevokedAt    time.Time `json:"revoked_at" xorm:"comment('撤销时间') index TIMESTAMP"`
}

func (Access) TableName() string {
	return "t_access"
}

type CronDaily struct {
	Id       int    `json:"id" xorm:"not null pk autoincr comment('主键') INT(10)"`
	TaskId   int    `json:"task_id" xorm:"not null default 0 comment('任务ID') index INT(10)"`
	IsActive int    `json:"is_active" xorm:"not null default b'0' comment('有效') BIT(1)"`
	Workday  int    `json:"workday" xorm:"not null default b'0' comment('工作日') BIT(1)"`
	Weekday  int    `json:"weekday" xorm:"not null default 0 comment('周X|周Y...') TINYINT(3)"`
	RunClock string `json:"run_clock" xorm:"not null default '' comment('具体时间') index CHAR(8)"`
}

func (CronDaily) TableName() string {
	return "t_cron_daily"
}

type CronNotice struct {
	Id             int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(10)"`
	UserId         int       `json:"user_id" xorm:"default 0 comment('用户ID') index INT(10)"`
	TaskId         int       `json:"task_id" xorm:"not null default 0 comment('任务ID') index INT(10)"`
	IsActive       int       `json:"is_active" xorm:"not null default b'0' comment('有效') BIT(1)"`
	Important      int       `json:"important" xorm:"not null default 0 comment('重要程度') TINYINT(3)"`
	Message        string    `json:"message" xorm:"comment('消息内容') TEXT"`
	ReadTime       time.Time `json:"read_time" xorm:"comment('阅读时间') index DATETIME"`
	DelayStartTime time.Time `json:"delay_start_time" xorm:"comment('推迟开始时间') index DATETIME"`
	StartTime      time.Time `json:"start_time" xorm:"comment('开始时间') DATETIME"`
	StopTime       time.Time `json:"stop_time" xorm:"comment('结束时间') DATETIME"`
	StartClock     string    `json:"start_clock" xorm:"comment('开始时刻') CHAR(8)"`
	StopClock      string    `json:"stop_clock" xorm:"comment('结束时刻') CHAR(8)"`
}

func (CronNotice) TableName() string {
	return "t_cron_notice"
}

type CronTask struct {
	Id         int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(10)"`
	UserId     int       `json:"user_id" xorm:"default 0 comment('用户ID') index INT(10)"`
	ReferId    int       `json:"refer_id" xorm:"not null default 0 comment('关联任务ID') index INT(10)"`
	IsActive   int       `json:"is_active" xorm:"not null default b'0' comment('有效') BIT(1)"`
	Behind     int       `json:"behind" xorm:"not null default 0 comment('相对推迟/提前多少分钟') SMALLINT(6)"`
	ActionType string    `json:"action_type" xorm:"not null default 'command' comment('动作类型') ENUM('command','function','http_get','http_post','message')"`
	CmdUrl     string    `json:"cmd_url" xorm:"not null default '' comment('指令或网址') VARCHAR(500)"`
	ArgsData   string    `json:"args_data" xorm:"comment('参数或消息体') TEXT"`
	LastTime   time.Time `json:"last_time" xorm:"comment('最后执行时间') index DATETIME"`
	LastResult string    `json:"last_result" xorm:"comment('执行结果') TEXT"`
	LastError  string    `json:"last_error" xorm:"comment('出错信息') TEXT"`
}

func (CronTask) TableName() string {
	return "t_cron_task"
}

type CronTimer struct {
	Id       int       `json:"id" xorm:"not null pk autoincr comment('主键') INT(10)"`
	TaskId   int       `json:"task_id" xorm:"not null default 0 comment('任务ID') index INT(10)"`
	IsActive int       `json:"is_active" xorm:"not null default b'0' comment('有效') BIT(1)"`
	RunDate  time.Time `json:"run_date" xorm:"comment('指定日期') index DATE"`
	RunClock string    `json:"run_clock" xorm:"not null default '' comment('具体时间') CHAR(8)"`
}

func (CronTimer) TableName() string {
	return "t_cron_timer"
}

type Group struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Gid       string    `json:"gid" xorm:"not null default '' comment('唯一ID') unique CHAR(16)"`
	Title     string    `json:"title" xorm:"not null default '' comment('名称') VARCHAR(50)"`
	Remark    string    `json:"remark" xorm:"comment('说明备注') TEXT"`
	CreatedAt time.Time `json:"created_at" xorm:"created comment('创建时间') TIMESTAMP"`
}

func (Group) TableName() string {
	return "t_group"
}

type Menu struct {
	Id                  int `json:"id" xorm:"not null pk autoincr INT(10)"`
	*models.NestedModel `xorm:"extends"`
	Path                string `json:"path" xorm:"not null default '' comment('路径') index VARCHAR(100)"`
	Title               string `json:"title" xorm:"not null default '' comment('名称') VARCHAR(50)"`
	Icon                string `json:"icon" xorm:"comment('图标') VARCHAR(30)"`
	Remark              string `json:"remark" xorm:"comment('说明备注') TEXT"`
	models.TimeModel    `xorm:"extends"`
}

func (Menu) TableName() string {
	return "t_menu"
}

type Role struct {
	Id               int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name             string `json:"name" xorm:"not null default '' comment('名称') unique VARCHAR(50)"`
	Remark           string `json:"remark" xorm:"comment('说明备注') TEXT"`
	models.TimeModel `xorm:"extends"`
}

func (Role) TableName() string {
	return "t_role"
}

type User struct {
	Id               int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Uid              string `json:"uid" xorm:"not null default '' comment('唯一ID') unique CHAR(16)"`
	Username         string `json:"username" xorm:"not null default '' comment('用户名') index VARCHAR(30)"`
	Password         string `json:"password" xorm:"not null default '' comment('密码') VARCHAR(60)"`
	Realname         string `json:"realname" xorm:"comment('昵称/称呼') VARCHAR(20)"`
	Mobile           string `json:"mobile" xorm:"comment('手机号码') index VARCHAR(20)"`
	Email            string `json:"email" xorm:"comment('电子邮箱') VARCHAR(50)"`
	PrinGid          string `json:"prin_gid" xorm:"not null default '' comment('主用户组') CHAR(16)"`
	ViceGid          string `json:"vice_gid" xorm:"comment('次用户组') CHAR(16)"`
	Avatar           string `json:"avatar" xorm:"comment('头像') VARCHAR(100)"`
	Introduction     string `json:"introduction" xorm:"comment('介绍说明') VARCHAR(500)"`
	models.TimeModel `xorm:"extends"`
}

func (User) TableName() string {
	return "t_user"
}

type UserRole struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	UserUid  string `json:"user_uid" xorm:"not null default '' comment('用户ID') index CHAR(16)"`
	RoleName string `json:"role_name" xorm:"not null default '' comment('角色名') index VARCHAR(50)"`
}

func (UserRole) TableName() string {
	return "t_user_role"
}
