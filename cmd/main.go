package main

import (
	routes "GOGIN/Routes"
	students "GOGIN/Students"
	users "GOGIN/Users"
	"GOGIN/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	//Auto Migration is Pending

	err = db.AutoMigrate(&students.Student{}, &users.User{})
	if err != nil {
		log.Fatal(err)
	}

	//Student DB to Handler Flow. Like Handler call -> Service -> Repo -> DB
	studentRepository := students.NewRepository(db)
	studentService := students.NewService(studentRepository)
	studentHandler := students.NewHandler(studentService)

	//User
	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	userHandler := users.NewHandler(userService)

	router := gin.Default()

	routes.RegisterRoutes(
		router,
		studentHandler,
		userHandler,
	)

	log.Println("🚀 Server Started at :8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
