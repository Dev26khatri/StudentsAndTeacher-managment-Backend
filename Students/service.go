package students

//Services takes only bussiness logics and validations
import (
	dto "GOGIN/DTO"
	users "GOGIN/Users"
	"errors"
	"log"

	"gorm.io/gorm"
)

// type StudentService interface {
// 	Create(student Student) (*Student, error)
// 	GetAll() ([]Student, error)
// 	GetByID(id int) (Student, error)
// 	Update(student dto.UpdateStudentRequest, id int) error
// 	Delete(id int) error
// }

type Service struct {
	db         *gorm.DB
	repository Repository
	userRepo   users.Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(student Student) (*Student, error) {
	if student.Name == "" {
		return nil, errors.New("name is required")
	}
	// if student.Email == "" {
	// 	return nil, errors.New("email is required")
	// }
	if student.Age < 18 {
		return nil, errors.New("you are not elegible. Student age must be at 18 ")
	}
	result, err := s.repository.Create(student)
	return result, err
}

func (s *Service) GetAll() ([]Student, error) {

	students, err := s.repository.GetAll()
	log.Printf("From service %v", students)
	return students, err
}

func (s *Service) GetByID(id int) (Student, error) {
	if id < 0 {
		return Student{}, errors.New("Invalid student Id")
	}
	return s.repository.GetById(id)
}
func (s *Service) Update(student dto.UpdateStudentRequest, id int) error {
	if id <= 0 {
		return errors.New("Invalid student Id")
	}
	if student.Name == "" {
		return errors.New("Name is required")
	}
	if student.Email == "" {
		return errors.New("email is required")
	}
	if student.Age < 18 {
		return errors.New("Student age must be 18")
	}
	return s.repository.Update(student, id)
}

func (s *Service) Delete(id int) error {
	if id == 0 {
		return errors.New("Invalid I")
	}
	return s.repository.Delete(id)
}

func (s *Service) CreateStudent(student Student, user users.User) error {

	return s.db.Transaction(func(tx *gorm.DB) error {

		if err := s.userRepo.CreateWithTx(tx, &user); err != nil {
			return err
		}
		student.UserID = user.ID
		if err := s.repository.CreateWithTx(tx, &student); err != nil {
			return err
		}

		return nil

	})

}
