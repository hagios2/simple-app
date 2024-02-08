package service

import (
	"github.com/hagios2/simple-app/entity"
	"github.com/hagios2/simple-app/repository"
)

type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
	Update(video entity.Video) entity.Video
	Delete(video entity.Video)
}

type videoService struct {
	repository repository.VideoRepository
}

func New(videoRepository repository.VideoRepository) VideoService {
	return &videoService{
		repository: videoRepository,
	}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	video = service.repository.Save(video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.repository.FindAll()
}

func (service *videoService) Update(video entity.Video) entity.Video {
	service.repository.Update(video)
	return video
}

func (service *videoService) Delete(video entity.Video) {
	service.repository.Delete(video)
}
