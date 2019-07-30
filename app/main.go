package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	localConnString       = "coffeeGo:qwe123@tcp(localhost:3307)/f1db?charset=utf8&parseTime=True&loc=Local"
	dockerMySQLConnString = "coffeeGo:qwe123@tcp(mysqlContainer:3306)/f1db?charset=utf8&parseTime=True&loc=Local&allowCleartextPasswords=true"
)

const (
	//mysqlF1 container creds
	dbContainer = "mysqlF1"
	dbPort      = "3306"
	dbUser      = "f1DBUser"
	dbPass      = "qwe123"
)

func main() {
	runtime.GOMAXPROCS(4)

	fmt.Println("Started")

	// Connect to DB
	connectionString := dockerMySQLConnString
	db, err := connectToDB(connectionString)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to DB %v", err))
	}
	defer db.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/driver/:name", func(c *gin.Context) {
		//get drivers
		driverCode := c.Param("name")
		msg := fmt.Sprintf("Requesting driver: %s", driverCode)
		fmt.Println(msg)
		// c.String(http.StatusOK, msg)
		driver := *fetchSingleDriver(db, driverCode)
		c.JSON(http.StatusOK, driver)
	})

	r.GET("/drivers", func(c *gin.Context) {
		//get drivers
		log.Println("Whats up fresh log message")
		drivers := fetchAllDrivers(db)
		c.JSON(http.StatusOK, drivers)
	})

	r.GET("/races/*year", func(c *gin.Context) {
		log.Println("Getting races")
		yearParam := c.Param("year")
		yearParam = yearParam[1:]
		year, err := strconv.Atoi(yearParam)
		if err != nil || year <= 0 {
			year = 2018
		}

		race := Race{Year: year}
		races := []Race{}
		db.Where(&race).Find(&races)
		c.JSON(http.StatusOK, races)
	})
	// r.Run(":1111") // listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

func connectToDB(mysqlConnStringContainer string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", mysqlConnStringContainer)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	return db, nil
}

type driver struct {
	Code        string    `json:"code"`
	Dob         time.Time `json:"dob"`
	Forename    string    `json:"forename"`
	Surname     string    `json:"surname"`
	Nationality string    `json:"nationality"`
	Number      int       `json:"number"`
	URL         string    `json:"url"`
}

func fetchSingleDriver(db *gorm.DB, code string) *driver {
	driver := &driver{Code: code}
	db.Where(&driver).First(&driver)
	return driver
}

func fetchAllDrivers(db *gorm.DB) []driver {
	drivers := []driver{}
	db.Find(&drivers)
	return drivers
}

/*
CREATE TABLE `races` (
  `raceId` int(11) NOT NULL AUTO_INCREMENT,
  `year` int(11) NOT NULL DEFAULT '0',
  `round` int(11) NOT NULL DEFAULT '0',
  `circuitId` int(11) NOT NULL DEFAULT '0',
  `name` varchar(255) NOT NULL DEFAULT '',
  `date` date NOT NULL DEFAULT '0000-00-00',
  `time` time DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`raceId`),
  UNIQUE KEY `url` (`url`)
) ENGINE=MyISAM AUTO_INCREMENT=1010 DEFAULT CHARSET=utf8;

*/
//Race struct to hold race rows
type Race struct {
	RaceID int       `json:"raceId"`
	Year   int       `json:"year"`
	Round  int       `json:"round"`
	Name   string    `json:"name"`
	URL    string    `json:"url"`
	Date   time.Time `json:"date"`
	Time   string    `json:"time"`
	// Date   time.Date `json:"date"`
}
