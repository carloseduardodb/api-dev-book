package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Posts struct {
	db *sql.DB
}

func NewRepositoryPosts(db *sql.DB) *Posts {
	return &Posts{db}
}

func (postRepo Posts) Create(post models.Post) (uint64, error) {
	statement, err := postRepo.db.Prepare("INSERT INTO posts (title, content, author_id, author_nick) VALUES(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(post.Title, post.Content, post.AuthorId, post.AuthorNick)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

func (postRepo *Posts) Find(titleOrContent string) ([]models.Post, error) {
	titleOrContent = fmt.Sprintf("%%%s%%", titleOrContent)
	statement, err := postRepo.db.Prepare("SELECT id, title, content, author_id, author_nick, likes, created_at FROM posts WHERE title LIKE ? OR content LIKE ?")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query(titleOrContent, titleOrContent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []models.Post{}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorId, &post.AuthorNick, &post.Likes, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (postRepo *Posts) FindById(id uint64) (models.Post, error) {
	statement, err := postRepo.db.Prepare("SELECT id, title, content, author_id, author_nick, likes, created_at FROM posts WHERE id = ?")
	if err != nil {
		return models.Post{}, err
	}
	defer statement.Close()
	row := statement.QueryRow(id)
	var post models.Post
	err = row.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorId, &post.AuthorNick, &post.Likes, &post.CreatedAt)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func (postRepo *Posts) Update(post models.Post) error {
	statement, err := postRepo.db.Prepare("UPDATE posts SET title = ?, content = ?, author_id = ?, author_nick = ?, likes = ?, created_at = ?, updated_at = ?, deleted_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(post.Title, post.Content, post.AuthorId, post.AuthorNick, post.Likes, post.CreatedAt, post.UpdatedAt, post.DeletedAt, post.ID)
	if err != nil {
		return err
	}
	return nil
}

func (postRepo *Posts) Delete(id uint64) error {
	statement, err := postRepo.db.Prepare("DELETE FROM posts WHERE id = ?")
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

func (postRepo *Posts) Like(id uint64) error {
	statement, err := postRepo.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
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

func (postRepo *Posts) Dislike(id uint64) error {
	statement, err := postRepo.db.Prepare("UPDATE posts SET likes = likes - 1 WHERE id = ?")
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
