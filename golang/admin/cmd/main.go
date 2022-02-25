package main

import (
	"flag"
	"fmt"

	"gitee.com/azhai/fiber-u8l/v2"
	"gitee.com/azhai/fiber-u8l/v2/middleware/compress"
	"gitee.com/azhai/fiber-u8l/v2/middleware/cors"
	"github.com/astro-bug/gondor/webapi"
	"github.com/astro-bug/gondor/webapi/config"
	"github.com/astro-bug/gondor/webapi/handlers"
	"github.com/astro-bug/gondor/webapi/models/db"
	"github.com/astro-bug/gondor/webapi/services"
)

var (
	configFile string
	port       uint // 运行端口
	verbose    bool // 详细输出
)

func init() {
	flag.UintVar(&port, "p", 8000, "运行端口")
	flag.StringVar(&configFile, "c", "settings.yml", "配置文件路径")
	flag.BoolVar(&verbose, "v", false, "输出详细信息")
	flag.Parse()

	cfg, err := config.ReadSettings(configFile)
	if err != nil {
		panic(err)
	}
	if verbose == false {
		verbose = cfg.Application.Debug
	}
	db.Initialize(cfg, verbose)
	services.Initialize(cfg, verbose)
}

func main() {
	app := fiber.New()
	app.Use(compress.New()).Use(cors.New())
	webapi.AddRoutes(app.Group("/api", handlers.RoleAuth))
	app.Listen(fmt.Sprintf(":%d", port))
}
