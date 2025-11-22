package ws

type Broadcast struct {
	Room string
	Data []byte
}

type Hub struct {
	clients    map[*Client]struct{}
	register   chan *Client
	unregister chan *Client
	broadcast  chan Broadcast
}

type Client struct {
	send   chan []byte
	room   string
	userID uint
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]struct{}),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan Broadcast, 256),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = struct{}{}
		case c := <-h.unregister:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				close(c.send)
			}
		case b := <-h.broadcast:
			for c := range h.clients {
				// only broadcast to clients in the same room
				if c.room != b.Room {
					continue
				}
				select {
				case c.send <- b.Data:
				default:
					delete(h.clients, c)
					close(c.send)
				}
			}
		}
	}
}

func (h *Hub) RedisBroadcast(b Broadcast) {
	h.broadcast <- b
}
