package models

import "errors"

type Post struct {
	ID         uint64 `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorId   uint64 `json:"author_id"`
	AuthorNick string `json:"author_nick"`
	Likes      uint64 `json:"likes"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}

func (post *Post) Prepare() error {
	post.validation()
	return nil
}

func (post *Post) validation() error {
	if post.Title == "" {
		return errors.New("title is required")
	}
	if post.Content == "" {
		return errors.New("content is required")
	}
	return nil
}
