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
	handler *students.Handler,
	userhandler *users.Handler,

) {
	student := router.Group("/students")
	{
		student.POST("/", handler.CreateStudent)
		student.GET("/", handler.GetAllStudents)
		student.GET("/:studentId", handler.GetStudentByID)
		student.PUT("/:studentId", handler.UpdateStudentByID)
		student.DELETE("/:studentId", handler.DeleteStudentByID)
	}
	user := router.Group("/users")
	{
		user.POST("/register")
		user.POST("/login")
	}
}
