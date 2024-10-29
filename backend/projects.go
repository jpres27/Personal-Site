package main

import (
	"database/sql"
	"errors"
	"time"
)

type Project struct {
	ID          int
	Title       string
	Description string
	Github      string
	Link        string
	Created     time.Time
}

type ProjectModel struct {
	DB *sql.DB
}

func (m *ProjectModel) Insert(title string, description string, github string, link string) (int, error) {
	stmt := `INSERT INTO projects (title, description, github, link, created) VALUES(?, ?, ?, ?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(stmt, title, description, github, link)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *ProjectModel) Get(id int) (Project, error) {
	var p Project

	stmt := `SELECT id, title, description, github, link, created FROM projects WHERE id = ?`
	err := m.DB.QueryRow(stmt, id).Scan(&p.ID, &p.Title, &p.Description, &p.Github, &p.Link, &p.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Project{}, ErrNoRecord
		} else {
			return Project{}, err
		}
	}
	return p, nil
}

func (m *ProjectModel) Latest() ([]Project, error) {
	stmt := `SELECT id, title, description FROM projects ORDER BY id DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var projects []Project

	for rows.Next() {
		var p Project
		err = rows.Scan(&p.ID, &p.Title, &p.Description)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
