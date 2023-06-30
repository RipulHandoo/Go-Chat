package ws

import (
	"encoding/json"
	"net/http"

	"github.com/RipulHandoo/goChat/utils"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)


type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}


func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
    type CreateRoomReq struct {
        ID   string `json:"id"`
        Name string `json:"name"`
    }

    decoder := json.NewDecoder(r.Body)
    params := CreateRoomReq{}

    err := decoder.Decode(&params)
    if err != nil {
        utils.ResponseWithError(w, http.StatusBadRequest, err)
        return
    }

    h.hub.Rooms[params.ID] = &Room{
        ID:      params.ID,
        Name:    params.Name,
        Clients: make(map[string]*Client),
    }

    utils.ResponseWithJson(w, http.StatusOK, params)
}
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	roomID := chi.URLParam(r, "roomId")

	// Get the userID and username from the query parameters
	clientID := r.URL.Query().Get("userID")
	username := r.URL.Query().Get("username")

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	m := &Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.hub)
}


func (h *Handler) GetRooms(w http.ResponseWriter, r *http.Request) {
	type RoomRes struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	utils.ResponseWithJson(w,http.StatusOK,rooms)
}


func (h *Handler) GetClients(w http.ResponseWriter,r *http.Request) {
	type ClientRes struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}
	var clients []ClientRes

	roomId := chi.URLParam(r, "roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientRes, 0)
		utils.ResponseWithJson(w,http.StatusOK,clients)
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	utils.ResponseWithJson(w,http.StatusOK,clients)
}