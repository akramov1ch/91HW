package services

import (
	"91HW/internal/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(db *mongo.Database) *TaskService {
	return &TaskService{
		collection: db.Collection("tasks"),
	}
}

func (s *TaskService) CreateTask(ctx context.Context, task *models.Task) error {
	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	_, err := s.collection.InsertOne(ctx, task)
	return err
}

func (s *TaskService) GetTasks(ctx context.Context) ([]models.Task, error) {
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	var task models.Task
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, task *models.Task) error {
	_, err := s.collection.UpdateOne(ctx, bson.M{"_id": task.ID}, bson.M{"$set": task})
	return err
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	_, err := s.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

