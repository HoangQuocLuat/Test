package ws

import (
	"fmt"
	"thuchanh_go/logic"
	"thuchanh_go/types/req"
)

type Room struct {
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *req.Message
	Chat       logic.ChatLogic
}

func NewRoom() *Room {
	return &Room{
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *req.Message),
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
			for _, cl := range r.Clients {
				cl.Message <- m
				fmt.Printf(m.Content)
			}
		}
	}
}
