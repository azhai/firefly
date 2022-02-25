package fakes

import (
	"fmt"
	"strings"

	"github.com/bxcodec/faker/v3"
)

const (
	baseContent = "<p>I am testing data, I am testing data.</p>" +
		"<p><img src=\\\"https://wpimg.wallstcn.com/4c69009c-0fd4-4153-b112-6cb53d1cf943\\\"></p>"
	imageUri = "https://wpimg.wallstcn.com/e4558086-631c-425c-9430-56ffb46e70b3"
)

func GenTitle() string {
	var words []string
	count := RandInt(5, 10)
	for i := 0; i < count; i++ {
		words = append(words, faker.Word())
	}
	return strings.Join(words, " ")
}

func GenArticle(id int) string {
	layout := "2006-01-02 15:04:05"
	displayTime := RandTime(25*365*86400, false)
	author, reviewer := faker.FirstName(), faker.FirstName()
	forecast := RandFloat(70, 100, true)
	importance := RandInt(1, 3)
	nation := RandItem([]string{"CN", "US", "JP", "EU"})
	status := RandItem([]string{"published", "draft", "deleted"})
	pageviews := RandInt(300, 5000)
	return ReduceBlanks(fmt.Sprintf(`{
    "id": %d, 
    "timestamp": %d, 
    "author": "%s", 
    "reviewer": "%s", 
    "title": "%s", 
    "content_short": "mock data", 
    "content": "%s", 
    "forecast": %.2f, 
    "importance": %d, 
    "type": "%s", 
    "status": "%s", 
    "display_time": "%s", 
    "comment_disabled": true, 
    "pageviews": %d, 
    "image_uri": "%s", 
    "platforms": [
        "a-platform"
    ]
}`, id, displayTime.Unix()*1000, author, reviewer, GenTitle(),
		baseContent, forecast, importance, nation, status,
		displayTime.Format(layout), pageviews, imageUri))
}

// 文章阅读量
func PageViewData() string {
	return `{
	"pvData": [
		{ "key": "PC", "pv": 1024 },
		{ "key": "mobile", "pv": 1024 },
		{ "key": "ios", "pv": 1024 },
		{ "key": "android", "pv": 1024 }
	]
}`
}
