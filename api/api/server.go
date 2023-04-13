package api

import (
	"api/api/handler"
	"api/config"
	"api/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(
	logger *zap.SugaredLogger,
	settings config.Settings,
	userService service.UserService,
) *Server {
	router := gin.Default()

	router.GET("/get/user/new", handler.GetInformationUsingAPI(logger, userService))
	router.GET("/get/user/all", handler.GetAllUsersOfDB(logger, userService))

	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", settings.Port),
			Handler: router,
		},
	}
}

func (s *Server) Start(logger *zap.SugaredLogger) {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			logger.Debugf("Listen %s\n", err)
		}
	}()
}

func (s *Server) Stop(ctx context.Context) {
	_ = s.httpServer.Shutdown(ctx)
}
