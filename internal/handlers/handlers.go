package handlers

type Handlers struct {
	URLHandler  *URLHandler
	PingHandler *PingHandler
}

func NewHandlers(url *URLHandler, ping *PingHandler) *Handlers {
	return &Handlers{
		URLHandler:  url,
		PingHandler: ping,
	}
}
