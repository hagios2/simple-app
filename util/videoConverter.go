package util

import (
	"github.com/hagios2/simple-app/entity"
	"github.com/hagios2/simple-app/graph/model"
)

func ModelVideoToEntityVideoConverter(modelVideo model.VideoInput) entity.Video {
	return entity.Video{
		Title:       modelVideo.Title,
		Description: modelVideo.Title,
		URL:         modelVideo.URL,
		Author: entity.Person{
			FirstName: modelVideo.Author.Firstname,
			LastName:  modelVideo.Author.Lastname,
			Age:       modelVideo.Author.Age,
			Email:     modelVideo.Author.Email,
		},
	}
}

func EntityVideoToModelVideoConverter(entityVideo entity.Video) *model.Video {
	return &model.Video{
		Title:       entityVideo.Title,
		Description: entityVideo.Title,
		URL:         entityVideo.URL,
		Author: &model.Person{
			Firstname: entityVideo.Author.FirstName,
			Lastname:  entityVideo.Author.LastName,
			Age:       entityVideo.Author.Age,
			Email:     entityVideo.Author.Email,
		},
	}
}
