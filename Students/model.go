package students

import (
	users "GOGIN/Users"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model

	// ID    int    `gorm:"id"`
	// Email string `gorm:"email"` Not using becuase we are using in the User model and students and User connected togather
	Name string `gorm:"name"`
	Age  int    `gorm:"age"`

	//Foreign key
	UserID uint `gorm:"not null;uniqueIndex"`

	//Association
	User users.User `gorm:"foreignKey:UserID;references:ID"`
}

// type CreateStudentRequest struct {
// 	Name  string `json:"name" validate:"required,min=3,max=50"`
// 	Email string `json:"email" validate:"required,email"`
// 	Age   int    `json:"age" validate:"required,gte=18 lte=100"`
// }
