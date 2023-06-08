package services

import (
	"context"
	"net/http"

	"github.com/huuhieunguyen/trip-plan-mobile-backend/go/pkg/utils"
	// "github.com/huuhieunguyen/trip-plan-mobile-backend/go/pkg/database"
	"github.com/huuhieunguyen/trip-plan-mobile-backend/go/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserServices represent the httphandler for user
type UserServices struct {
	DB *mongo.Database
}

// FetchArticle will fetch the article based on given params
func (a *UserServices) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()
	var Users []models.User
	cursor, err := a.DB.Collection("Users").Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &Users); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, utils.SimpleSuccessResponse(Users))
}

// GetByID will get Users by given id
func (a *UserServices) GetUserByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	// Find the user with the matching ID in the "Users" collection
	var User models.User
	err = a.DB.Collection("Users").FindOne(ctx, bson.M{"_id": objectID}).Decode(&User)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, utils.SimpleSuccessResponse(User))
}

// Create User will create a new User based on given request body
func (a *UserServices) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	var User models.User
	if err := c.ShouldBindJSON(&User); err != nil {
		return
	}

	User.ID = primitive.NewObjectID()
	_, err := a.DB.Collection("Users").InsertOne(ctx, User)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, utils.SimpleSuccessResponse(User))
}

// Update will update a User by given id
// func (a *UserServices) Update(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	id := c.Param("id")
// 	var requestBody models.User

// 	if err := c.BindJSON(&requestBody); err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}
// 	updates := map[string]interface{}{
// 		"name":           requestBody.Name,
// 		"code":           requestBody.Code,
// 		"numOfCorrect":   requestBody.NumOfCorrect,
// 		"numOfWrong":     requestBody.NumOfWrong,
// 		"imageAvatarUrl": requestBody.ImageAvatarUrl,
// 		"parentEmail":    requestBody.ParentEmail,
// 		"note":           requestBody.Note,
// 	}

// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	// Update the user with the matching ID in the "Users" collection
// 	_, err = a.DB.Collection("Users").UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": updates})
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
// }

// Delete will delete a User by given id
// func (a *UserServices) Delete(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	id := c.Param("id")

// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	// Delete the user with the matching ID in the "Users" collection
// 	_, err = a.DB.Collection("Users").DeleteOne(ctx, bson.M{"_id": objectID})
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}
// 	c.JSON(http.StatusNoContent, gin.H{"message": "success"})
// }
