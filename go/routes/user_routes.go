package routes

import (
	"github.com/huuhieunguyen/trip-plan-mobile-backend/go/handlers"
	"github.com/huuhieunguyen/trip-plan-mobile-backend/go/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// func UserRouter(r *gin.RouterGroup) {
// 	userRouter := r.Group("/users")
// 	{
// 		userRouter.POST("/", handlers.CreateUser)
// 		userRouter.GET("/", handlers.GetUser)
// 		userRouter.PUT("/", handlers.UpdateUser)
// 		userRouter.DELETE("/", handlers.DeleteUser)

// 		userRouter.POST("/login", handlers.Login)
// 		userRouter.POST("/google-login", handlers.GoogleLogin)
// 		userRouter.PUT("/password", handlers.UpdatePassword)
// 		userRouter.PUT("/password/reset", handlers.ResetPassword)

// 	}
// }

func UserRoutes(gin *gin.RouterGroup, db *mongo.Database) {

	handler := handlers.UserHandlers{
		Handler: services.UserServices{
			DB: db,
		},
	}

	users := gin.Group("/users")
	{
		users.GET("", handler.GetUsers)
		users.GET("/:id", handler.GetUserByID)
		users.POST("", handler.CreateUser)
		// users.PUT("/:id", handler.UpdateAnArticle)
		// users.DELETE("/:id", handler.DeleteAnArticle)
	}
}
