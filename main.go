package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	//handler "github.com/RipulHandoo/goChat/handler/Auth"
	handler "github.com/RipulHandoo/goChat/handler/auth"
	"github.com/RipulHandoo/goChat/middleware"
	"github.com/RipulHandoo/goChat/ws"
	"github.com/RipulHandoo/goChat/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Could not load .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Could not get the port")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Links"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", utils.ServerHealth)
	v1Router.Post("/singUp", handler.CreateUser)
	v1Router.Post("/login", handler.LoginUser)
	v1Router.Post("/logout", middleware.Auth(middleware.AuthHandler(handler.LogOut)))
	// v1Router.Post("/delete", middleware.Auth(middleware.AuthHandler(handler.DeleteUser)))
	// v1Router.Delete("/delete", middleware.Auth(middleware.AuthHandler(handler.DeleteUser)))
	// v1Router.Post("/follow", middleware.Auth(middleware.AuthHandler(handler.FollowUser)))
	// v1Router.Post("/Unfollow", middleware.Auth(middleware.AuthHandler(handler.Unfollow)))

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	v1Router.Post("/ws/createRoom", wsHandler.CreateRoom)
	v1Router.Get("/ws/joinRoom/{roomId}", wsHandler.JoinRoom)
	v1Router.Get("/ws/getRooms", wsHandler.GetRooms)
	v1Router.Get("/ws/getClients/:roomId", wsHandler.GetClients)


	go hub.Run()

	router.Mount("/v1", v1Router)

	fmt.Printf("Server is running on port: %s\n", port)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Could not run the server at port")
	}
}
