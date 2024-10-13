package internal

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"quiz.com/quiz/internal/collection"
	"quiz.com/quiz/internal/controller"
	"quiz.com/quiz/internal/service"
)

var quizCollection *mongo.Collection

type App struct {
	httpServer  *fiber.App
	database    *mongo.Database
	quizService *service.QuizService
	netService  *service.NetService
}

func (a *App) Init() {
	a.setupDb()
	a.setupServices()
	a.setupHttp()

	log.Fatal(a.httpServer.Listen(":3000"))

}

func (a *App) setupHttp() {
	app := fiber.New()
	app.Use(cors.New())

	quizController := controller.Quiz(a.quizService)
	app.Get("/api/quizzes", quizController.GetQuizzes)

	wsController := controller.Ws(a.netService)
	app.Get("/ws", websocket.New(wsController.Ws))

	log.Fatal(app.Listen(":3000"))
	a.httpServer = app

}
func (a *App) setupServices() {
	a.quizService = service.Quiz(collection.Quiz(a.database.Collection("quizzes")))
	a.netService = service.Net(a.quizService)
}

func (a *App) setupDb() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	a.database = client.Database("quiz")
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	setupDb()

	app.Get("/", index)
	app.Get("/api/quizzes", getQuizzes)

	// establish websocket connection
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}

	}))

	log.Fatal(app.Listen("127.0.0.1:3000"))
}

// set up database using mongodb by connecting to the local mongodb server
func setupDb() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	// connect to the local mongodb server
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		// any error? crash the program
		panic(err)
	}
	// select the database and collection
	quizCollection = client.Database("quiz").Collection("quizzes")
}

func index(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// getQuizzes returns all quizzes in the database
func getQuizzes(c *fiber.Ctx) error {
	// find all quizzes in the db
	cursor, err := quizCollection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}

	// create a variable to store all the quizzes inside a slice
	quizzes := []map[string]any{}
	// iterate through all the quizzes
	cursor.All(context.Background(), &quizzes)

	return c.JSON(quizzes)
}
