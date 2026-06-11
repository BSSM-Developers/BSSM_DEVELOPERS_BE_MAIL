package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BSSM-Developers/BSSM_DEVELOPERS_BE_MAIL/internal/config"
	"github.com/BSSM-Developers/BSSM_DEVELOPERS_BE_MAIL/internal/handler"
	"github.com/BSSM-Developers/BSSM_DEVELOPERS_BE_MAIL/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()
	logger := newLogger()
	defer logger.Sync()

	mailSvc := service.NewMailService(cfg.SMTP)
	mailHandler := handler.NewMailHandler(mailSvc)

	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/mail/send", mailHandler.Send)

	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		logger.Info("mail service 시작", zap.String("port", cfg.Server.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("서버 시작 실패", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("서버 종료 중...")
}

func newLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}
