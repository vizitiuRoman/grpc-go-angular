package pg

import (
	"github.com/jmoiron/sqlx"
	. "github.com/user-service/pkg/domain"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (repo *userRepo) CreateUser(user *User) (*User, error) {
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
		err = rows.StructScan(&user)
		if err != nil {
			return &User{}, err
		}
	}
	return user, nil
}

func (repo *userRepo) GetUser(userID uint64) (*User, error) {
	user := &User{}
	err := repo.db.Get(user, "SELECT email, id FROM users WHERE id = $1", userID)
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (repo *userRepo) UpdateUser(user *User) (*User, error) {
	rows, err := repo.db.Query(`
			UPDATE users 
			SET email=$2, password=$3 
			WHERE id = $1
		`,
		user.ID, user.Email, user.Password,
	)
	if err != nil {
		return &User{}, err
	}
	if rows.Next() {
		err = rows.Scan(&user)
		if err != nil {
			return &User{}, err
		}
	}
	return user, nil
}

func (repo *userRepo) DeleteUser(userID uint64) error {
	_, err := repo.db.Query(`DELETE FROM users WHERE id = $1`,
		userID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepo) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	err := repo.db.Get(user, "SELECT email, id, password FROM users WHERE email = $1", email)
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
