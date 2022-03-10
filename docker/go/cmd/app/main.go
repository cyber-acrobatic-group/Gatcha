package main

import (
	//"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	// "github.com/golang-migrate/migrate/v4"
	// "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// db, _ := sql.Open("mysql", "docker:docker@tcp(mysql:3306)/sample-mysql")
	// driver, _ := mysql.WithInstance(db, &mysql.Config{})
	// m, _ := migrate.NewWithDatabaseInstance(
	// 	"file:/app/mysql/migration",
	// 	"mysql",
	// 	driver,
	// )
	// version, dirty, err := m.Version()
	// fmt.Println("-----version-----")
	// fmt.Println(version)
	// fmt.Println("-----dirty-----")
	// fmt.Println(dirty)
	// fmt.Println("-----error-----")
	// fmt.Println(err)
	// fmt.Println("----------")
	// if m == nil {
	// 	fmt.Println("エラー")
	// }
	// if err := m.Up().Error(); err != "" {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success!",
		})
		fmt.Println("Hello, World!!")
	})
	router.Run(":3000")

	//fmt.Println("Hello, World!!")
}

// docker-compose exec app migrate -database "mysql://docker:docker@tcp(mysql:3306)/sample-mysql" -path ./app/mysql/migration -verbose up
