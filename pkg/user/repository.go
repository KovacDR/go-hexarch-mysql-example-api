package user

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetOne(ctx context.Context, id int) (User, error)
    GetByUsername(ctx context.Context, username string) (User, error)
    Create(ctx context.Context, user *User) error
    Update(ctx context.Context, id int, user User) error
    Delete(ctx context.Context, id int) error
}