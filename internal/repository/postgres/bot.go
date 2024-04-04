package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"skin-monkey/internal/entity"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) CreateUser(user *entity.User) error {
	query := fmt.Sprintf(`INSERT INTO %s (id, name, user_name, language_code, is_premium, is_bot, date_add, active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) returning id`, usersTable)

	_, err := r.db.Exec(query, user.Id, user.Name, user.UserName, user.LanguageCode, user.IsPremium, user.IsBot, user.DateAdd, user.Active)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserPostgres) GetUser(id string) (*entity.User, error) {

	var user entity.User

	return &user, nil
}

func (r *UserPostgres) GetUsersFilter() (*[]entity.User, error) {
	query := fmt.Sprintf(`SELECT id, name, user_name, language_code, is_premium, is_bot, date_add, active FROM %s WHERE active = true`, usersTable)

	var users []entity.User

	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return &users, nil
}
