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

func (postRepo *Posts) Find(titleOrContent string, userID uint64) ([]models.Post, error) {
	titleOrContent = fmt.Sprintf("%%%s%%", titleOrContent)
	statement, err := postRepo.db.Prepare(
		`SELECT 
		p.id, p.title, p.content, p.author_id, p.author_nick, p.likes, p.created_at, 
		(SELECT user_id from user_like where user_id = ? and post_id = p.id LIMIT 1) IS NOT NULL as liked 
		FROM posts AS p WHERE p.title LIKE ? OR p.content LIKE ?`)
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query(userID, titleOrContent, titleOrContent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []models.Post{}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorId, &post.AuthorNick, &post.Likes, &post.CreatedAt, &post.Liked)
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
	statement, err := postRepo.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(post.Title, post.Content, post.ID)
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

func (postRepo *Posts) Like(id uint64, userID uint64) error {
	createLikeStatement, err := postRepo.db.Prepare("INSERT INTO user_like (user_id, post_id) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer createLikeStatement.Close()
	_, err = createLikeStatement.Exec(userID, id)
	if err != nil {
		return err
	}
	updateStatement, err := postRepo.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer updateStatement.Close()
	_, err = updateStatement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (postRepo *Posts) Dislike(postID uint64, userID uint64) error {
	deleteLikeStatement, err := postRepo.db.Prepare("DELETE FROM user_like WHERE post_id = ? AND user_id = ?")
	if err != nil {
		return err
	}
	defer deleteLikeStatement.Close()
	_, err = deleteLikeStatement.Exec(postID, userID)
	if err != nil {
		return err
	}
	updateStatement, err := postRepo.db.Prepare("UPDATE posts SET likes = likes - 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer updateStatement.Close()
	_, err = updateStatement.Exec(postID)
	if err != nil {
		return err
	}
	return nil
}
