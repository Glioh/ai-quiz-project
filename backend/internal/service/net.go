package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/contrib/websocket"
	"quiz.com/quiz/internal/entity"
)

type NetService struct {
	quizService *QuizService //\ might need to remove

	host *websocket.Conn

	tick int
}

// Net is a factory function that returns a new instance of the NetService struct.
func Net(quizService *QuizService) *NetService {
	return &NetService{
		quizService: quizService,
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
	}

	return nil
}

func (c *NetService) packetToPacketId(packet any) (uint8, error) {
	switch packet.(type) {
	case QuestionShowPacket:
		return 2, nil // Changed to 2 since 0 and 1 are used for incoming packets
	case ConnectPacket:
		return 0, nil
	case HostGamePacket:
		return 1, nil
	}

	return 0, errors.New("invalid packet type")
}

// The OnIncomingMessage method is called when a message is received from a websocket connection.
func (c *NetService) OnIncomingMessage(con *websocket.Conn, mt int, msg []byte) {
	// Add debug logging
	fmt.Printf("Received message: %v\n", msg)

	if len(msg) < 2 {
		fmt.Println("Message too short")
		return
	}

	packetId := msg[0]
	data := msg[1:]

	fmt.Printf("Packet ID: %d, Data: %v\n", packetId, string(data))

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
			fmt.Printf("Connect packet received: Name=%s, Code=%s\n", data.Name, data.Code)
			break
		}
	case *HostGamePacket:
		{
			fmt.Printf("Host game packet received: QuizId=%s\n", data.QuizId)
			go func() {
				time.Sleep(time.Second * 5)
				c.SendPacket(con, QuestionShowPacket{
					Question: entity.QuizQuestion{
						Name: "What is 2+2",
						Choices: []entity.QuizChoice{
							{
								Name: "4",
							},
							{
								Name: "9",
							},
						},
					},
				})
			}()
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
