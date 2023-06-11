package service

import "gin_api/app/models"

type VideoService interface {
	Save(models.Video) models.Video
	FindAll() []models.Video
}

type videoService struct {
	videos []models.Video
}

func (service *videoService) Save(video models.Video) models.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []models.Video {
	return service.videos
}

func New() VideoService {
	return &videoService{}
}
