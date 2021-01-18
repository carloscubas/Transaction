package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Handler struct
type Handler struct {
	svc    *Service
	logger *zap.Logger
}

func NewHandler(s *Service, l *zap.Logger) *Handler {
	return &Handler{svc: s, logger: l}
}

func (h Handler) NewTransaction(c *gin.Context) {
	h.svc.insertTransaction()

	c.JSON(http.StatusOK, gin.H{"status": "success", "users": "cubas"})
}
