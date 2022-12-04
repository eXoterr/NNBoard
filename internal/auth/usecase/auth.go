package usecase

import (
	"errors"
	"log"

	"github.com/eXoterr/NNBoard/internal/auth/models"
	"github.com/eXoterr/NNBoard/internal/auth/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUC struct {
	repo repository.Repository
}

func (uc *AuthUC) Register(model models.User) error {
	if !(len(model.Login) > 0) {
		return errors.New("empty login")
	}
	if !(len(model.Password) > 0) {
		return errors.New("empty password")
	}
	if !(len(model.Password2) > 0) {
		return errors.New("empty password2")
	}
	if model.Password != model.Password2 {
		return errors.New("Password doesnt match")
	}

	user, err := uc.repo.GetByLogin(model.Login)

	if user != nil {
		return errors.New("login already registered!")
	}

	hash, err := hashPassword(model.Password)

	if err != nil {
		return err
	}

	return uc.repo.Create(model.Login, string(hash))

}

func (uc *AuthUC) LoginIn(login string, password string) error {

	if isEmpty(login) {
		return errors.New("empty login")
	}

	if isEmpty(password) {
		return errors.New("empty password")
	}

	user, err := uc.repo.GetByLogin(login)

	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("invalid login or password")

	}

	if verifyPassword(user.Password, password) || err != nil {
		if err != nil {
			log.Println(err)
		}
		return errors.New("invalid login or password")
	}

	return nil
}

func NewAuthUC(repo repository.Repository) *AuthUC {
	return &AuthUC{
		repo: repo,
	}
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 14)
}

func verifyPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}

func isEmpty(str string) bool {
	return len(str) == 0
}
