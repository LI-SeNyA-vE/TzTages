package repository

import (
	"TzTages/internal/domain"
	"context"
)

type StorageImage interface {
	// Create сохраняет изображение в хранилище.
	// Принимает UUID, имя файла и бинарные данные.
	// Возвращает ID изображения в БД и ошибку, если есть.
	Create(ctx context.Context, uuid string, fileName string, data []byte, size int64) (err error)

	// Search ищет изображение по имени файла.
	// Возвращает структуру ImageFromBD (включая только метаданные, без бинарных данных) и ошибку, если есть.
	Search(ctx context.Context, uuid string) (image domain.ImageFromBD, err error)

	// Get получает изображение из БД по ID.
	// Возвращает структуру ImageFromBD (включая метаданные и бинарные данные) и ошибку, если есть.
	Get(ctx context.Context, uuid string) (image domain.ImageFromBD, err error)

	// List возвращает список всех изображений без бинарных данных.
	// Используется для отображения списка файлов (ID, имя, размер и т.д.).
	List(ctx context.Context) (list []domain.ImageFromBD, err error)
}
