package users

import (
	"github.com/white43/SEP401-pdf-generator/pkg/database"
	"github.com/white43/SEP401-pdf-generator/pkg/errors"
)

type UserRepository struct {
	database *database.Database
}

func NewUserRepository(db *database.Database) *UserRepository {
	return &UserRepository{db}
}

func (ur UserRepository) InsertUser(firstName, lastName, email, password string) (User, error) {
	stmt, err := ur.database.Prepare("INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return User{}, err
	}

	_, err = stmt.Exec(firstName, lastName, email, password)
	if err != nil {
		return User{}, err
	}

	return ur.GetOneByEmail(email)
}

func (ur UserRepository) GetByID(value any) (User, error) {
	return User{}, nil
}

func (ur UserRepository) GetOneByEmail(value any) (User, error) {
	stmt, err := ur.database.Prepare("SELECT * FROM users WHERE email = ?")
	if err != nil {
		return User{}, err
	}

	e := User{}

	if err = stmt.QueryRow(value).Scan(&e.ID, &e.FirstName, &e.LastName, &e.Email, &e.Password, &e.Token, &e.Balance); err != nil {
		return User{}, err
	}

	return e, nil
}

func (ur UserRepository) GetOneByToken(value any) (User, error) {
	stmt, err := ur.database.Prepare("SELECT * FROM users WHERE token = ?")
	if err != nil {
		return User{}, err
	}

	e := User{}

	if err = stmt.QueryRow(value).Scan(&e.ID, &e.FirstName, &e.LastName, &e.Email, &e.Password, &e.Token, &e.Balance); err != nil {
		return User{}, err
	}

	return e, nil
}

func (ur UserRepository) UpdateToken(userId int, token string) (int64, error) {
	result, err := ur.database.Exec("UPDATE users SET token = ? WHERE id = ?", token, userId)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil
}

func (ur UserRepository) AddBalance(userId int, balance float64) error {
	result, err := ur.database.Exec("UPDATE users SET balance = balance + ? WHERE id = ?", balance, userId)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.NoRowsAffected
	}

	return nil
}

func (ur UserRepository) DeductBalance(userId int, balance float64) error {
	return ur.AddBalance(userId, -balance)
}
