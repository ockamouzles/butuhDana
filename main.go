package main

import (
	"log"
	"projekStartup/handler"
	"projekStartup/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "idris:idris@tcp(127.0.0.1:3306)/projekStartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Munir At Tarbiy"
	// userInput.Email = "Hilang@amsterdan2004.com"
	// userInput.Occupation = "Gigolo"
	// userInput.Password = "password"

	// userService.Register(userInput)

	// fmt.Println("Koneksi aman")

	// var users []user.User
	// db.Find(&users)

	// for _, data := range users {
	// 	fmt.Println(data.Name)
	// 	fmt.Println(data.Email)
	// 	fmt.Println()
	// }
	// route := gin.Default()
	// route.GET("/handler", handler)
	// route.Run()
}

// func handler(c *gin.Context) {
// 	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
// 	dsn := "idris:idris@tcp(127.0.0.1:3306)/projekStartup?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	var users []user.User

// 	db.Find(&users)
// 	c.JSON(http.StatusOK, users)
// }
