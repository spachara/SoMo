package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	//"time"

	"somo/service/cmd/api/handlers"
	"somo/service/internal/database"

	"github.com/gin-gonic/gin"
)

// type Room struct {
//   ID     string  `json:"id"`
//   Name  string  `json:"name"`
//   Status string  `json:"statuse"`
//   Price  float64 `json:"price"`
//   CreateOn  float64 `json:"createOn"`
//   ModifyOn time `json:"modifyOn"`
// }

func main() {

	var port, dbhost, dbschema, dbusername, dbpassword, disableTLS string
	var dbport int
	flag.StringVar(&port, "port", "3000", "port for open service")
	flag.StringVar(&dbhost, "dbhost", "localhost", "database host name")
	flag.IntVar(&dbport, "dbport", 5432, "database port")
	flag.StringVar(&dbschema, "dbschema", "first", "database schema name")
	flag.StringVar(&dbusername, "dbusername", "postgres", "database user name")
	flag.StringVar(&dbpassword, "dbpassword", "1234", "database password")
	flag.StringVar(&disableTLS, "disableTLS", "Y", "database disableTLS[Y/n]")
	flag.Parse()

	var databaseTSL bool
	if disableTLS == "n" {
		databaseTSL = false
	} else {
		databaseTSL = true
	}
	dbConfig := database.Config{
		User:       dbusername,
		Password:   dbpassword,
		Host:       dbhost,
		Port:       dbport,
		Name:       dbschema,
		DisableTLS: databaseTSL,
	}

	db, err := database.Open(dbConfig)
	dbp, errp := database.OpenPostgreSQL()
	if err != nil {
		log.Fatal("connecting database fail", err)
	}
	if db == nil {
		log.Fatal("database fail")
	}
	if errp != nil {
		log.Fatal("connecting gorm database fail", errp)
	}
	if dbp == nil {
		log.Fatal("database gorm fail")
	}
	r := gin.Default()

	handlers.API(db, r)
	handlers.SomoAPI(dbp, r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Default",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	log.Fatal(r.Run(fmt.Sprintf(":%s", port)))
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
