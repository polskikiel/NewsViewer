package main

import (
	"NewsViewer/internal/handlers"
	"NewsViewer/internal"
)

func main() {
	handlers.Run(handlers.NewServer(":8888", handlers.NewHandler(internal.NewUsersInMemory())))
}
