package handlers

import (
	"strconv"
	"strings"

	"gitee.com/azhai/fiber-u8l/v2"
	"github.com/astro-bug/gondor/webapi/fakes"
	"github.com/astro-bug/gondor/webapi/utils"
)

// 文章列表
func ArticleListHandler(ctx *fiber.Ctx) (err error) {
	var (
		pageno, pagesize int
		sort             string
		arts             []string
		// importance int
		// title, nation string
	)
	pageno, _ = strconv.Atoi(utils.QueryDefault(ctx, "page", "1"))
	pagesize, _ = strconv.Atoi(utils.QueryDefault(ctx, "limit", "20"))
	if pagesize < 0 {
		pagesize = 100
	}
	if sort = ctx.Query("sort"); sort == "-id" {
		offset := 99
		if pageno > 0 {
			offset -= (pageno - 1) * pagesize
		}
		if pagesize > offset+1 {
			pagesize = offset + 1
		}
		for i := 0; i < pagesize; i++ {
			arts = append(arts, fakes.Articles[offset-i])
		}
	} else {
		offset := 0
		if pageno > 0 {
			offset += (pageno - 1) * pagesize
		}
		if pagesize > fakes.ArticleTotal-offset {
			pagesize = fakes.ArticleTotal - offset
		}
		for i := 0; i < pagesize; i++ {
			arts = append(arts, fakes.Articles[offset+i])
		}
	}
	err = ctx.Type("json").Send([]byte(`{"code":200, "total":` +
		strconv.Itoa(fakes.ArticleTotal) + `, "data":[` + strings.Join(arts, ", ") + `]}`))
	return
}

// 文章详情
func ArticleDetailHandler(ctx *fiber.Ctx) (err error) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	result := fiber.Map{
		"code": 200,
		"data": fakes.Articles[id-1],
	}
	err = ctx.JSON(result)
	return
}

// 文章阅读量
func ArticleReadHandler(ctx *fiber.Ctx) (err error) {
	result := fakes.ReduceBlanks(`{"code":200, "data":` + fakes.PageViewData() + `}`)
	err = ctx.Type("json").Send([]byte(result))
	return
}

// 添加修改文章
func ArticleModHandler(ctx *fiber.Ctx) (err error) {
	result := fiber.Map{
		"code": 200,
		"data": "success",
	}
	err = ctx.JSON(result)
	return
}
