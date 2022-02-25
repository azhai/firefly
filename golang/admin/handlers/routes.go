package handlers

var (
	permRoute = `
	{
		"path": "/permission",
		"component": "layout/Layout",
		"redirect": "/permission/index",
		"alwaysShow": true,
		"meta": {
			"title": "权限",
			"icon": "lock",
			"roles": ["superuser", "member"]
		},
		"children": [
		{
			"path": "role",
			"component": "views/permission/role",
			"name": "RolePermission",
			"meta": {
				"title": "角色权限",
				"roles": ["superuser"]
			}
		}
	]
	}`

	constantRoutes = []string{`
{
	"path": "/auth-redirect",
	"component": "views/login/auth-redirect",
	"hidden": true
}`, `
	{
		"path": "/redirect",
		"component": "layout/Layout",
		"hidden": true,
		"children": [
		{
			"path": "/redirect/:path*",
			"component": "views/redirect/index"
		}
	]
	}`, `
	{
		"path": "/login",
		"component": "views/login/index",
		"hidden": true
	}`, `

	{
		"path": "/404",
		"component": "views/error-page/404",
		"hidden": true
	}`, `

	{
		"path": "",
		"component": "layout/Layout",
		"redirect": "dashboard",
		"children": [
		{
			"path": "dashboard",
			"component": "views/dashboard/index",
			"name": "Dashboard",
			"meta": { "title": "面板", "icon": "dashboard", "affix": true }
		}
	]
	}`}

	asyncRoutes = []string{`
	{
		"path": "/table",
		"component": "layout/Layout",
		"redirect": "/table/complex-table",
		"name": "Table",
		"meta": {
			"title": "Table",
			"icon": "table"
		},
		"children": [
		{
			"path": "complex-table",
			"component": "views/table/complex-table",
			"name": "ComplexTable",
			"meta": { "title": "复杂Table" }
		},
		{
			"path": "inline-edit-table",
			"component": "views/table/inline-edit-table",
			"name": "InlineEditTable",
			"meta": { "title": "内联编辑" }
		}
	]
	}`, `

	{
		"path": "/excel",
		"component": "layout/Layout",
		"redirect": "/excel/export-excel",
		"name": "Excel",
		"meta": {
			"title": "Excel",
			"icon": "excel"
		},
		"children": [
		{
			"path": "export-selected-excel",
			"component": "views/excel/select-excel",
			"name": "SelectExcel",
			"meta": { "title": "选择导出" }
		},
		{
			"path": "upload-excel",
			"component": "views/excel/upload-excel",
			"name": "UploadExcel",
			"meta": { "title": "上传Excel" }
		}
	]
	}`, `

	{
		"path": "/theme",
		"component": "layout/Layout",
		"redirect": "noRedirect",
		"children": [
		{
			"path": "index",
			"component": "views/theme/index",
			"name": "Theme",
			"meta": { "title": "主题", "icon": "theme" }
		}
	]
	}`, `

	{
		"path": "/error",
		"component": "layout/Layout",
		"redirect": "noRedirect",
		"name": "ErrorPages",
		"meta": {
			"title": "Error Pages",
			"icon": "404"
		},
		"children": [
		{
			"path": "404",
			"component": "views/error-page/404",
			"name": "Page404",
			"meta": { "title": "404错误", "noCache": true }
		}
	]
	}`, `

	{
		"path": "external-link",
		"component": "layout/Layout",
		"children": [
		{
			"path": "https://cn.vuejs.org/",
			"meta": { "title": "外部链接", "icon": "link" }
		}
	]
	}`, `

	{ "path": "*", "redirect": "/404", "hidden": true }`}
)
