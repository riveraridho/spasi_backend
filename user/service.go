package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	//IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID() (User, error)
	UpdateUser(input FormUpdateUserInput) (User, error)
	Delete(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {

	user := User{}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.Nama = input.Nama
	//user.Alamat = input.Alamat
	user.Email = input.Email
	user.Password = string(passwordHash)
	user.CreatedAt = time.Now()

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	nama := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(nama)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

// func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
// 	email := input.Email

// 	user, err := s.repository.FindByEmail(email, password)
// 	if err != nil {
// 		return false, err
// 	}

// 	if user.ID == 0 {
// 		return true, nil
// 	}

// 	return false, nil
// }

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID() (User, error) {
	user, err := s.repository.ShowUser()
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (s *service) GetUserByIDAdmin() (User, error) {
	user, err := s.repository.FindByIDAdmin()
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) UpdateUser(input FormUpdateUserInput) (User, error) {
	user, err := s.repository.FindByID(input.ID)
	if err != nil {
		return user, err
	}

	user.Nama = input.Nama
	user.Alamat = input.Alamat
	user.Email = input.Email
	user.Password = input.Password

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) Delete(ID int) (User, error) {
	user, err := s.repository.Delete(ID)
	if err != nil {
		return user, err
	}

	return user, nil
}
