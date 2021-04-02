package post

import (
	"context"
	"time"

	"github.com/KovacDR/go-mysql-api/internal/storage"
)


type PostRepository struct {
	Storage *storage.Storage
}

func (pr *PostRepository) GetAll(ctx context.Context) ([]Post, error) {
	q := `
	SELECT id, body, image, user_id, created_at, updated_at
	FROM posts;
	`

	rows, err := pr.Storage.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		err = rows.Scan(&p.ID, &p.Body, &p.Image, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return posts, err
		}
		
		posts = append(posts, p)
	}

	return posts, nil
}

func (pr *PostRepository) GetOne(ctx context.Context, id int) (Post, error) {
	q := `
    SELECT id, body, image, user_id, created_at, updated_at
        FROM posts WHERE id = ?;
    `

	row := pr.Storage.DB.QueryRowContext(ctx, q, id)

	var p Post
	err := row.Scan(&p.ID, &p.Body, &p.Image, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return p, err
	}
	
	return p, nil
}

func (pr *PostRepository) GetByUser(ctx context.Context, userID int) ([]Post, error) {
	q := `
    SELECT id, body, image, user_id, created_at, updated_at
        FROM posts WHERE user_id = ?;
    `

	rows, err := pr.Storage.DB.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var p Post
		err = rows.Scan(&p.ID, &p.Body, &p.Image, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return posts, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func (pr *PostRepository) Create(ctx context.Context, post *Post) error {
	q := `
	INSERT INTO posts (body, user_id, image)
	VALUES (?, ?, ?)
	RETURNING id;
	`

	stmt, err := pr.Storage.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	
	defer stmt.Close()

	_, err = stmt.Exec(post.Body, post.Image, post.UserID)
	if err != nil {
		return err
	}
	
	return nil

}

func (pr *PostRepository) Update(ctx context.Context, id int, post Post) error {
	q := `
	UPDATE posts SET body=?, image=?, created_at=?
	WHERE id = ?;
	`

	stmt, err := pr.Storage.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	
	defer stmt.Close()

	_, err = stmt.Exec(post.Body, post.Image, time.Now(), id)
	if err != nil {
		return err
	}
	
	return nil
}

func (pr *PostRepository) Delete(ctx context.Context, id int) error {
	q := `
	DELETE FROM posts WHERE id = ?;
	`

	stmt, err := pr.Storage.DB.PrepareContext(ctx, q)
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