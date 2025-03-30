package handler

import (
	pb "TzTages/api/proto/v1/images"
	"TzTages/internal/repository"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverForImage struct {
	pb.UnimplementedImageServiceServer
	log     *logrus.Entry
	storage repository.StorageImage
}

func NewImageServer(storage repository.StorageImage, log *logrus.Entry) pb.ImageServiceServer {
	return &serverForImage{
		storage: storage,
		log:     log,
	}
}

// UploadImage сохраняет изображение, переданное клиентом, в хранилище и возвращает его метаданные.
// В случае ошибки сохранения возвращает статус Internal.
func (s *serverForImage) UploadImage(ctx context.Context, in *pb.UploadImageRequest) (*pb.UploadImageResponse, error) {
	var response pb.UploadImageResponse
	var err error

	imageUuid := uuid.New().String()
	err = s.storage.Create(context.TODO(), imageUuid, in.FileName, in.Data, int64(len(in.Data)))
	if err != nil {
		// Внутренняя ошибка сервера при сохранении файла
		err = fmt.Errorf("ошибка сохранения изображения: %w", err)
		s.log.Errorf(err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	s.log.Info("Фото успешно загружено в хранилище")
	response.ImageMetadata = &pb.ImageMetadata{
		ImageId:   imageUuid,
		FileName:  in.FileName,
		SizeBytes: int64(len(in.Data)),
	}
	s.log.Infof("ID изображения: %s\nFileName: %s\n", response.ImageMetadata.ImageId, response.ImageMetadata.FileName)

	return &response, nil
}

// DownloadImage возвращает бинарные данные изображения по заданному image_id (uuid).
// Если изображение не найдено — возвращает NotFound, если возникла ошибка при чтении — Internal.
func (s *serverForImage) DownloadImage(ctx context.Context, in *pb.DownloadImageRequest) (*pb.DownloadImageResponse, error) {
	var response pb.DownloadImageResponse

	_, err := s.storage.Search(context.TODO(), in.ImageId)
	if err != nil {
		// Возвращаем NotFound, если изображение не найдено
		err = fmt.Errorf("изображение не найдено: %w", err)
		s.log.Errorf(err.Error())
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	imageGet, err := s.storage.Get(context.TODO(), in.ImageId)
	if err != nil {
		err = fmt.Errorf("ошибка при получении изображения: %w", err)
		s.log.Errorf(err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	s.log.Info("Изображение успешно получено из хранилища")
	response.ImageMetadata = &pb.ImageMetadata{
		ImageId:   imageGet.Uuid,
		FileName:  imageGet.FileName,
		SizeBytes: imageGet.Size,
	}
	s.log.Infof("ID изображения: %s\nFileName: %s\n", response.ImageMetadata.ImageId, response.ImageMetadata.FileName)
	response.Data = imageGet.Data

	return &response, nil
}

// ListImages возвращает список всех изображений без содержимого (только метаданные).
// В случае ошибки при получении списка возвращает статус Internal.
func (s *serverForImage) ListImages(ctx context.Context, in *pb.Empty) (*pb.ListImagesResponse, error) {
	var response pb.ListImagesResponse
	var err error

	listImages, err := s.storage.List(context.TODO())
	if err != nil {
		// Внутренняя ошибка при получении списка изображений
		err = fmt.Errorf("не удалось получить список изображений: %w", err)
		s.log.Errorf(err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if len(listImages) == 0 {
		s.log.Info("в хранилище нет изображений")
		return &response, nil
	}

	s.log.Info("успешно получен список изображений сохранённых в хранилище\n")

	for _, image := range listImages {
		response.ImageMetadata = append(response.ImageMetadata, &pb.ImageMetadata{
			ImageId:   image.Uuid,
			FileName:  image.FileName,
			SizeBytes: image.Size,
		})
		s.log.Infof("ID изображения: %s\nFileName: %s\n", image.Uuid, image.FileName)
	}

	return &response, nil
}
