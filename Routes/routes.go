package routes

import (
	students "GOGIN/Students"
	users "GOGIN/Users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	studentHandler *students.Handler,
	userHanlder *users.Handler,

) {
	RegisterStudentRoutes(router, studentHandler, userHanlder)
}

func RegisterStudentRoutes(
	router *gin.Engine,
	studenthandler *students.Handler,
	userhandler *users.Handler,

) {
	student := router.Group("/students")
	{
		student.POST("/", studenthandler.CreateStudent)
		student.GET("/", studenthandler.GetAllStudents)
		student.GET("/:studentId", studenthandler.GetStudentByID)
		student.PUT("/:studentId", studenthandler.UpdateStudentByID)
		student.DELETE("/:studentId", studenthandler.DeleteStudentByID)
	}
	user := router.Group("/users")
	{
		user.POST("/register", userhandler.Register)
		user.POST("/login", userhandler.Login)
	}
}
