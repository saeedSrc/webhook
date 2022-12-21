package controller

import (
	"fmt"
	"net/http"
)

type Controller interface {
	Update(w http.ResponseWriter, r *http.Request)
}

type controller struct {
}

func NewController() Controller {
	c := &controller{}
	return c
}

func (c *controller) Update(w http.ResponseWriter, r *http.Request) {

	//message := &m.ReceiveMessage{}

	chatID := 75277134
	msgText := "thesecondapiisdown"

	//err := json.NewDecoder(r.Body).Decode(&message)
	//if err != nil {
	//	fmt.Println("step 1", err)
	//}

	client := &http.Client{}
	respMsg := fmt.Sprintf("%s%s/sendMessage?chat_id=%d&text=%s", "URL", "TOKEN", chatID, msgText)

	req, err := http.NewRequest("GET", respMsg, http.NoBody)
	if err != nil {
		fmt.Println("step 2", err)
	}

	req.Close = true
	req.Header.Set("Accept-Encoding", "identity")

	_, err = client.Do(req)
	if err != nil {
		fmt.Println("step 3", err)
	}

}
