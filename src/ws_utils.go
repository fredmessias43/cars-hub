package main

import (
	"log"

	"github.com/fredmessias43/car-hub/src/config"
	"github.com/fredmessias43/car-hub/src/contracts"
	"github.com/fredmessias43/car-hub/src/websocket"
)

func Emit(roomName string, model contracts.Model) error {
	room := config.WS.FindRoomByName(roomName)
	if room == nil {
		log.Println("Room \"" + roomName + "\" not found")
		return nil
	}
	// log.Println(strconv.Itoa(model.ID) + " => " + room.Name)

	message := websocket.Message{
		Action:  websocket.SendMessageAction,
		Message: string(model.ToJson()),
		Target:  room,
	}

	room.BroadcastToClientsInRoom(message.Encode())
	return nil
}
