package db

import (
	"github.com/astro-bug/gondor/webapi/models"
	"xorm.io/xorm"
)

// 添加一个菜单
func (m *Menu) AddTo(parent *Menu) (err error) {
	var parentNode *models.NestedModel
	if parent != nil {
		parentNode = parent.NestedModel
	}
	if m.NestedModel == nil {
		m.NestedModel = new(models.NestedModel)
	}
	query := Table(m.TableName())
	err = m.NestedModel.AddToParent(parentNode, query)
	if err == nil {
		_, err = engine.InsertOne(m)
	}
	return
}

// 添加菜单
func AddMenu(path, title string, icon string, parent *Menu) (menu *Menu, err error) {
	menu = &Menu{Path: path, Title: title, Icon: icon}
	err = menu.AddTo(parent)
	return
}

// 写入必须的初始化数据
func FillRequiredData(drv string, query *xorm.Session) *xorm.Session {
	// 菜单
	menu := new(Menu)
	if count, _ := Table(menu).Count(); count == 0 {
		_, _ = AddMenu("/dashboard", "面板", "dashboard", nil)
		menu, _ = AddMenu("/permission", "权限", "lock", nil)
		_, _ = AddMenu("role", "角色权限", "", menu)
		menu, _ = AddMenu("/table", "Table", "table", nil)
		_, _ = AddMenu("complex-table", "复杂Table", "", menu)
		_, _ = AddMenu("inline-edit-table", "内联编辑", "", menu)
		menu, _ = AddMenu("/excel", "Excel", "excel", nil)
		_, _ = AddMenu("export-selected-excel", "选择导出", "", menu)
		_, _ = AddMenu("upload-excel", "上传Excel", "", menu)
		menu, _ = AddMenu("/theme/index", "主题", "theme", nil)
		menu, _ = AddMenu("/error/404", "404错误", "404", nil)
		menu, _ = AddMenu("https://cn.vuejs.org/", "外部链接", "link", nil)
	}
	// 权限
	if count, _ := Table(new(Access)).Count(); count == 0 {
		AddAccess("superuser", "menu", ACCESS_ALL, "*") // 超管可以访问所有菜单
		// 基本用户
		AddAccess("visitor", "menu", ACCESS_VIEW, "/dashboard")
		AddAccess("visitor", "menu", ACCESS_VIEW, "/error/404")
	}
	return query
}
