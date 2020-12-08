package pg

import (
	. "github.com/user-service/pkg/models"
)

type UserRepo struct {
	db *DB
}

func NewUserRepo(db *DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) CreateUser(user *User) (*User, error) {
	rows, err := repo.db.NamedQuery(`
		INSERT INTO users (email, password)
		VALUES (:email, :password)
		RETURNING email, id;
		`,
		&user,
	)
	if err != nil {
		return &User{}, err
	}
	if rows.Next() {
		rows.StructScan(&user)
	}
	return user, nil
}

func (repo *UserRepo) GetUser(userID uint64) (*User, error) {
	user := &User{}
	err := repo.db.Get(user, "SELECT email, id FROM users WHERE id = $1", userID)
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (repo *UserRepo) UpdateUser(user *User) error {
	_, err := repo.db.Query(`
		UPDATE users 
		SET email=$2, password=$3 
		WHERE id = $1`,
		user.ID, user.Email, user.Password,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) DeleteUser(userID uint64) error {
	_, err := repo.db.Query(`DELETE FROM users WHERE id = $1`,
		userID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	err := repo.db.Get(user, "SELECT email, id, password FROM users WHERE email = $1", email)
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
