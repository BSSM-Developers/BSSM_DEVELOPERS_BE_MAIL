package service_test

import (
	"testing"

	"github.com/BSSM-Developers/BSSM_DEVELOPERS_BE_MAIL/internal/config"
	"github.com/BSSM-Developers/BSSM_DEVELOPERS_BE_MAIL/internal/service"
)

func TestMailService_BuildMessage(t *testing.T) {
	cfg := config.SMTPConfig{
		From: "sender@gmail.com",
	}
	svc := service.NewMailService(cfg)
	msg := svc.BuildMessage("to@example.com", "제목", "본문 내용")

	if len(msg) == 0 {
		t.Fatal("expected non-empty message")
	}
}
