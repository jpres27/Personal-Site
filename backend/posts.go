package main

import (
	"database/sql"
	"errors"
	"time"
)

type Post struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) Insert(title string, content string) (int, error) {
	stmt := `INSERT INTO posts (title, content, created) VALUES(?, ?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(stmt, title, content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *PostModel) Get(id int) (Post, error) {
	var p Post

	stmt := `SELECT id, title, content, created FROM posts WHERE id = ?`
	err := m.DB.QueryRow(stmt, id).Scan(&p.ID, &p.Title, &p.Content, &p.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Post{}, ErrNoRecord
		} else {
			return Post{}, err
		}
	}
	return p, nil
}

func (m *PostModel) Latest() ([]Post, error) {
	stmt := `SELECT id, title, content, created FROM posts ORDER BY id DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var p Post
		err = rows.Scan(&p.ID, &p.Title, &p.Content, &p.Created)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
