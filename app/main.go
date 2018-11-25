package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	mysqlConnStringContainer = "coffeeGo:qwe123@tcp(localhost:3307)/f1db?charset=utf8&parseTime=True&loc=Local"
)

func main() {
	runtime.GOMAXPROCS(4)

	// Connect to DB
	db, err := connectToDB()
	if err != nil {
		// panic(fmt.Sprintf("Error connecting to DB %v", err))
		fmt.Printf("Error connecting to db %v", err)
	}
	defer db.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/driver", func(c *gin.Context) {
		//get drivers
		driver := *fetchSingleDriver(db)
		c.JSON(http.StatusOK, driver)
	})

	r.GET("/drivers", func(c *gin.Context) {
		//get drivers
		drivers := fetchAllDrivers(db)
		c.JSON(http.StatusOK, drivers)
	})

	r.Run() // listen and serve on 0.0.0.0:8080}
}

func connectToDB() (*gorm.DB, error) {
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

func fetchSingleDriver(db *gorm.DB) *driver {
	driver := &driver{}
	db.First(&driver)
	return driver
}

func fetchAllDrivers(db *gorm.DB) []driver {
	drivers := []driver{}
	db.Find(&drivers)
	return drivers
}
