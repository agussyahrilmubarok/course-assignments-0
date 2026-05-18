package apiserver

import (
	"app/course/booking"
	"app/course/catalog"
	bookingV1 "app/course/server/booking/v1"
	catalogV1 "app/course/server/catalog/v1"
	"app/internal/config"
	"app/internal/logger"
	"context"
	"fmt"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Option struct {
	Cfg *config.Config
	DB  *gorm.DB
	Log *zap.Logger
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
	log := opts.Log
	log.Info("initializing server")

	s := &Server{
		option: opts,
	}

	s.catalogStore = catalog.NewStore(opts.DB)
	s.catalogService = catalog.NewService(s.catalogStore)
	s.catalogHandler = catalogV1.NewHandler(s.catalogService)

	s.bookingStore = booking.NewStore(opts.DB)
	s.bookingService = booking.NewService(s.bookingStore, s.catalogStore)
	s.bookingHandler = bookingV1.NewHandler(s.bookingService)

	log.Info("server initialized")
	return s
}

func (s *Server) Run(ctx context.Context) error {
	log := s.option.Log

	s.http = s.newHTTPServer()
	log.Info("http server starting",
		zap.String("addr", s.http.Addr),
	)

	go func() {
		if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server error",
				zap.Error(err),
			)
		}
	}()

	<-ctx.Done()

	log.Info("shutdown signal received")

	gracefulShutdownPeriod := 30 * time.Second

	shutdownCtx, cancel := context.WithTimeout(context.Background(), gracefulShutdownPeriod)
	defer cancel()

	if err := s.http.Shutdown(shutdownCtx); err != nil {
		log.Error("graceful shutdown failed",
			zap.Error(err),
		)
		return err
	}

	log.Info("server stopped gracefully")

	return nil
}

func (s *Server) newHTTPServer() *http.Server {
	log := s.option.Log

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(s.requestIDMiddleware())
	router.Use(ginzap.RecoveryWithZap(log, true))

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

func (s *Server) requestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		requestID := c.Request.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.NewString()
		}

		c.Writer.Header().Set("X-Request-ID", requestID)

		reqLogger := logger.Log.With(
			zap.String("request_id", requestID),
		)

		ctx := logger.WithCtx(c.Request.Context(), reqLogger)

		c.Request = c.Request.WithContext(ctx)

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()
		reqLogger.Info("http request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.Int("status", status),
			zap.Duration("latency", duration),
		)
	}
}
