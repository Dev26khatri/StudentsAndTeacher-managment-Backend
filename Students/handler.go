package students

import (
	dto "GOGIN/DTO"
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Handler responsibilites Receive HTTP Request
//Read JSON Body , Read URL Parameters, Read Query Parameters ,Call Service, Return JSON Response

type Handler struct {
	service *Service
}

// type pagination struct{
// 	page int ``
// }

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

type queryRequest struct {
	Id int `uri:"studentId"`
}

func (h *Handler) CreateStudent(c *gin.Context) {
	var req dto.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	student := Student{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}
	result, err := h.service.Create(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Status": "Student Created",
		"id":     result.ID,
	})
}
func (h *Handler) GetAllStudents(c *gin.Context) {

	students, err := h.service.GetAll()
	log.Printf("From Handler %v", students)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": students,
	})
}
func (h *Handler) GetStudentByID(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))
	var query queryRequest
	log.Printf("ID: %v", query.Id)
	if err := c.ShouldBindUri(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Id",
		})
		return
	}
	student, err := h.service.GetByID(query.Id)
	log.Printf("Data : %v", student)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"errors": "student not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}
func (h *Handler) UpdateStudentByID(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))
	var query queryRequest

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Invalid student id",
	// 	})
	// 	return
	// }

	if err := c.ShouldBindUri(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Id",
		})
		return
	}
	var student dto.UpdateStudentRequest

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.Update(student, query.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			c.JSON(http.StatusNotFound, gin.H{
				"error": "Student Not Found",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Student Updated Successfully",
	})
}
func (h *Handler) DeleteStudentByID(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Invailid student Id",
	// 	})
	// 	return
	// }

	var query queryRequest

	if err := c.ShouldBindUri(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Id",
		})
		return
	}

	err := h.service.Delete(query.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Student Not found",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Student Deleted Successfully",
	})
}
