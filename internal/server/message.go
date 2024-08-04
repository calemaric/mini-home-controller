package server

type Message struct {
	ActionType string
	Client     Client
	originId   string
}
