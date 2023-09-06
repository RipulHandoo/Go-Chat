package main

import (
	"fmt"
	"log"
	"net/http"
	//"os"

	handler "github.com/RipulHandoo/goChat/handler/auth" // Importing the authentication handler.
	"github.com/RipulHandoo/goChat/middleware"
	"github.com/RipulHandoo/goChat/users"
	"github.com/RipulHandoo/goChat/utils"
	"github.com/RipulHandoo/goChat/ws"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	//"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from a .env file.
	//err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Could not load .env file")
	// }

	// Get the port from environment variables.
	// port := os.Getenv("PORT")
	port := "8080"
	if port == "" {
		log.Fatal("Could not get the port")
	}

	// Create a new Chi router.
	router := chi.NewRouter()

	// Configure CORS (Cross-Origin Resource Sharing) settings.
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Links"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Create a sub-router for version 1 (v1) of your API.
	v1Router := chi.NewRouter()

	// Define API endpoints for v1.
	v1Router.Get("/health", utils.ServerHealth)
	v1Router.Post("/singUp", handler.CreateUser)
	v1Router.Post("/login", handler.LoginUser)
	v1Router.Post("/logout", middleware.Auth(middleware.AuthHandler(handler.LogOut)))

	v1Router.Delete("/delete", middleware.Auth(middleware.AuthHandler(users.DeleteUser)))
	v1Router.Post("/follow", middleware.Auth(middleware.AuthHandler(users.FollowUser)))
	v1Router.Post("/Unfollow", middleware.Auth(middleware.AuthHandler(users.UnFollowUser)))

	// Create a WebSocket hub and handler.
	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	// Define WebSocket endpoints.
	v1Router.Post("/ws/createRoom", wsHandler.CreateRoom)
	v1Router.Get("/ws/joinRoom/{roomId}", wsHandler.JoinRoom)
	v1Router.Get("/ws/getRooms", wsHandler.GetRooms)
	v1Router.Get("/ws/getClients/:roomId", wsHandler.GetClients)

	// Start the WebSocket hub in a goroutine.
	go hub.Run()

	// Mount the v1Router under the "/v1" path.
	router.Mount("/v1", v1Router)

	// Print a message indicating the server is running.
	fmt.Printf("Server is running on port: %s\n", port)

	// Create an HTTP server and start listening on the specified port.
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Could not run the server at port")
	}
}
