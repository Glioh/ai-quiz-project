package service

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"quiz.com/quiz/internal/collection"
	"quiz.com/quiz/internal/entity"
)

type QuizService struct {
	quizCollection *collection.QuizCollection
}

func Quiz(quizCollection *collection.QuizCollection) *QuizService {
	return &QuizService{
		quizCollection: quizCollection,
	}
}

func (s QuizService) CreateQuiz(name string) (*entity.Quiz, error) {
	quiz := entity.Quiz{
		Id:        primitive.NewObjectID(),
		Name:      name,
		Questions: []entity.QuizQuestion{},
	}

	err := s.quizCollection.InsertQuiz(quiz)
	if err != nil {
		return nil, err
	}

	return &quiz, nil
}

func (s QuizService) GetQuizById(id primitive.ObjectID) (*entity.Quiz, error) {
	return s.quizCollection.GetQuizById(id)
}

func (s QuizService) UpdateQuiz(id primitive.ObjectID, name string, questions []entity.QuizQuestion) error {
	quiz, err := s.quizCollection.GetQuizById(id)
	if err != nil {
		return err
	}

	if quiz == nil {
		return errors.New("quiz not found")
	}

	quiz.Name = name
	quiz.Questions = questions
	return s.quizCollection.UpdateQuiz(*quiz)
}

func (s QuizService) DeleteQuiz(id primitive.ObjectID) error {
	return s.quizCollection.DeleteQuiz(id)
}

func (s *QuizService) GenerateAIQuiz(name string, prompt string, numQuestions int) (*entity.Quiz, error) {
	aiService := NewAIService()
	questions, err := aiService.GenerateQuiz(prompt, numQuestions)
	if err != nil {
		return nil, err
	}

	quiz := entity.Quiz{
		Id:        primitive.NewObjectID(),
		Name:      name,
		Questions: questions,
	}

	err = s.quizCollection.InsertQuiz(quiz)
	if err != nil {
		fmt.Println("Error generating quiz:", err)
		return nil, err
	}

	fmt.Println("Generated quiz:", quiz)
	return &quiz, nil
}

func (s QuizService) GetQuizzes() ([]entity.Quiz, error) {
	return s.quizCollection.GetQuizzes()
}
