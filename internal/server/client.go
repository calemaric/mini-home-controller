package server

import (
	"bytes"
	"calemaric/mini-home-controller/internal/actions"
	"encoding/json"
	"html/template"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	hub        *Connections
	connection *websocket.Conn
	id         string
}

func NewClient(hub *Connections, connection *websocket.Conn) *Client {
	id := uuid.New()

	return &Client{hub: hub, connection: connection, id: id.String()}
}

func (c *Client) SendMessage(message *Message) {
	action, actionErr := actions.GetActionByType(message.ActionType)

	if actionErr != nil {
		log.Println(actionErr)
		return
	}

	err := c.connection.WriteMessage(websocket.TextMessage, GetActionTemplate(action, message))

	if err != nil {
		log.Print("Cannot send a message", err)
	}
}

func (c *Client) Listener() {
	for {
		_, p, err := c.connection.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := &Message{originId: c.id}
		unmarshalErr := json.Unmarshal(p, &message)

		if unmarshalErr != nil {
			log.Println(err)
			return
		}

		c.hub.broadcast <- message
	}
}

func GetActionTemplate(action actions.Action, message *Message) []byte {
	templ, err := template.ParseFiles("web/templates/actions/" + action.GetTemplateName() + ".html")

	if err != nil {
		log.Panic("canont parse template", err)
	}

	doc := bytes.Buffer{}

	actionResult := action.Execute()

	templErr := templ.Execute(&doc, actionResult)

	if templErr != nil {
		log.Panic("cannot execute template")
	}

	return doc.Bytes()
}
