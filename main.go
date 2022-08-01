package main

import (
	"fmt"

	"altaproject2/config"
	"altaproject2/factory"
	awss3 "altaproject2/utils/aws"
	"altaproject2/utils/database/mysql"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.Getconfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateDB(db)

	e := echo.New()
	awsConn := awss3.InitS3(cfg.Keys3, cfg.Secrets3, cfg.Regions3)
	factory.InitFactory(e, db, awsConn)

	fmt.Println("==== STARTING PROGRAM ====")
	address := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(address))
}
