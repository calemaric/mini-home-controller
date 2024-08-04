package server

type Connections struct {
	clients map[*Client]bool

	broadcast chan *Message

	register chan *Client

	unregister chan *Client
}

func SetupConnections() *Connections {
	return &Connections{
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (c *Connections) Run() {
	for {
		select {
		case client := <-c.register:
			c.clients[client] = true
			go client.Listener()
		case client := <-c.unregister:
			delete(c.clients, client)
		case message := <-c.broadcast:
			clients := filterClients(c.clients, message)
			for _, client := range clients {
				client.SendMessage(message)
			}
		}
	}
}

func filterClients(clients map[*Client]bool, message *Message) []*Client {
	filteredClients := make([]*Client, 0)

	for client := range clients {
		if client.id == message.originId {
			filteredClients = append(filteredClients, client)
		}
	}

	return filteredClients

}
