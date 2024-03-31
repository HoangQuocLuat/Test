package chathandler

import (
	"net/http"
	"thuchanh_go/logic"
	"thuchanh_go/models"
	"thuchanh_go/types/req"
	"thuchanh_go/ws"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	Chat logic.ChatLogic
}

func NewHandler(c logic.ChatLogic) *Handler {
	return &Handler{
		Chat: c,
	}
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req req.CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Chat.Insert(&gin.Context{}, req)
	if err != nil {
		c.JSON(http.StatusConflict, models.Response{
			StatusCode: http.StatusConflict,
			Message:    "không tạo được phòng",
			Data:       nil,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetRooms(c *gin.Context) {
	var req req.CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}
	res, err := h.Chat.Select(c, req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Message:    "Lấy thông tin thành công",
		Data:       res,
	})
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
	//nâng cấp kết nối HTTP lên kết nối Websocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//lấy dữ liệu từ URL
	roomID := c.Param("roomId")
	userID := c.Query("userId")
	username := c.Query("username")

	cl := &ws.Client{
		Conn:     conn,
		Message:  make(chan *ws.Message),
		ID:       userID,
		RoomID:   roomID,
		UserName: username,
	}

	m := &ws.Message{
		Content: "A new user has joined the room",
		RoomID:  roomID,
		UserID:  userID,
	}

	//Đăng ký một client mới thông qua channal đăng ký
	h.Hub.Register <- cl
	// và phát tin nhắn đó
	h.Hub.Broadcast <- m
	//writeMess()
	go cl.WriteMess()
	//readMess()
	cl.ReadMess(h.Hub)
}

// func (h *Handler) GetClient(c *gin.Context) {
// 	var client []res.ClientRes
// 	roomId := c.Param("roomId")

// 	if _, ok := h.Hub.Rooms[roomId]; !ok {
// 		client = make([]res.ClientRes, 0)
// 		c.JSON(http.StatusOK, client)
// 	}

// 	for _, c := range h.Hub.Rooms[roomId].Clients {
// 		client = append(client, res.ClientRes{
// 			ID:       c.ID,
// 			Username: c.Username,
// 		})
// 	}

// 	c.JSON(http.StatusOK, client)
// }
