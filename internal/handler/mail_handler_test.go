package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BSSM-Developers/BSSM_DEVELOPERS_BE_MAIL/internal/handler"
	"github.com/gin-gonic/gin"
)

type mockMailService struct{ err error }

func (m *mockMailService) Send(to, subject, body string) error { return m.err }

func TestMailHandler_Send_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := handler.NewMailHandler(&mockMailService{})
	r.POST("/mail/send", h.Send)

	payload := map[string]string{
		"to":      "user@bssm.hs.kr",
		"subject": "테스트",
		"body":    "본문",
	}
	b, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/mail/send", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestMailHandler_Send_MissingTo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := handler.NewMailHandler(&mockMailService{})
	r.POST("/mail/send", h.Send)

	payload := map[string]string{"subject": "제목", "body": "본문"}
	b, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/mail/send", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
