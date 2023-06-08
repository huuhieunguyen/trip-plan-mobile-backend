package handlers

import (
	"github.com/huuhieunguyen/trip-plan-mobile-backend/go/services"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	Handler services.UserServices
}

func (ah *UserHandlers) GetUsers(c *gin.Context) {
	ah.Handler.GetUsers(c)
}

func (ah *UserHandlers) GetUserByID(c *gin.Context) {
	ah.Handler.GetUserByID(c)
}

func (ah *UserHandlers) CreateUser(c *gin.Context) {
	ah.Handler.CreateUser(c)
}

// func (ah *UserHandlers) UpdateAnUser(c *gin.Context) {
// 	ah.Handler.UpdateUser(c)
// }

// func (ah *UserHandlers) DeleteAnUser(c *gin.Context) {
// 	ah.Handler.DeleteUser(c)
// }
