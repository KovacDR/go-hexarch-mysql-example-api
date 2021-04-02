package user

import (
	"context"
	"time"

	"github.com/KovacDR/go-mysql-api/internal/storage"
)


type UserRepository struct {
	Storage *storage.Storage
}

func (ur *UserRepository) GetAll(ctx context.Context) ([]User, error) {
	q := `
	SELECT id, username, email, avatar, created_at, updated_at
	FROM users;
	`

	rows, err := ur.Storage.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.UserName, &u.Email, &u.Avatar, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return users, err
		}
		
		users = append(users, u)
	}

	return users, nil
}

func (ur *UserRepository) GetOne(ctx context.Context, id int) (User, error) {
	q := `
	SELECT id, username, email, avatar, created_at, updated_at
	FROM users
	WHERE id = ?;
	`

	row := ur.Storage.DB.QueryRowContext(ctx, q, id)

	var u User
	err := row.Scan(&u.ID, &u.UserName, &u.Email, &u.Avatar, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return u, err
	}
	
	return u, nil
}

func (ur *UserRepository) GetByUsername(ctx context.Context, username string) (User, error) {
	q := `
	SELECT id, username, email, avatar, created_at, updated_at
	FROM users
	WHERE username = ?;
	`
	row := ur.Storage.DB.QueryRowContext(ctx, q, username)

	var u User
	err := row.Scan(&u.ID, &u.UserName, &u.Email, &u.Avatar, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return u, err
	}
	
	return u, nil

}

func (ur *UserRepository) Create(ctx context.Context, user *User) error {
	q := `
	INSERT INTO users (username, email, password)
	VALUES (?, ?, ?)
	RETURNING id;`

	err := user.HashPassword()
	if err != nil {
		return err
	}
	
	row := ur.Storage.DB.QueryRowContext(ctx, q, user.UserName, user.Email, user.Password)

	err = row.Scan(&user.ID)
	if err != nil {
		return err
	}
	
	return nil
}

func (ur *UserRepository) Update(ctx context.Context, id int, user User) error {
	q := `
	SET users username=?, email=?, avatar=?, updated_at=?
	WHERE id = ?;
	`

	stmt, err := ur.Storage.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	
	defer stmt.Close()

	_, err = stmt.Exec(user.UserName, user.Email, user.Avatar, time.Now(), id)
	if err != nil {
		return err
	}
	
	return nil
}

func (ur *UserRepository) Delete(ctx context.Context, id int) error {
	q := `
	DELETE FROM users WHERE id = ?;
	`

	stmt, err := ur.Storage.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	
	return nil
}