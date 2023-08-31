package database

import (
	"TranscribeHub_HTMX/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

type Dao interface {
	RegisterUser(user models.RegisterUser) error
	LoginUser(user models.LoginUser) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UsernameExists(username string) (bool, error)
}

type DaoImpl struct {
	*pgx.Conn
}

func NewDao(db *pgx.Conn) Dao {
	return &DaoImpl{db}
}

func (d *DaoImpl) RegisterUser(user models.RegisterUser) error {
	_, err := d.Exec(context.Background(), "INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4);", user.Id, user.Username, user.Email, user.Password)
	return err
}

func (d *DaoImpl) LoginUser(user models.LoginUser) (*models.User, error) {
	var out models.User
	err := d.QueryRow(context.Background(), "SELECT id, username, email FROM users WHERE email = $1 AND password = $2;", user.Email, user.Password).Scan(&out.Id, &out.Username, &out.Email)
	return &out, err
}

func (d *DaoImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := d.QueryRow(context.Background(), "SELECT * FROM users WHERE email = $1;", email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	return &user, err
}

func (d *DaoImpl) UsernameExists(username string) (bool, error) {
	err := d.QueryRow(context.Background(), "SELECT username FROM users WHERE username = $1;", username).Scan(&username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
