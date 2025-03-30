package liteSQL

import (
	"TzTages/internal/domain"
	"context"
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteForImage struct {
	db *sql.DB
}

func NewSQLiteStorage() *SQLiteForImage {
	dbPath := filepath.Join("data", "images.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(fmt.Errorf("ошибка подключения к SQL базе %w", err))
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS images (
		id INTEGER PRIMARY KEY,
		uuid TEXT NOT NULL UNIQUE,
		filename TEXT NOT NULL,
		data BLOB NOT NULL,
		size INTEGER NOT NULL
	);
	`
	if _, err = db.Exec(createTable); err != nil {
		log.Fatal(fmt.Errorf("ошибка инициализации таблицы: %w", err))
	}

	return &SQLiteForImage{
		db: db,
	}
}

func (s SQLiteForImage) Create(ctx context.Context, uuid string, fileName string, data []byte, size int64) (err error) {
	_, err = s.db.Exec("INSERT INTO images (uuid, fileName, data, size) VALUES($1, $2, $3, $4)", uuid, fileName, data, size)
	return err
}

func (s SQLiteForImage) Search(ctx context.Context, uuid string) (image domain.ImageFromBD, err error) {
	err = s.db.QueryRow("SELECT id, uuid, filename FROM images WHERE uuid = $1", uuid).
		Scan(&image.ID, &image.Uuid, &image.FileName)
	return image, err
}

func (s SQLiteForImage) Get(ctx context.Context, uuid string) (image domain.ImageFromBD, err error) {
	err = s.db.QueryRow("SELECT id, uuid, filename, data FROM images WHERE uuid = $1", uuid).
		Scan(&image.ID, &image.Uuid, &image.FileName, &image.Data)
	return image, err
}

func (s SQLiteForImage) List(ctx context.Context) (list []domain.ImageFromBD, err error) {
	query, err := s.db.Query("SELECT id, uuid, filename, size FROM images")
	if err != nil {
		return nil, err
	}
	defer query.Close()

	for query.Next() {
		var image domain.ImageFromBD
		err = query.Scan(&image.ID, &image.Uuid, &image.FileName, &image.Size)
		if err != nil {
			return nil, err
		}

		list = append(list, image)
	}

	return list, nil
}
