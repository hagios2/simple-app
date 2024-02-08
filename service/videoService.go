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
	videoRepository repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()
}

func (service *videoService) Update(video entity.Video) entity.Video {
	service.videoRepository.Update(video)
	return video
}

func (service *videoService) Delete(video entity.Video) {
	service.videoRepository.Delete(video)
}
