package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"net/http"
	"webhook/config"
	"webhook/models"
)

type Controller interface {
	Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type controller struct {
	config *config.Config
	logger *zap.SugaredLogger
}

func NewController(config *config.Config, l *zap.SugaredLogger) Controller {
	c := &controller{
		config: config,
		logger: l,
	}
	return c
}

func (c *controller) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	message := &models.ReceiveMessage{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		c.logger.Errorf("can not decode request body, err is: %v", err)
	}

	client := &http.Client{}
	respMsg := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=%s", c.config.WebHook.Url, c.config.WebHook.Token, c.config.WebHook.ChatID, message.Message)

	req, err := http.NewRequest("GET", respMsg, http.NoBody)
	if err != nil {
		c.logger.Errorf("error on requesting webhook, err is: %v", err)
	}

	req.Close = true
	req.Header.Set("Accept-Encoding", "identity")

	_, err = client.Do(req)
	if err != nil {
		c.logger.Errorf("error on sending request to webhook, err is: %v", err)
	}

}
