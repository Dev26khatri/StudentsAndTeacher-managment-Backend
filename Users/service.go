package users

import (
	dto "GOGIN/DTO"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Register(req dto.RegisterUserRequest) error {

	// check wather the ema	il already exists
	existingUser, err := s.repository.GetByEmail(req.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already exist")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	//Hash the Entered password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	//Convert DTO to model

	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	//save user
	_, err = s.repository.CreateUser(user)

	if err != nil {
		return err
	}
	return nil

}
func (s *Service) Login(req dto.LoginUserRequest) (*User, error) {
	//step 1 Find User by email
	user, err := s.repository.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("Invalid email or password")
	}

	// step 2 compare the enterd password with stored has password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, errors.New("Invalid email or password")
	}

	//Step 3 login succesffuly
	return user, nil
}
