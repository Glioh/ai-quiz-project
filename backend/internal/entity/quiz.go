// WIP file might not work properly.

package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Quiz struct
type Quiz struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name"`
	Questions []QuizQuestion     `json:"questions"`
}

// QuizQuestion struct
type QuizQuestion struct {
	Id      string       `json:"id"`
	Name    string       `json:"name"`
	Choices []QuizChoice `json:"choices"`
}

// QuizChoice struct
type QuizChoice struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Correct bool   `json:"correct"`
}
