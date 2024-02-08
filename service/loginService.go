package service

import "github.com/gin-gonic/gin"

type LoginService interface {
	Login(username string, password string) bool
	Logout(ctx *gin.Context)
}

type loginService struct {
	username string
	password string
}

func NewLoginService() LoginService {
	return &loginService{
		username: "hagios",
		password: "password",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.username == username &&
		service.password == password
}

func (service *loginService) Logout(ctx *gin.Context) {

}
