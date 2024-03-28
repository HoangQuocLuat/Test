package chathandler

import (
	"net/http"
	"thuchanh_go/types/req"
	"thuchanh_go/types/res"
	"thuchanh_go/ws"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *ws.Hub
}

func NewHandler(h *ws.Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req req.CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.hub.Rooms[req.ID] = &ws.Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*ws.Client),
	}

	c.JSON(http.StatusOK, req)
}

// xử lý websocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // thử trên postman
		//origin := r.Header.Get("Origin")
		//return origin == "http://localhost:3000"
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomID := c.Param("roomId")
	clientID := c.Query("userId")
	username := c.Query("username")

	cl := &ws.Client{
		Conn:     conn,
		Message:  make(chan *ws.Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	m := &ws.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	//Đăng ký một client mới thông qua channal đăng ký
	h.hub.Register <- cl
	// và phát tin nhắn đó
	h.hub.Broadcast <- m
	//writeMess()
	go cl.WriteMess()
	//readMess()
	cl.ReadMess(h.hub)
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]res.RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, res.RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
}

func (h *Handler) GetClient(c *gin.Context) {
	var client []res.ClientRes
	roomId := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		client = make([]res.ClientRes, 0)
		c.JSON(http.StatusOK, client)
	}

	for _, c := range h.hub.Rooms[roomId].Clients{
		client = append(client, res.ClientRes{
			ID : c.ID,
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, client)
}
