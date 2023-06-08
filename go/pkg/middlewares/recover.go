package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/huuhieunguyen/trip-plan-mobile-backend/go/pkg/database"
	"github.com/huuhieunguyen/trip-plan-mobile-backend/go/pkg/utils"
)

func Recover(sc database.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(utils.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					// panic(err)
					return
				}

				appErr := utils.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				// panic(err)
				return
			}
		}()

		c.Next()
	}
}
