package server

import (
	"baguette/go-todo-c/config"
	// "context"
	// "os"
	// "os/signal"
	// "syscall"
	// "time"

	// "net/http"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	gin    *gin.Engine
	cfg    *config.Configuration
	db     *gorm.DB
	logger *logrus.Logger
}

func NewServer(cfg *config.Configuration, db *gorm.DB, logger *logrus.Logger) *Server {
	server := &Server{
		gin: gin.Default(),
		cfg:    cfg,
		db:     db,
		logger: logger,
		
	}
	server.MapHandler(server.gin)
	server.logger.Logf(logrus.InfoLevel, "Server listening on port %s", server.cfg.Port)
	return server
}

func (s *Server) Run() {
	p := ":" +s.cfg.Port
	s.gin.Run(p)
	// server := &http.Server{
	// 	Addr:    ":" + s.cfg.Port,
	// 	Handler: s.gin,
	// }
	// if err := s.MapHandler(s.gin); err != nil {
	// 	s.logger.Fatal(err)
	// }
	// // p := ":"+s.cfg.Port
	// s.logger.Logf(logrus.InfoLevel, "Server listening on port %s", s.cfg.Port)
	// if err := server.ListenAndServe(); err != nil {
	// 	s.logger.Fatal("Error starting server ", err)
	// }
	
	// if s.ready != nil {
	// 	s.ready <- true
	// }
	// quit := make(chan os.Signal,1)
	// signal.Notify(quit,os.Interrupt,syscall.SIGTERM)
	// <- quit

	// ctx,cancel := context.WithTimeout(context.Background(),5 * time.Second)

	// defer cancel()
	// s.logger.Fatalln("server exited properly")

	// return server.Shutdown(ctx)
}
