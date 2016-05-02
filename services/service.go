package services

import "github.com/gin-gonic/gin"

// Service ...
type Service interface {
	Register(*gin.RouterGroup)
}

// Services ...
type Services []Service

// Register ...
func (s Services) Register(r *gin.RouterGroup) {
	for _, service := range s {
		service.Register(r)
	}
}
