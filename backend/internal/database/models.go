package database

import "time"

type Photo struct {
	ID         string     `json:"id"`
	Path       string     `json:"path"`
	Filename   string     `json:"filename"`
	Timestamp  time.Time  `json:"timestamp"`
	Albums     []Album    `json:"albums"`
	Categories []Category `json:"categories"`
}

type Category struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}

type Album struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Model struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
}
