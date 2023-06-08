package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	GetMainDBConnection() *mongo.Database
}

type appCtx struct {
	db *mongo.Database
}

func NewAppContext(db *mongo.Database) *appCtx {
	return &appCtx{
		db: db,
	}
}

func (ctx *appCtx) GetMainDBConnection() *mongo.Database {
	return ctx.db
}

/* import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	*mongo.Database
}

func NewDatabase() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client.Database("your-database-name"), nil
} */
