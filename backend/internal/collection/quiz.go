package collection

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"quiz.com/quiz/internal/entity"
)

// Struct wraps MongoDB collection. Acts as a pointer to the MongoDB quizzes collection-where data is stored
type QuizCollection struct {
	collection *mongo.Collection
}

func Quiz(collection *mongo.Collection) *QuizCollection {
	return &QuizCollection{
		collection: collection,
	}
}

// Function to insert a new quiz into the MongoDB collection
func (c QuizCollection) InsertQuiz(quiz entity.Quiz) error {
	_, err := c.collection.InsertOne(context.Background(), quiz)
	return err
}

// Function retrieves all quizzes from the MongoDB collection
func (c QuizCollection) GetQuizzes() ([]entity.Quiz, error) {
	cursor, err := c.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var quiz []entity.Quiz
	err = cursor.All(context.Background(), &quiz)
	if err != nil {
		return nil, err
	}

	return quiz, nil
}

// Retrieves a single quiz by its MongoDB ObjectID
func (c QuizCollection) GetQuizById(id primitive.ObjectID) (*entity.Quiz, error) {
	result := c.collection.FindOne(context.Background(), bson.M{"_id": id})

	var quiz entity.Quiz
	err := result.Decode(&quiz)
	if err != nil {
		return nil, err
	}

	return &quiz, nil
}
