package main

import (
	"context"
	"fmt"
	"log"

	"github.com/LastBit97/todolist-telegram-bot/bot"
	"github.com/LastBit97/todolist-telegram-bot/configs"
	"github.com/LastBit97/todolist-telegram-bot/repository"
	"github.com/LastBit97/todolist-telegram-bot/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx             context.Context
	mongoClient     *mongo.Client
	mongoCollection *mongo.Collection
	taskRepository  repository.TaskRepository
	taskService     service.TaskService
	telegramBot     bot.Handler
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

	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	mongoCollection = mongoClient.Database("todolist_db").Collection("tasks")
	taskRepository = repository.NewTaskRepository(mongoCollection, ctx)
	taskService = service.NewTaskService(taskRepository)
	telegramBot = bot.NewHandler(taskService)
}

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	defer mongoClient.Disconnect(ctx)

	telegramBot.StartBot(config.ApiToken)
}
