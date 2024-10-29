package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
	"quiz.com/quiz/internal/entity"
)

type Player struct {
	Name       string
	Connection *websocket.Conn
}

type GameState int

const (
	LobbyState GameState = iota
	PlayState
	RevealState
	EndState
)

// Game struct
type Game struct {
	Id      uuid.UUID
	Quiz    entity.Quiz
	Code    string
	Players []Player

	Host *websocket.Conn
}

func generateCode() string {
	return strconv.Itoa(100000 + rand.Intn(900000)) // Generates a random 4-digit code
}

func New(quiz entity.Quiz, host *websocket.Conn) Game {
	return Game{
		Id:      uuid.New(),
		Quiz:    quiz,
		Code:    generateCode(),
		Players: []Player{},
		Host:    host,
	}
}

func (g *Game) Start() {
	go func() {
		for {
			g.Tick()
			time.Sleep(time.Second)
		}
	}()
}

func (g *Game) Tick() {
	// Implement game tick logic here

}

func (g *Game) OnPlayerJoin(name string, connection *websocket.Conn) {
	fmt.Println("Player", name, "joined the game")
	g.Players = append(g.Players, Player{
		Name:       name,
		Connection: connection,
	})
}
