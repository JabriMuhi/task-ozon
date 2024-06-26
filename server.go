package main

import (
	"flag"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	generated "taskOzon/graph"
	"taskOzon/graph/model"
	"taskOzon/internal/service"
	"taskOzon/pkg/db/in_memory"
	database "taskOzon/pkg/db/postgresql"
	"time"
)

const defaultPort = "8080"

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func main() {

	err := godotenv.Load("./docker/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	fmt.Printf("Server started at port: %s \n", port)

	fmt.Printf("Storage type is: %s \n", getEnv("STORAGE_TYPE", "in_memory"))

	storageType := flag.String("storageType", "in_memory", "Storage type: in_memory/db")
	flag.Parse()

	var resolver *generated.Resolver

	switch *storageType {
	case "in_memory":
		// Initialize in-memory
		users := make(map[int]in_memory.User)
		posts := make(map[int]in_memory.Post)
		comments := make(map[int]in_memory.Comment)
		commentsParentsChild := make([]in_memory.CommentParentChild, 0)

		subscriptionMap := make(map[int][]chan<- *model.Comment)

		inMemory := in_memory.InitInMemory(users, posts, comments, commentsParentsChild)

		resolver = generated.NewResolver(service.InitPostServiceInMemory(inMemory), service.InitUserServiceInMemory(inMemory), service.InitCommentServiceInMemory(inMemory, subscriptionMap))
	case "db":
		// Initialize the database
		db := database.InitBD()
		defer func() {
			if err := db.Close(); err != nil {
				log.Printf("Error closing database: %v", err)
			}
		}()

		subscriptionMap := make(map[int][]chan<- *model.Comment)

		resolver = generated.NewResolver(service.InitPostService(db), service.InitUserService(db), service.InitCommentService(db, subscriptionMap))

	default:
		fmt.Println("Invalid storage type!")
	}

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		KeepAlivePingInterval: 10 * time.Second,
	})

	srv.AddTransport(transport.POST{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
