package users

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/white43/SEP401-pdf-generator/pkg/dto"
	"github.com/white43/SEP401-pdf-generator/pkg/errors"
	"github.com/white43/SEP401-pdf-generator/pkg/mail"
	"github.com/white43/SEP401-pdf-generator/pkg/random"
	"regexp"
	"strconv"
)

const emailRegexp string = "^[a-zA-Z0-9!#$%&\\'*+\\\\/=?^_`{|}~-]+(?:\\.[a-zA-Z0-9!#$%&\\'*+\\\\/=?^_`{|}~-]+)*@(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?\\.)+[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$"

type Service struct {
	repository *UserRepository
	mail       *mail.Mail
}

func NewService(repository *UserRepository, mail *mail.Mail) *Service {
	return &Service{repository, mail}
}

func (s Service) ValidateRegistrationRequest(request dto.RegistrationRequest) error {
	if request.Email == "" {
		return errors.EmptyEmail
	}
	if request.FirstName == "" {
		return errors.EmptyFirstName
	}
	if request.LastName == "" {
		return errors.EmptyLastName
	}

	check := regexp.MustCompile(emailRegexp)

	if !check.MatchString(request.Email) {
		return errors.InvalidEmail
	}

	user, err := s.repository.GetOneByEmail(request.Email)
	if err != nil && err != sql.ErrNoRows {
		return err
	} else if user.ID > 0 {
		return errors.EmailIsAlreadyTaken
	}

	return nil
}

func (s Service) ValidateLoginRequest(request dto.LoginRequest) error {
	if request.Email == "" {
		return errors.EmptyEmail
	}
	if request.Password == "" {
		return errors.EmptyPassword
	}

	check := regexp.MustCompile(emailRegexp)

	if !check.MatchString(request.Email) {
		return errors.InvalidEmail
	}

	return nil
}

func (s Service) ValidateTopupRequest(request dto.TopupRequest) error {
	if request.Amount == "" {
		return errors.EmptyAmount
	}

	amount, err := strconv.ParseFloat(request.Amount, 10)
	if err != nil {
		return errors.WrongAmount
	}

	if amount < 10 {
		return errors.MinimumPayment10Dollars
	}

	return nil
}

func (s Service) Register(body dto.RegistrationRequest) error {
	password := random.String(4)

	user, err := s.repository.InsertUser(body.FirstName, body.LastName, body.Email, password)
	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.NewServerError("something went wrong during user creation process")
	}

	err = s.mail.SendMail(
		"no-reply@pdf-generator.io",
		user.Email,
		"Your password for pdf-generator.io",
		fmt.Sprintf("Your login: %s\r\nYour password: %s", user.Email, password),
	)

	if err != nil {
		return errors.NewServerError(err.Error())
	}

	return errors.NewMessage(201, "check the email you have provided for login credentials", errors.Options{Headers: []string{}})
}

func (s Service) Login(body dto.LoginRequest) (string, error) {
	randomBytes := make([]byte, 32)
	_, _ = rand.Read(randomBytes)

	user, err := s.repository.GetOneByEmail(body.Email)
	if err == sql.ErrNoRows {
		return "", errors.UserNotFound
	} else if err != nil {
		return "", errors.NewServerError(err.Error())
	}

	if user.Password == "" {
		return "", errors.NewServerError("no password defined")
	}

	if user.Password != body.Password {
		return "", errors.WrongPassword
	}

	token := hex.EncodeToString(randomBytes)

	affected, err := s.repository.UpdateToken(user.ID, token)
	if err != nil {
		return "", err
	} else if affected == 0 {
		return "", errors.NoRowsAffected
	}

	return token, nil
}

func (s Service) GetUserByToken(token string) (User, error) {
	return s.repository.GetOneByToken(token)
}

func (s Service) AddBalance(userId int, body dto.TopupRequest) error {
	amount, err := strconv.ParseFloat(body.Amount, 10)
	if err != nil {
		return err
	}

	return s.repository.AddBalance(userId, amount)
}

func (s Service) DeductBalance(userId int, balance float64) error {
	return s.repository.DeductBalance(userId, balance)
}
