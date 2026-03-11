package apiserver

import (
	"app/course/booking"
	"app/course/catalog"
	bookingV1 "app/course/server/booking/v1"
	catalogV1 "app/course/server/catalog/v1"
	"app/internal/config"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Option struct {
	Cfg *config.Config
	DB  *gorm.DB
}

type Server struct {
	option *Option
	http   *http.Server

	catalogHandler *catalogV1.Handler
	catalogService *catalog.Service
	catalogStore   *catalog.Store

	bookingHandler *bookingV1.Handler
	bookingService *booking.Service
	bookingStore   *booking.Store
}

func New(opts *Option) *Server {
	s := &Server{
		option: opts,
	}

	s.catalogStore = catalog.NewStore(opts.DB)
	s.catalogService = catalog.NewService(s.catalogStore)
	s.catalogHandler = catalogV1.NewHandler(s.catalogService)

	s.bookingStore = booking.NewStore(opts.DB)
	s.bookingService = booking.NewService(s.bookingStore, s.catalogStore)
	s.bookingHandler = bookingV1.NewHandler(s.bookingService)

	return s
}

func (s *Server) Run(ctx context.Context) error {
	s.http = s.newHTTPServer()

	go func() {
		if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("server error:", err)
		}
	}()

	<-ctx.Done()

	gracefulShutdownPeriod := 30 * time.Second

	shutdownCtx, cancel := context.WithTimeout(context.Background(), gracefulShutdownPeriod)
	defer cancel()

	return s.http.Shutdown(shutdownCtx)
}

func (s *Server) newHTTPServer() *http.Server {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")

	apiV1.GET("/health", s.healthz)

	s.catalogHandler.Register(apiV1)
	s.bookingHandler.Register(apiV1)

	addr := fmt.Sprintf(":%v", s.option.Cfg.Server.Port)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func (s *Server) healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
