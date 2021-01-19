package account

import (
	"encoding/json"
	"io/ioutil"
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

	var transaction Transaction

	data, _ := ioutil.ReadAll(c.Request.Body)
	if e := json.Unmarshal(data, &transaction); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}

	err := transaction.validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if err = h.svc.insertTransaction(transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "users": "cubas"})
}
