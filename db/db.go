package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	TestRun  bool   `json:"TestRun"`
	IP       string `json:"IP"`
	User     string `json:"user"`
	Password string `json:"password"`
}

//TODO: Add a config file with the needed information in the config folder to ensure your backend won't crash

func GetConfig() Config {
	jsonFile, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Config
	json.Unmarshal(byteValue, &config)
	return config
}

var port = 3306

type _config struct {
	DbUser     string
	DbPassword string
	DbIP       string
	DbPort     int
}

type Connection interface {
	DB1() *sql.DB
	Close()
}

type Conn struct {
	db *sql.DB
}

func NewConnection(dbName string) Connection {
	var err error
	var c Conn
	var server = GetConfig()
	config := _config{
		DbUser:     server.User,
		DbPassword: server.Password,
		DbPort:     port,
		DbIP:       server.IP,
	}
	connection := server.TestRun
	if connection == false {
		connString1 := fmt.Sprint(config.DbUser, ":"+config.DbPassword, "@tcp(", config.DbIP, ":", config.DbPort, ")/", dbName, "?parseTime=true")
		c.db, err = sql.Open("mysql", connString1)
		if err != nil {
			log.Fatal("Error creating connection 1st pool: " + err.Error())
		}

		log.Println("Connected")

	} else {
		log.Println("Test Run")
	}
	return &c
}

func (c *Conn) DB1() *sql.DB {
	return c.db
}

func (c *Conn) Close() {
	c.db.Close()
}
