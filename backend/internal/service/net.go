package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quiz.com/quiz/internal/entity"
)

type NetService struct {
	quizService *QuizService //\ might need to remove

	games []*Game
}

// Net is a factory function that returns a new instance of the NetService struct.
func Net(quizService *QuizService) *NetService {
	return &NetService{
		quizService: quizService,

		games: []*Game{},
	}
}

type ConnectPacket struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type HostGamePacket struct {
	QuizId string `json:"quizId"`
}

type QuestionShowPacket struct {
	Question entity.QuizQuestion `json:"question"`
}

type ChangeGameStatePacket struct {
	State GameState `json:"state"`
}

type PlayerJoinPacket struct {
	Player Player `json:"player"`
}

type StartGamePacket struct {
}

type PlayerDisconnectPacket struct {
	PlayerId uuid.UUID `json:"playerId"`
}

type SkipPacket struct {
}

type TickPacket struct {
	Tick int `json:"tick"`
}

type QuestionAnswerPacket struct {
	Question int `json:"question"`
}

type PlayerRevealPacket struct {
	Points int `json:"points"`
}

type LeaderboardPacket struct {
	Points []LeaderboardEntry `json:"points"`
}

func (c *NetService) packetIdToPacket(packetId uint8) any {
	switch packetId {
	case 0:
		{
			return &ConnectPacket{}
		}
	case 1:
		{
			return &HostGamePacket{}
		}
	case 5:
		{
			return &StartGamePacket{}
		}
	case 7:
		{
			return &QuestionAnswerPacket{}
		}
	case 11:
		{
			return &SkipPacket{}
		}
	}

	return nil
}

func (c *NetService) packetToPacketId(packet any) (uint8, error) {
	switch packet.(type) {
	case ConnectPacket:
		return 0, nil
	case HostGamePacket:
		return 1, nil
	case QuestionShowPacket:
		return 2, nil // Changed to 2 since 0 and 1 are used for incoming packets
	case ChangeGameStatePacket:
		return 3, nil
	case PlayerJoinPacket:
		return 4, nil
	case TickPacket:
		return 6, nil
	case PlayerRevealPacket:
		return 8, nil
	case LeaderboardPacket:
		return 9, nil
	case PlayerDisconnectPacket:
		return 10, nil
	case SkipPacket:
		return 11, nil
	}

	return 0, errors.New("invalid packet type")
}

func (c *NetService) getGameByCode(code string) *Game {
	for _, game := range c.games {
		if game.Code == code {
			return game
		}
	}
	return nil
}

func (c *NetService) getGameByHost(host *websocket.Conn) *Game {
	for _, game := range c.games {
		if game.Host == host {
			return game
		}
	}
	return nil
}

func (c *NetService) getGameByPlayer(con *websocket.Conn) (*Game, *Player) {
	for _, game := range c.games {
		for _, player := range game.Players {
			if player.Connection == con {
				return game, player
			}
		}
	}
	return nil, nil
}

func (c *NetService) OnDisconnect(con *websocket.Conn) {
	game, player := c.getGameByPlayer(con)
	if game == nil {
		return
	}

	game.OnPlayerDisconnect(player)
}

// The OnIncomingMessage method is called when a message is received from a websocket connection.
func (c *NetService) OnIncomingMessage(con *websocket.Conn, mt int, msg []byte) {

	if len(msg) < 2 {
		fmt.Println("Message too short")
		return
	}

	packetId := msg[0]
	data := msg[1:]

	packet := c.packetIdToPacket(packetId)
	if packet == nil {
		fmt.Printf("Invalid packet ID: %d\n", packetId)
		return
	}

	err := json.Unmarshal(data, packet)
	if err != nil {
		fmt.Printf("Error unmarshaling data: %v\n", err)
		return
	}

	switch data := packet.(type) {
	case *ConnectPacket:
		{
			game := c.getGameByCode(data.Code)
			if game == nil {
				return
			}

			game.OnPlayerJoin(data.Name, con)
			break
		}
	case *HostGamePacket:
		{
			quizId, err := primitive.ObjectIDFromHex(data.QuizId)
			if err != nil {
				fmt.Printf("Invalid Quiz ID: %v\n", err)
				return
			}

			quiz, err := c.quizService.quizCollection.GetQuizById(quizId)
			if err != nil {
				fmt.Printf("Error fetching quiz: %v\n", err)
				return
			}

			if quiz == nil {
				return
			}
			game := newGame(*quiz, con, c)
			fmt.Printf("New game created with code: %s\n", game.Code)
			c.games = append(c.games, &game)

			c.SendPacket(con, HostGamePacket{
				QuizId: game.Code,
			})

			c.SendPacket(con, ChangeGameStatePacket{
				State: game.State,
			})
			break
		}
	case *StartGamePacket:
		{
			fmt.Println("Received StartGamePacket")
			game := c.getGameByHost(con)
			if game == nil {
				fmt.Println("No game found for host connection")
				return
			}
			fmt.Printf("Found game with code: %s, current state: %v\n", game.Code, game.State)
			game.StartOrSkip()
			break
		}
	case *QuestionAnswerPacket:
		{
			game, player := c.getGameByPlayer(con)
			if game == nil {
				return
			}

			game.OnPlayerAnswer(data.Question, player)
			break
		}
	case *SkipPacket:
		{
			game := c.getGameByHost(con)
			if game == nil {
				return
			}
			game.Skip()
			break
		}
	}
}

func (c *NetService) SendPacket(connection *websocket.Conn, packet any) error {
	bytes, err := c.PacketToBytes(packet)
	if err != nil {
		return err
	}

	return connection.WriteMessage(websocket.BinaryMessage, bytes)
}

func (c *NetService) PacketToBytes(packet any) ([]byte, error) {
	packetId, err := c.packetToPacketId(packet)
	if err != nil {
		return nil, errors.New("invalid packet type")
	}

	bytes, err := json.Marshal(packet)
	if err != nil {
		return nil, err
	}

	final := append([]byte{packetId}, bytes...)
	return final, nil
}
