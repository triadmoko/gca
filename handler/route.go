package handler

import (
	"gca/model"
	"gca/service"

	"github.com/gin-gonic/gin"
)

type route struct {
	userService service.Service
}

func NewHandlerUser(userService service.Service) *route {
	return &route{userService}
}

func (h *route) GetUser(c *gin.Context) {
	var user model.User
	user, err := h.userService.GetUser()
	if err != nil {
		c.JSON(500, user)
	}
	c.JSON(200, user)

}
