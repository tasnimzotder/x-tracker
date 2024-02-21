package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ClientMessageT struct {
	UserID   int64 `json:"user_id"`
	DeviceID int64 `json:"device_id"`
}

type MessageT struct {
	DeviceID  int64     `json:"device_id"`
	Lat       float32   `json:"lat"`
	Lon       float32   `json:"lon"`
	Timestamp time.Time `json:"timestamp"`
}

func (s *Server) wsLatestLocation(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
	}(conn)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// read message from client
	_, p, err := conn.ReadMessage()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var clientMessage ClientMessageT
	err = json.Unmarshal(p, &clientMessage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// get user and device
	user, err := s.queries.GetUser(ctx, clientMessage.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	device, err := s.queries.GetDevice(ctx, clientMessage.DeviceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if user.ID != device.UserID {
		ctx.JSON(http.StatusForbidden, errorResponse(
			err,
		))
		return
	}

	// send location data to client
	for {
		select {
		case <-ticker.C:
			locations, err := GetLocationsByDeviceID(s, device.ID)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			jsonMessage, err := json.Marshal(locations)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			err = conn.WriteMessage(websocket.TextMessage, jsonMessage)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			_, reqMessage, _ := conn.ReadMessage()
			if string(reqMessage) == "disconnect" {
				log.Println("Client disconnected")
				err := conn.Close()
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, errorResponse(err))
					return
				}

				return
			}
		}
	}
}
