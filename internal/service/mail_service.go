package service

import (
	"fmt"
	"net/smtp"

	"github.com/BSSM-Developers/BSSM_DEVELOPERS_BE_MAIL/internal/config"
)

type MailService struct {
	cfg config.SMTPConfig
}

func NewMailService(cfg config.SMTPConfig) *MailService {
	return &MailService{cfg: cfg}
}

func (s *MailService) Send(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.cfg.Username, s.cfg.Password, s.cfg.Host)
	msg := s.BuildMessage(to, subject, body)
	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	return smtp.SendMail(addr, auth, s.cfg.From, []string{to}, msg)
}

func (s *MailService) BuildMessage(to, subject, body string) []byte {
	headers := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n",
		s.cfg.From, to, subject,
	)
	return []byte(headers + body)
}
