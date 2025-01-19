package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Manager struct {
	db *sql.DB
}

func NewManager(dbPath string) (*Manager, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Manager{db: db}, nil
}

func (m *Manager) Close() error {
	return m.db.Close()
}

func (m *Manager) GetPhotos(filter PhotoFilter) ([]Photo, error) {
	query := `
		SELECT DISTINCT p.id, p.path, p.filename, p.timestamp
		FROM photos p
		LEFT JOIN photo_categories pc ON p.id = pc.photo_id
		LEFT JOIN photo_albums pa ON p.id = pa.photo_id
		WHERE 1=1
	`
	args := []interface{}{}

	if filter.Category != "" {
		query += " AND pc.category_id = ?"
		args = append(args, filter.Category)
	}

	if filter.Album != "" {
		query += " AND pa.album_id = ?"
		args = append(args, filter.Album)
	}

	if filter.Sort == "date_asc" {
		query += " ORDER BY p.timestamp ASC"
	} else {
		query += " ORDER BY p.timestamp DESC"
	}

	rows, err := m.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query photos: %w", err)
	}
	defer rows.Close()

	var photos []Photo
	for rows.Next() {
		var p Photo
		if err := rows.Scan(&p.ID, &p.Path, &p.Filename, &p.Timestamp); err != nil {
			return nil, fmt.Errorf("failed to scan photo row: %w", err)
		}

		if err := m.loadPhotoRelations(&p); err != nil {
			return nil, err
		}

		photos = append(photos, p)
	}

	return photos, nil
}

func (m *Manager) loadPhotoRelations(p *Photo) error {
	// Load categories
	rows, err := m.db.Query(`
		SELECT c.id, c.name, pc.confidence
		FROM categories c
		JOIN photo_categories pc ON c.id = pc.category_id
		WHERE pc.photo_id = ?
	`, p.ID)
	if err != nil {
		return fmt.Errorf("failed to query photo categories: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var c Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Confidence); err != nil {
			return fmt.Errorf("failed to scan category row: %w", err)
		}
		p.Categories = append(p.Categories, c)
	}

	// Load albums
	rows, err = m.db.Query(`
		SELECT a.id, a.name, a.description, a.created_at, a.updated_at
		FROM albums a
		JOIN photo_albums pa ON a.id = pa.album_id
		WHERE pa.photo_id = ?
	`, p.ID)
	if err != nil {
		return fmt.Errorf("failed to query photo albums: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var a Album
		if err := rows.Scan(&a.ID, &a.Name, &a.Description, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return fmt.Errorf("failed to scan album row: %w", err)
		}
		p.Albums = append(p.Albums, a)
	}

	return nil
}

func (m *Manager) GetAlbums() ([]Album, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM albums ORDER BY name`
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query albums: %w", err)
	}
	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var a Album
		if err := rows.Scan(&a.ID, &a.Name, &a.Description, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan album row: %w", err)
		}
		albums = append(albums, a)
	}
	return albums, nil
}

func (m *Manager) CreateAlbum(album Album) (*Album, error) {
	query := `INSERT INTO albums (id, name, description, created_at, updated_at) VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`
	result, err := m.db.Exec(query, album.ID, album.Name, album.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to create album: %w", err)
	}

	// Get the created album
	id, _ := result.LastInsertId()
	var created Album
	err = m.db.QueryRow(`SELECT id, name, description, created_at, updated_at FROM albums WHERE id = ?`, id).
		Scan(&created.ID, &created.Name, &created.Description, &created.CreatedAt, &created.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch created album: %w", err)
	}

	return &created, nil
}

func (m *Manager) UpdateAlbum(album Album) (*Album, error) {
	query := `UPDATE albums SET name = ?, description = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	if _, err := m.db.Exec(query, album.Name, album.Description, album.ID); err != nil {
		return nil, fmt.Errorf("failed to update album: %w", err)
	}

	var updated Album
	err := m.db.QueryRow(`SELECT id, name, description, created_at, updated_at FROM albums WHERE id = ?`, album.ID).
		Scan(&updated.ID, &updated.Name, &updated.Description, &updated.CreatedAt, &updated.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated album: %w", err)
	}

	return &updated, nil
}

func (m *Manager) DeleteAlbum(id string) error {
	query := `DELETE FROM albums WHERE id = ?`
	if _, err := m.db.Exec(query, id); err != nil {
		return fmt.Errorf("failed to delete album: %w", err)
	}
	return nil
}
