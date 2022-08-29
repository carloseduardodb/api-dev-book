package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewRepositoryUsers(db *sql.DB) *Users {
	return &Users{db}
}

func (userRepo Users) Create(user models.User) (uint64, error) {
	statement, err := userRepo.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

func (userRepo *Users) Find(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	statement, err := userRepo.db.Prepare("SELECT id, name, nick, email, created_at FROM users WHERE name LIKE ? OR nick LIKE ?")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query(nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (userRepo *Users) FindById(id uint64) (models.User, error) {
	statement, err := userRepo.db.Prepare("SELECT id, name, nick, email, created_at FROM users WHERE id = ?")
	if err != nil {
		return models.User{}, err
	}
	defer statement.Close()
	var user models.User
	err = statement.QueryRow(id).Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (userRepo *Users) Update(user models.User) error {
	statement, err := userRepo.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(user.Name, user.Nick, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (userRepo *Users) Delete(id uint64) error {
	statement, err := userRepo.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (userRepo *Users) FindByEmail(email string) (models.User, error) {
	statement, err := userRepo.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer statement.Close()
	var user models.User
	if statement.Next() {
		if err = statement.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}
