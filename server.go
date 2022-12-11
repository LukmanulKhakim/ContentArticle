package main

import (
	"content/config"
	dContent "content/feature/article/delivery"
	rContent "content/feature/article/repository"
	sContent "content/feature/article/services"
	dUser "content/feature/user/delivery"
	rUser "content/feature/user/repository"
	sUser "content/feature/user/services"
	"content/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)

	mdlUser := rUser.New(db)
	serUser := sUser.New(mdlUser)
	dUser.New(e, serUser)

	mdlContent := rContent.New(db)
	serContent := sContent.New(mdlContent)
	dContent.New(e, serContent)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":8000"))
}
