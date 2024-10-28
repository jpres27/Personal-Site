package main

import (
	"database/sql"
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

func (m *ProjectModel) Insert(title string, content string, github string, link string) (int, error) {
	return 0, nil
}

func (m *ProjectModel) Get(id int) (Project, error) {
	return Project{}, nil
}

func (m *ProjectModel) Latest() ([]Project, error) {
	return nil, nil
}
