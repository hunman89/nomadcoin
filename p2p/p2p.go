package p2p

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hunman89/nomadcoin/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	_, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
}
