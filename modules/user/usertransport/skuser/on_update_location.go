package skuser

import (
	"fmt"
	"food_delivery_be/common"
	"food_delivery_be/component"
	socketio "github.com/googollee/go-socket.io"
)

type LocationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func OnUserUpdateLocation(appCtx component.AppContext, requester common.Requester) func(s socketio.Conn, location LocationData) {
	return func(s socketio.Conn, location LocationData) {

		//location belong to ????
		fmt.Println("UserId: ", requester.GetUserId(), "Location: ", location)
	}
}
