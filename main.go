package main

import (
	"context"
	"log"

	"github.com/LastBit97/todolist-telegram-bot/bot"
	"github.com/LastBit97/todolist-telegram-bot/configs"
	"github.com/LastBit97/todolist-telegram-bot/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx             context.Context
	mongoClient     *mongo.Client
	mongoCollection *mongo.Collection
	taskRepository  repository.TaskRepository
)

func init() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	ctx = context.TODO()
	mongoOptions := options.Client().ApplyURI(config.DBUri)
	mongoClient, err = mongo.Connect(ctx, mongoOptions)

	if err != nil {
		panic(err)
	}

	mongoCollection = mongoClient.Database("todolist_db").Collection("tasks")
	taskRepository = repository.NewTaskRepository(mongoCollection, ctx)
}

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	defer mongoClient.Disconnect(ctx)

	bot.StartBot(config.ApiToken)

}
