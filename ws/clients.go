package ws

import (
	"log"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID 		 string `json:"id"`
	RoomID   string `json:"roomId"`
	UserName string `json:"username"`
}

type Message struct {
	ID      string  `json:"id"`
	RoomID  string  `json:"roomId"`
	Sender  *Client `json:"sender"`
	Content string  `json:"content"`
}

func (c *Client) WriteMess() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)

	}
}

func (c *Client) ReadMess(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			msg := &Message{
				Content:  string(m),
				RoomID:   c.RoomID,
				UserName: c.UserName,
			}

			hub.Broadcast <- msg
		}
	}
}
