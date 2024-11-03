package service

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
	Id         uuid.UUID       `json:"id"`
	Name       string          `json:"name"`
	Connection *websocket.Conn `json:"-"`
	Answered   bool            `json:"-"`
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
	Id              uuid.UUID
	Quiz            entity.Quiz
	CurrentQuestion int
	Code            string
	Players         []*Player

	Time       int
	Host       *websocket.Conn
	netService *NetService
	State      GameState
}

func generateCode() string {
	return strconv.Itoa(100000 + rand.Intn(900000)) // Generates a random 4-digit code
}

func newGame(quiz entity.Quiz, host *websocket.Conn, netService *NetService) Game {
	return Game{
		Id:              uuid.New(),
		Quiz:            quiz,
		Code:            generateCode(),
		Players:         []*Player{},
		State:           LobbyState,
		CurrentQuestion: -1,
		Host:            host,
		Time:            60,
		netService:      netService,
	}
}

func (g *Game) Start() {
	g.ChangeState(PlayState)
	g.NextQuestion()

	go func() {
		for {
			g.Tick()
			time.Sleep(time.Second)
		}
	}()
}

func (g *Game) NextQuestion() {
	g.CurrentQuestion++

	g.ChangeState(PlayState)
	g.Time = 60

	g.netService.SendPacket(g.Host, QuestionShowPacket{
		Question: g.Quiz.Questions[g.CurrentQuestion],
	})

}

func (g *Game) Reveal() {
	g.Time = 10
	g.ChangeState(RevealState)
}

func (g *Game) Tick() {
	g.Time--
	g.netService.SendPacket(g.Host, TickPacket{
		Tick: g.Time,
	})

	if g.Time == 0 {
		switch g.State {
		case PlayState:
			{
				g.Reveal()
				break
			}
		case RevealState:
			{
				g.NextQuestion()
				break
			}
		}
	}
}

func (g *Game) ChangeState(state GameState) {
	g.State = state
	g.BroadcastPacket(ChangeGameStatePacket{
		State: state,
	}, true)

}

func (g *Game) BroadcastPacket(packet any, includeHost bool) error {
	for _, player := range g.Players {
		err := g.netService.SendPacket(player.Connection, packet)
		if err != nil {
			return err
		}
	}

	if includeHost {
		err := g.netService.SendPacket(g.Host, packet)
		if err != nil {
			return err
		}
	}

	return nil
}
func (g *Game) OnPlayerJoin(name string, connection *websocket.Conn) {
	fmt.Println("Player", name, "joined the game")

	player := Player{
		Id:         uuid.New(),
		Name:       name,
		Connection: connection,
	}
	g.Players = append(g.Players, &player)

	g.netService.SendPacket(connection, ChangeGameStatePacket{
		State: g.State,
	})

	g.netService.SendPacket(g.Host, PlayerJoinPacket{
		Player: player,
	})
}

func (g *Game) getAnsweredPlayers() []*Player {
	players := []*Player{}

	for _, player := range g.Players {
		if player.Answered {
			players = append(players, player)
		}
	}

	return players
}

func (g *Game) OnPlayerAnswer(choice int, player *Player) {
	player.Answered = true

	if len(g.getAnsweredPlayers()) == len(g.Players) {
		g.Reveal()
	}
}
