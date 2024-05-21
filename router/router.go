package router

import (
	"NotificationService/constant"
	"NotificationService/controller"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Router interface {
	Serve()
}

type routerImpl struct {
	notificationController controller.NotificationController
}

func NewRouter(
	notificationController controller.NotificationController,
) Router {
	return &routerImpl{
		notificationController: notificationController,
	}
}

func (routes routerImpl) Serve() {
	server := echo.New()

	router := server.Group(constant.PathPrefix)
	router.POST("/send", routes.notificationController.Send)

	routes.start(server)
}

func (routes routerImpl) start(server *echo.Echo) {
	server.Debug = true

	go func() {
		if err := server.Start(constant.ServerAddress); !errors.Is(err, http.ErrServerClosed) {
			server.Logger.Fatalf("%v", err)
		}
	}()

	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, syscall.SIGTERM)
	<-sigquit
	if err := server.Shutdown(context.Background()); err != nil {
		server.Logger.Fatal(err)
	}

	server.Logger.Info("Server stopped")
}
