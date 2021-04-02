package post

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Post, error)
	GetOne(ctx context.Context, id int) (Post, error)
	GetByUser(ctx context.Context, userID int) ([]Post, error)
	Create(ctx context.Context, post *Post) error
	Update(ctx context.Context, id int, post Post) error
	Delete(ctx context.Context, id int) error
}