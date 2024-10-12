package entity

import (
	"github.com/google/uuid"
)

// Game struct
type Game struct {
	Id              uuid.UUID
	Quiz            Quiz
	CurrentQuestion int
	Code            string
}
