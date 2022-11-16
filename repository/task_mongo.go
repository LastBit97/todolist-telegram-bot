package repository

import (
	"context"
	"errors"
	"log"

	"github.com/LastBit97/todolist-telegram-bot/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskMongo struct {
	tasksCollection *mongo.Collection
	ctx             context.Context
}

func NewTaskRepository(mongoCollection *mongo.Collection, ctx context.Context) TaskRepository {
	return &TaskMongo{mongoCollection, ctx}
}

func (tm *TaskMongo) CreateTask(task *model.Task) error {
	_, err := tm.tasksCollection.InsertOne(tm.ctx, task)
	log.Print("task added")

	if err != nil {
		return err
	}

	return nil
}

func (tm *TaskMongo) GetTasks(chatId int64) ([]*model.Task, error) {
	query := bson.M{"chat_id": chatId}

	cursor, err := tm.tasksCollection.Find(tm.ctx, query)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(tm.ctx)

	var tasks []*model.Task

	for cursor.Next(tm.ctx) {
		task := &model.Task{}
		err := cursor.Decode(task)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return []*model.Task{}, nil
	}

	return tasks, nil
}

func (tm *TaskMongo) DeleteTask(taskId string) error {
	obId, _ := primitive.ObjectIDFromHex(taskId)
	query := bson.M{"_id": obId}

	res, err := tm.tasksCollection.DeleteOne(tm.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no task with that Id exists")
	}

	return nil
}
