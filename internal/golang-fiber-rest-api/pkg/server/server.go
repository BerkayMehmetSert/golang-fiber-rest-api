package server

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang-fiber-rest-api/configuration"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// server, Fiber uygulamasını yöneten bir yapıdır.
type server struct {
	app *fiber.App
}

// NewServer, yeni bir server örneği oluşturur.
func NewServer(app *fiber.App) *server {
	return &server{
		app: app,
	}
}

// StartHttpServer, HTTP sunucusunu başlatır ve graceful shutdown işlemini yönetir.
func (s *server) StartHttpServer() {
	go func() {
		gracefulShutdown(s.app)
	}()

	if err := s.app.Listen(fmt.Sprintf(":%s", configuration.Port)); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Cannot start server - ERROR: %v\n", err)
		panic("cannot start server")
	}
}

// gracefulShutdown, sunucunun düzgün bir şekilde kapanmasını sağlar.
func gracefulShutdown(app *fiber.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutdown Server")

	_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		fmt.Printf("Server Shutdown Error: %v\n", err)
	}

	fmt.Println("Server exiting")
}
