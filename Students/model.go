package students

type Student struct {
	ID    int    `gorm:"id"`
	Name  string `gorm:"name"`
	Email string `gorm:"email"`
	Age   int    `gorm:"age"`
}

// type CreateStudentRequest struct {
// 	Name  string `json:"name" validate:"required,min=3,max=50"`
// 	Email string `json:"email" validate:"required,email"`
// 	Age   int    `json:"age" validate:"required,gte=18 lte=100"`
// }
