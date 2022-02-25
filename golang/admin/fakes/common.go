package fakes

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bxcodec/faker/v3"
)

var (
	FakeUsers, Articles []string
	ArticleTotal        = 100
)

func init() {
	for i := 0; i < ArticleTotal; i++ {
		Articles = append(Articles, GenArticle(i+1))
		FakeUsers = append(FakeUsers, faker.FirstNameMale())
		FakeUsers = append(FakeUsers, faker.FirstNameFemale())
	}
}

// 将多个空白或换行合并成一个空格
func ReduceBlanks(lines string) string {
	return strings.Join(strings.Fields(lines), " ")
}

// 随机整数
func RandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// 随机浮点数
func RandFloat(min, max float64, inclRight bool) float64 {
	offset := (max - min) * rand.Float64()
	if inclRight {
		return max - offset
	} else {
		return min + offset
	}
}

// 随机元素
func RandItem(items []string) string {
	return items[rand.Intn(len(items))]
}

// 随机时间
func RandTime(secs int64, afterNow bool) time.Time {
	offset := rand.Int63n(secs)
	duration := time.Second * time.Duration(offset)
	if afterNow {
		return time.Now().Add(duration)
	} else {
		return time.Now().Add(0 - duration)
	}
}
