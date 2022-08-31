package models

import "errors"

type Post struct {
	ID         uint64 `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	AuthorId   uint64 `json:"author_id,omitempty"`
	AuthorNick string `json:"author_nick,omitempty"`
	Likes      uint64 `json:"likes,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	DeletedAt  string `json:"deleted_at,omitempty"`
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
