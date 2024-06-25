package main

import (
	"flag"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	generated "taskOzon/graph"
	"taskOzon/internal/service"
	"taskOzon/pkg/db/in_memory"
	database "taskOzon/pkg/db/postgresql"
	"time"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	storageType := flag.String("storageType", "db", "Storage type: In memory/PostgreSQL")
	flag.Parse()
	var resolver *generated.Resolver
	switch *storageType {
	case "in_memory":
		// Initialize in-memory
		users := make(map[int]in_memory.User)
		posts := make(map[int]in_memory.Post)
		comments := make(map[int]in_memory.Comment)
		commentsParentsChild := make([]in_memory.CommentParentChild, 0)

		inMemory := in_memory.InitInMemory(users, posts, comments, commentsParentsChild)

		resolver = generated.NewResolver(service.InitPostServiceInMemory(inMemory), service.InitUserServiceInMemory(inMemory), service.InitCommentServiceInMemory(inMemory))
	case "db":
		// Initialize the database
		db := database.InitBD()
		defer func() {
			if err := db.Close(); err != nil {
				log.Printf("Error closing database: %v", err)
			}
		}()

		resolver = generated.NewResolver(service.InitPostService(db), service.InitUserService(db), service.InitCommentService(db))

	default:
		fmt.Println("Invalid storage type!")
	}

	// Create a new resolver and pass the database connection if needed

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
