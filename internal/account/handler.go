package account

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

func (h Handler) NewAccounts(c *gin.Context) {
	var account Account

	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if e := json.Unmarshal(data, &account); e != nil {
		h.logger.Error(e.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}

	err = account.validate()
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	response, err := h.svc.InsertAccount(account)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

}

func (h Handler) NewTransaction(c *gin.Context) {

	var transaction Transaction

	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if e := json.Unmarshal(data, &transaction); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}

	err = transaction.validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	response, err := h.svc.InsertTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) GetAccounts(c *gin.Context) {
	idAccount := c.Param("accountID")

	id, err := strconv.ParseInt(idAccount, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	response, err := h.svc.GetAccount(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if response == nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": fmt.Sprintf("Account id: %s not found", idAccount)})
		return
	}

	c.JSON(http.StatusOK, response)
}
