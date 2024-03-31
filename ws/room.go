package ws

type Room struct {
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewRoom() *Room {
	return &Room{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func (r *Room) Run() {
	for {
		select {
		case cl := <-r.Register:
			// đăng ký người dùng
			if _, ok := r.Clients[cl.ID]; !ok {
				r.Clients[cl.ID] = cl
			}
		case cl := <-r.Unregister:
			// xóa người dùng
			if _, ok := r.Clients[cl.ID]; ok {
				delete(r.Clients, cl.ID)
			}
		case m := <-r.Broadcast:
			for _, ms := range r.Clients {
				ms.Message <- m
			} 
		}
	}
}
