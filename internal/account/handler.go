package account

import (
	"encoding/json"
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

	response, err := h.svc.insertAccount(account)
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

	response, err := h.svc.insertTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h Handler) GetAccounts(c *gin.Context) {
	idAccount := c.Param("accountID")

	if len(idAccount) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "id não deve ser nulo"})
		return
	}

	id, err := strconv.ParseInt(idAccount, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	response, err := h.svc.getAccount(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

}
