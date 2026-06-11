package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MailSender interface {
	Send(to, subject, body string) error
}

type MailHandler struct {
	svc MailSender
}

func NewMailHandler(svc MailSender) *MailHandler {
	return &MailHandler{svc: svc}
}

type sendRequest struct {
	To      string `json:"to"      binding:"required,email"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body"    binding:"required"`
}

func (h *MailHandler) Send(c *gin.Context) {
	var req sendRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.svc.Send(req.To, req.Subject, req.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "메일 발송 실패"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
