package utils

import (
	"net/http"
	"sort"
	"strings"

	"gitee.com/azhai/fiber-u8l/v2"
)

// 用户分类
const (
	USER_ANONYMOUS = iota // 匿名用户（未登录/未注册）
	USER_FORBIDDEN        // 封禁用户（有违规被封号）
	USER_LIMITED          // 受限用户（未过审或被降级）
	USER_REGULAR          // 正常用户（正式会员）
	USER_SUPER            // 超级用户（后台管理权限）
)

// 字符串比较方式
const (
	CMP_STRING_OMIT             = iota // 不比较
	CMP_STRING_CONTAINS                // 包含
	CMP_STRING_STARTSWITH              // 打头
	CMP_STRING_ENDSWITH                // 结尾
	CMP_STRING_IGNORE_SPACES           // 忽略空格
	CMP_STRING_CASE_INSENSITIVE        // 不分大小写
	CMP_STRING_EQUAL                   // 相等
)

func Abort(ctx *fiber.Ctx, code int) {
	ctx.SendStatus(code)
}

func AbortJSON(ctx *fiber.Ctx, code int, body interface{}) {
	ctx.Status(code).JSON(body)
}

// 拒绝访问
func AccessDenied(ctx *fiber.Ctx, msg string) {
	code := http.StatusForbidden
	AbortJSON(ctx, code, fiber.Map{"code": code, "data": msg})
}

func QueryDefault(ctx *fiber.Ctx, key, val string) (value string) {
	if value = ctx.Query(key); value == "" {
		value = val
	}
	return
}

// 删除空格
func RemoveSpaces(s string) string {
	return strings.Join(strings.Fields(s), "")
}

// 是否在字符串列表中
func InStringList(x string, lst []string, cmp int) bool {
	size := len(lst)
	if size == 0 {
		return false
	}
	if !sort.StringsAreSorted(lst) {
		sort.Strings(lst)
	}
	i := sort.Search(size, func(i int) bool { return lst[i] >= x })
	if i >= size {
		return false
	}

	// 比较是否相符
	switch cmp {
	default:
		return false
	case CMP_STRING_OMIT:
		return true
	case CMP_STRING_CONTAINS:
		return strings.Contains(x, lst[i])
	case CMP_STRING_STARTSWITH:
		return strings.HasPrefix(x, lst[i])
	case CMP_STRING_ENDSWITH:
		return strings.HasSuffix(x, lst[i])
	case CMP_STRING_IGNORE_SPACES:
		xx, yy := RemoveSpaces(x), RemoveSpaces(lst[i])
		return strings.EqualFold(xx, yy)
	case CMP_STRING_CASE_INSENSITIVE:
		return strings.EqualFold(x, lst[i])
	case CMP_STRING_EQUAL:
		return x == lst[i]
	}
}

// 用户分类，无法区分内部用户和普通用户
func GetUserType(ctx *fiber.Ctx) (int, error) {
	return USER_REGULAR, nil
}

// 获取正常用户权限可访问的网址
func GetRegularPermissionUrls(ctx *fiber.Ctx) (urls []string) {
	return
}

// 获取超级用户权限可访问的网址，不再检查正常用户权限
func GetSuperPermissionUrls(ctx *fiber.Ctx) (urls []string) {
	return
}

// 是否静态资源网址
func IsStaticResourceUrl(url string) bool {
	return false
}

// 获取可公开访问的网址
func GetAnonymousOpenUrls() (urls []string) {
	return
}

// 获取受限用户黑名单中的的网址，与白名单二选一
func GetLimitedBlackListUrls() (urls []string) {
	return
}

// 获取受限用户白名单中的的网址，不再检查正常用户权限，与黑名单二选一
func GetLimitedWhiteListUrls() (urls []string) {
	return
}
