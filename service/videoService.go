package service

import (
	"github.com/hagios2/simple-app/entity"
	"github.com/hagios2/simple-app/graph/model"
	"github.com/hagios2/simple-app/repository"
	"log"
	"strconv"
)

type VideoService interface {
	Save(video entity.Video) entity.Video
	SaveGQL(video entity.Video) model.Video
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

func (service *videoService) SaveGQL(video entity.Video) model.Video {
	video = service.repository.Save(video)
	log.Println("saved video response", video)
	newVideo := model.Video{
		ID:          strconv.FormatUint(video.ID, 10),
		Title:       video.Title,
		URL:         video.URL,
		Description: video.Description,
		Author: &model.Person{
			ID:        strconv.FormatUint(video.Author.ID, 10),
			Firstname: video.Author.FirstName,
			Lastname:  video.Author.LastName,
			Age:       video.Author.Age,
			Email:     video.Author.Email,
		},
	}
	return newVideo
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
