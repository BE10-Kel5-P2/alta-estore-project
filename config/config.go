package config

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type Appconfig struct {
	Username  string
	Password  string
	Address   string
	Port      int
	DBName    string
	Keys3     string
	Secrets3  string
	Regions3  string
	Midserver string
}

var (
	lock      = &sync.Mutex{}
	appconfig *Appconfig
)

func Getconfig() *Appconfig {
	lock.Lock()
	defer lock.Unlock()

	if appconfig == nil {
		appconfig = initConfig()
	}

	return appconfig
}

func initConfig() *Appconfig {
	var defaultconfig Appconfig

	// err := godotenv.Load("local.env")
	// if err != nil {
	// 	log.Println("cant load env file")
	// 	return nil
	// }

	defaultconfig.Username = os.Getenv("Username")
	defaultconfig.Password = os.Getenv("Password")
	defaultconfig.Address = os.Getenv("Address")

	port, err := strconv.Atoi(os.Getenv("Port"))
	if err != nil {
		log.Println("Cant convert port")
		return nil
	}

	defaultconfig.Port = port
	defaultconfig.DBName = os.Getenv("DBName")
	defaultconfig.Keys3 = os.Getenv("Keys3")
	defaultconfig.Secrets3 = os.Getenv("Secrets3")
	defaultconfig.Regions3 = os.Getenv("Regions3")
	defaultconfig.Midserver = os.Getenv("Midserver")
	return &defaultconfig
}
